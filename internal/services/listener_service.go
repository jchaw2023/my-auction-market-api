package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ethclientpkg "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/contracts/my_auction"
	ethclientwrapper "my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/websocket"
)

// ListenerService 链上事件监听服务（使用事件订阅方式）
// 包含两部分功能：
// 1. 拍卖合约事件监听：监听拍卖合约发出的所有事件（AuctionCreated, BidPlaced 等）
// 2. 钱包授权事件监听：监听钱包地址的 ERC721 Approval 事件
type ListenerService struct {
	// ========== 基础配置 ==========
	ethClient              *ethclientwrapper.Client
	client                 *ethclient.Client
	config                 config.EthereumConfig
	auctionContractAddress common.Address                   // 拍卖合约地址
	auctionContract        *my_auction.MyXAuctionV2Filterer // 拍卖合约过滤器，用于解析事件

	// 服务管理器（用于访问其他业务服务）
	serviceManager *ServiceManager

	// ========== 监听控制（共享）==========
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	isRunning bool
	mu        sync.RWMutex

	// ========== 拍卖合约事件监听（现有功能）==========
	auctionContractLogsSub           ethclientpkg.Subscription // 拍卖合约日志订阅
	auctionContractReconnectAttempts int                       // 拍卖合约重连尝试次数

	// ========== 钱包授权事件监听（新增功能）==========
	walletSubscriptions         map[common.Address]*WalletSubscription // 每个钱包地址的订阅信息
	walletSubscriptionsMu       sync.RWMutex                           // 保护钱包订阅映射的锁
	walletApprovalLogsChan      chan types.Log                         // 共享的钱包授权日志通道（所有订阅的事件都发送到这里）
	walletApprovalHandlerCtx    context.Context                        // 钱包授权处理 goroutine 的上下文
	walletApprovalHandlerCancel context.CancelFunc                     // 钱包授权处理 goroutine 的取消函数
	walletApprovalHandlerWg     sync.WaitGroup                         // 钱包授权处理 goroutine 的等待组

	// ========== 监听配置（共享）==========
	confirmations uint64 // 确认区块数

	// ========== WebSocket 心跳机制 ==========
	heartbeatInterval time.Duration // 心跳间隔（默认30秒）
	heartbeatCtx      context.Context
	heartbeatCancel   context.CancelFunc
	heartbeatWg       sync.WaitGroup

	// ========== WebSocket Hub（用于向前端推送消息）==========
	wsHub *websocket.Hub
}

// WalletSubscription 单个钱包地址的订阅信息
type WalletSubscription struct {
	walletAddress common.Address            // 钱包地址
	subscription  ethclientpkg.Subscription // 订阅对象
	ctx           context.Context           // 上下文
	cancel        context.CancelFunc        // 取消函数
	logsChan      chan types.Log            // 订阅的日志通道
}

// NewListenerService 创建新的监听服务实例
// serviceManager 用于在事件处理时访问其他业务服务
// wsHub 用于向前端推送实时消息
func NewListenerService(ethCfg config.EthereumConfig, serviceManager *ServiceManager, wsHub *websocket.Hub) (*ListenerService, error) {
	// 初始化以太坊客户端
	ethClient, err := ethclientwrapper.NewClientWithWSS(ethCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client: %w", err)
	}

	auctionContractAddress := common.HexToAddress(ethCfg.AuctionContractAddress)
	if auctionContractAddress == (common.Address{}) {
		return nil, fmt.Errorf("auction contract address is not configured")
	}

	// 创建拍卖合约过滤器实例，用于解析事件
	auctionContractFilterer, err := my_auction.NewMyXAuctionV2Filterer(auctionContractAddress, ethClient.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create auction contract filterer: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	heartbeatCtx, heartbeatCancel := context.WithCancel(context.Background())

	// 根据 WebSocket 超时时间计算心跳间隔
	// 心跳间隔 = WebSocket 超时时间的 60%（确保在超时前至少发送一次心跳）
	websocketTimeout := ethCfg.WebSocketTimeout
	if websocketTimeout == 0 {
		websocketTimeout = 60 * time.Second // 默认60秒超时
	}
	heartbeatInterval := time.Duration(float64(websocketTimeout) * 0.6) // 60% 的超时时间
	// 限制心跳间隔范围：最小15秒，最大60秒
	if heartbeatInterval < 15*time.Second {
		heartbeatInterval = 15 * time.Second
	}
	if heartbeatInterval > 60*time.Second {
		heartbeatInterval = 60 * time.Second
	}

	return &ListenerService{
		ethClient:              ethClient,
		client:                 ethClient.GetClient(),
		config:                 ethCfg,
		auctionContractAddress: auctionContractAddress,
		auctionContract:        auctionContractFilterer,
		serviceManager:         serviceManager,
		wsHub:                  wsHub,
		ctx:                    ctx,
		cancel:                 cancel,
		confirmations:          12, // 默认等待12个确认
		walletSubscriptions:    make(map[common.Address]*WalletSubscription),
		heartbeatInterval:      heartbeatInterval, // 根据 WebSocket 超时时间自动计算
		heartbeatCtx:           heartbeatCtx,
		heartbeatCancel:        heartbeatCancel,
	}, nil
}

// Start 启动监听服务（启动拍卖合约事件监听和钱包授权事件监听）
func (s *ListenerService) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		return fmt.Errorf("listener service is already running")
	}

	logger.Info("starting blockchain event listener service")
	logger.Info("auction contract address: %s", s.auctionContractAddress.Hex())
	logger.Info("confirmations: %d", s.confirmations)
	logger.Info("websocket timeout: %v, heartbeat interval: %v", s.config.WebSocketTimeout, s.heartbeatInterval)

	s.isRunning = true

	// 启动 WebSocket 心跳机制
	s.wg.Add(1)
	go s.startHeartbeat()

	// 启动拍卖合约事件订阅
	s.wg.Add(1)
	go s.subscribeAuctionContractLogs()

	// 从数据库加载所有用户的钱包地址并纳入监听
	s.wg.Add(1)
	go s.loadAllUserWalletAddresses()

	return nil
}

// Stop 停止监听服务（停止所有监听）
func (s *ListenerService) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		return fmt.Errorf("listener service is not running")
	}

	logger.Info("stopping blockchain listener service")

	// 取消拍卖合约事件订阅
	if s.auctionContractLogsSub != nil {
		s.auctionContractLogsSub.Unsubscribe()
	}

	// 取消所有钱包地址的订阅
	s.walletSubscriptionsMu.Lock()
	subscriptionCount := len(s.walletSubscriptions)
	logger.Info("stopping %d wallet subscription(s)", subscriptionCount)

	for addr, walletSub := range s.walletSubscriptions {
		logger.Debug("stopping subscription for wallet: %s", addr.Hex())
		if walletSub.subscription != nil {
			walletSub.subscription.Unsubscribe()
		}
		walletSub.cancel()
	}
	s.walletSubscriptionsMu.Unlock()

	// 停止共享的钱包授权处理 goroutine
	if s.walletApprovalHandlerCancel != nil {
		s.walletApprovalHandlerCancel()
	}
	if s.walletApprovalLogsChan != nil {
		close(s.walletApprovalLogsChan)
	}

	// 取消上下文
	s.cancel()

	// 停止心跳机制
	if s.heartbeatCancel != nil {
		s.heartbeatCancel()
	}

	// 等待所有 goroutine 完成
	s.wg.Wait()

	// 等待心跳 goroutine 完成（heartbeatWg 是 WaitGroup，不需要 nil 检查）
	s.heartbeatWg.Wait()

	// 等待钱包授权处理 goroutine 完成
	if s.walletApprovalHandlerCancel != nil {
		s.walletApprovalHandlerWg.Wait()
	}

	// 等待所有钱包订阅的转发 goroutine 完成
	// 转发 goroutine 会在 cancel 后自动退出，通过 wg.Wait() 等待

	s.isRunning = false

	logger.Info("blockchain listener service stopped")

	return nil
}

// IsRunning 检查服务是否正在运行
func (s *ListenerService) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isRunning
}

// SetConfirmations 设置确认区块数
func (s *ListenerService) SetConfirmations(confirmations uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.confirmations = confirmations
}

// ========== 拍卖合约事件监听相关方法 ==========

// subscribeAuctionContractLogs 订阅拍卖合约日志事件
// 使用 WebSocket 订阅方式实时接收拍卖合约事件，无需轮询
// 订阅后会持续监听拍卖合约地址发出的所有事件日志
func (s *ListenerService) subscribeAuctionContractLogs() {
	defer s.wg.Done()

	// 构建过滤查询 - 只订阅拍卖合约地址的日志
	query := ethclientpkg.FilterQuery{
		Addresses: []common.Address{s.auctionContractAddress},
		// Topics 可以根据需要添加特定事件的过滤
		// 例如：只订阅特定事件签名
		// Topics: [][]common.Hash{
		//     {common.HexToHash("0x...")}, // 事件签名的哈希
		// },
	}

	// 创建日志通道和订阅
	// SubscribeFilterLogs 会通过 WebSocket 实时推送匹配的日志
	logsChan := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(s.ctx, query, logsChan)
	if err != nil {
		logger.Error("failed to subscribe to auction contract logs: %v", err)
		logger.Error("note: SubscribeFilterLogs requires WebSocket connection (ws:// or wss://)")
		return
	}

	// 保存订阅引用，以便后续取消订阅
	s.mu.Lock()
	s.auctionContractLogsSub = sub
	s.mu.Unlock()

	logger.Info("successfully subscribed to auction contract logs at address: %s", s.auctionContractAddress.Hex())

	// 实时监听新日志事件
	logger.Info("listening for new auction contract events...")
	for {
		select {
		case <-s.ctx.Done():
			logger.Info("auction contract log subscription stopped by context")
			return

		case err := <-sub.Err():
			logger.Error("auction contract log subscription error: %v", err)

			// WebSocket 连接错误，需要重连
			logger.Warn("auction contract subscription disconnected, attempting to reconnect...")

			// 取消旧订阅
			sub.Unsubscribe()

			// 重连逻辑（带退避策略）
			maxRetries := 10
			retryDelay := 5 // 秒
			reconnected := false

			for attempt := 1; attempt <= maxRetries; attempt++ {
				// 检查服务是否还在运行
				if !s.isRunning {
					logger.Info("listener service stopped, aborting reconnection")
					return
				}

				// 等待一段时间后重试（第一次立即重试）
				if attempt > 1 {
					select {
					case <-s.ctx.Done():
						logger.Info("auction contract subscription stopped by context during reconnection")
						return
					case <-time.After(time.Duration(retryDelay) * time.Second):
						// 继续重连
					}
				}

				logger.Info("reconnecting auction contract subscription (attempt %d/%d)...", attempt, maxRetries)

				// 尝试重新订阅
				newSub, newErr := s.client.SubscribeFilterLogs(s.ctx, query, logsChan)
				if newErr != nil {
					logger.Error("failed to reconnect auction contract subscription (attempt %d/%d): %v", attempt, maxRetries, newErr)
					// 指数退避：每次重试延迟时间翻倍，但不超过60秒
					retryDelay = retryDelay * 2
					if retryDelay > 60 {
						retryDelay = 60
					}
					continue
				}

				// 重连成功
				logger.Info("successfully reconnected auction contract subscription")

				// 更新订阅引用
				s.mu.Lock()
				s.auctionContractLogsSub = newSub
				s.auctionContractReconnectAttempts = 0
				s.mu.Unlock()

				// 更新局部变量，继续监听新订阅
				sub = newSub
				reconnected = true
				break
			}

			// 如果重连失败，记录错误并返回（让外层循环重新启动订阅）
			if !reconnected {
				logger.Error("failed to reconnect auction contract subscription after %d attempts", maxRetries)
				s.auctionContractReconnectAttempts++
				// 返回，让 Start() 方法可以重新启动订阅
				return
			}

		case log := <-logsChan:
			// 收到新的日志事件
			logger.Debug("received new auction contract log event: block=%d, tx=%s, topics=%d",
				log.BlockNumber, log.TxHash.Hex(), len(log.Topics))

			// 检查确认数 - 确保区块有足够的确认后再处理
			// 这可以防止链重组导致的数据不一致
			// TODO: 暂时注释掉，后续根据需要启用
			// if err := s.checkConfirmations(log.BlockNumber); err != nil {
			// 	logger.Debug("auction contract log not confirmed yet, waiting for more confirmations: block=%d, error=%v", log.BlockNumber, err)
			// 	// 注意：这里直接 continue，日志会被跳过
			// 	// 如果需要处理未确认的日志，可以考虑放入队列延迟处理
			// 	continue
			// }

			// 处理日志事件
			// processAuctionContractLog 会根据事件类型路由到相应的处理函数
			if err := s.processAuctionContractLog(&log); err != nil {
				logger.Error("failed to process auction contract log at block %d, tx: %s: %v",
					log.BlockNumber, log.TxHash.Hex(), err)
				// 继续处理其他日志，不中断整个流程
				// 单个日志处理失败不应该影响其他日志的处理
				continue
			}
		}
	}
}

// processAuctionContractLog 处理拍卖合约日志，根据事件类型路由到相应的处理函数
func (s *ListenerService) processAuctionContractLog(log *types.Log) error {
	if len(log.Topics) == 0 {
		logger.Debug("auction contract log has no topics, skipping: block=%d, tx=%s",
			log.BlockNumber, log.TxHash.Hex())
		return nil
	}

	eventSignature := log.Topics[0]

	// 根据事件签名路由到不同的处理函数
	// 使用合约绑定中的 Parse 方法来解析事件

	// BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount, address indexed paymentToken, uint256 bidCount)
	if event, err := s.auctionContract.ParseBidPlaced(*log); err == nil {
		return s.handleAuctionBidPlaced(event, log)
	}

	// AuctionCreated(uint256 indexed auctionId, address indexed creator, address indexed nftAddress, uint256 tokenId)
	if event, err := s.auctionContract.ParseAuctionCreated(*log); err == nil {
		return s.handleAuctionCreated(event, log)
	}

	// AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 finalBid, address seller, address paymentToken)
	if event, err := s.auctionContract.ParseAuctionEnded(*log); err == nil {
		return s.handleAuctionEnded(event, log)
	}

	// AuctionCancelled(uint256 indexed auctionId, address indexed cancelledBy, address indexed bidder, uint256 refundAmount)
	if event, err := s.auctionContract.ParseAuctionCancelled(*log); err == nil {
		return s.handleAuctionCancelled(event, log)
	}

	// AuctionForceEnded(uint256 indexed auctionId, address indexed endedBy)
	if event, err := s.auctionContract.ParseAuctionForceEnded(*log); err == nil {
		return s.handleAuctionForceEnded(event, log)
	}

	// PlatformFeeUpdated(uint256 oldFee, uint256 newFee)
	if event, err := s.auctionContract.ParsePlatformFeeUpdated(*log); err == nil {
		return s.handleAuctionPlatformFeeUpdated(event, log)
	}

	// FeeTierUpdated(uint256 indexed tierIndex, uint256 threshold, uint256 feeRate)
	if event, err := s.auctionContract.ParseFeeTierUpdated(*log); err == nil {
		return s.handleAuctionFeeTierUpdated(event, log)
	}

	// DynamicFeeEnabled(bool enabled)
	if event, err := s.auctionContract.ParseDynamicFeeEnabled(*log); err == nil {
		return s.handleAuctionDynamicFeeEnabled(event, log)
	}

	// Paused(address account)
	if event, err := s.auctionContract.ParsePaused(*log); err == nil {
		return s.handleAuctionPaused(event, log)
	}

	// Unpaused(address account)
	if event, err := s.auctionContract.ParseUnpaused(*log); err == nil {
		return s.handleAuctionUnpaused(event, log)
	}

	// TotalValueLockedUpdated(uint256 newTVL, uint256 totalBidsPlaced, uint256 change, bool isIncrease)
	if event, err := s.auctionContract.ParseTotalValueLockedUpdated(*log); err == nil {
		return s.handleAuctionTotalValueLockedUpdated(event, log)
	}

	// NFTApproved(address indexed owner, address indexed nftAddress, uint256 tokenId)
	if event, err := s.auctionContract.ParseNFTApproved(*log); err == nil {
		return s.handleAuctionNFTApproved(event, log)
	}

	// NFTApprovalCancelled(address indexed owner, address indexed nftAddress, uint256 tokenId)
	if event, err := s.auctionContract.ParseNFTApprovalCancelled(*log); err == nil {
		return s.handleAuctionNFTApprovalCancelled(event, log)
	}

	// 如果所有事件都不匹配，记录未知事件
	logger.Debug("unknown auction contract event signature: %s, block=%d, tx=%s", eventSignature.Hex(), log.BlockNumber, log.TxHash.Hex())
	return nil
}

// ========== 钱包授权事件监听相关方法 ==========

// AddWalletAddress 添加要监听的钱包地址（创建独立订阅，但共享处理 goroutine）
func (s *ListenerService) AddWalletAddress(walletAddress common.Address) error {
	if walletAddress == (common.Address{}) {
		return fmt.Errorf("wallet address cannot be zero")
	}

	s.walletSubscriptionsMu.Lock()
	defer s.walletSubscriptionsMu.Unlock()

	// 检查是否已存在
	if _, exists := s.walletSubscriptions[walletAddress]; exists {
		logger.Debug("wallet address already being monitored: %s", walletAddress.Hex())
		return nil
	}

	// 如果监听器未运行，只记录地址，不创建订阅
	if !s.isRunning {
		logger.Debug("listener service not running, wallet address will be monitored when service starts")
		return nil
	}

	// 如果处理 goroutine 还没启动，先启动它
	if s.walletApprovalLogsChan == nil {
		s.walletApprovalLogsChan = make(chan types.Log, 100) // 带缓冲的通道
		s.walletApprovalHandlerCtx, s.walletApprovalHandlerCancel = context.WithCancel(context.Background())
		s.walletApprovalHandlerWg.Add(1)
		go s.handleAllWalletApprovalEvents() // 启动共享的处理 goroutine
		logger.Info("started shared wallet approval event handler goroutine")
	}

	// 创建独立的订阅
	sub, err := s.createWalletSubscription(walletAddress)
	if err != nil {
		return fmt.Errorf("failed to create subscription for wallet %s: %w", walletAddress.Hex(), err)
	}

	s.walletSubscriptions[walletAddress] = sub
	logger.Info("added wallet address to approval listener: %s (total subscriptions: %d)",
		walletAddress.Hex(), len(s.walletSubscriptions))

	// 启动一个 goroutine 将订阅的事件直接转发到共享通道
	// 注意：这个 goroutine 只负责转发，不做业务处理
	s.wg.Add(1)
	go s.forwardWalletSubscriptionToSharedChannel(walletAddress, sub)

	return nil
}

// RemoveWalletAddress 移除要监听的钱包地址（取消独立订阅）
func (s *ListenerService) RemoveWalletAddress(walletAddress common.Address) error {
	s.walletSubscriptionsMu.Lock()
	sub, exists := s.walletSubscriptions[walletAddress]
	if !exists {
		s.walletSubscriptionsMu.Unlock()
		logger.Debug("wallet address not in monitoring list: %s", walletAddress.Hex())
		return nil
	}

	delete(s.walletSubscriptions, walletAddress)
	subscriptionCount := len(s.walletSubscriptions)
	s.walletSubscriptionsMu.Unlock()

	// 取消订阅
	logger.Info("removing wallet address from approval listener: %s", walletAddress.Hex())

	// 取消订阅
	if sub.subscription != nil {
		sub.subscription.Unsubscribe()
	}

	// 取消上下文（会停止转发 goroutine）
	sub.cancel()

	logger.Info("removed wallet address from approval listener: %s (remaining subscriptions: %d)",
		walletAddress.Hex(), subscriptionCount)

	// 如果所有订阅都移除了，停止处理 goroutine
	if subscriptionCount == 0 && s.walletApprovalLogsChan != nil {
		s.walletApprovalHandlerCancel()
		s.walletApprovalHandlerWg.Wait()
		close(s.walletApprovalLogsChan)
		s.walletApprovalLogsChan = nil
		logger.Info("stopped shared wallet approval event handler goroutine (no more subscriptions)")
	}

	return nil
}

// GetWalletAddresses 获取当前监听的钱包地址列表
func (s *ListenerService) GetWalletAddresses() []common.Address {
	s.walletSubscriptionsMu.RLock()
	defer s.walletSubscriptionsMu.RUnlock()

	addresses := make([]common.Address, 0, len(s.walletSubscriptions))
	for addr := range s.walletSubscriptions {
		addresses = append(addresses, addr)
	}

	return addresses
}

// GetWalletSubscriptionCount 获取当前钱包订阅数量
func (s *ListenerService) GetWalletSubscriptionCount() int {
	s.walletSubscriptionsMu.RLock()
	defer s.walletSubscriptionsMu.RUnlock()
	return len(s.walletSubscriptions)
}

// createWalletSubscription 为单个钱包地址创建独立的订阅
func (s *ListenerService) createWalletSubscription(walletAddress common.Address) (*WalletSubscription, error) {
	// ERC721 Approval 事件签名
	approvalEventSig := common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")

	// 将钱包地址转换为 Hash（用于 Topics 过滤）
	ownerHash := common.BytesToHash(common.LeftPadBytes(walletAddress.Bytes(), 32))

	// 构建过滤查询 - 只监听这个钱包地址的授权
	query := ethclientpkg.FilterQuery{
		Addresses: nil, // 监听所有地址
		Topics: [][]common.Hash{
			{approvalEventSig}, // Topics[0]: 事件签名
			{ownerHash},        // Topics[1]: owner = 这个钱包地址
			{common.BytesToHash(common.LeftPadBytes(s.auctionContractAddress.Bytes(), 32))}, // Topics[2]: approved = 拍卖合约地址
			nil, // Topics[3]: tokenId（不过滤）
		},
	}

	// 创建上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())

	// 创建日志通道和订阅
	logsChan := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(ctx, query, logsChan)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to subscribe to wallet approval logs: %w", err)
	}

	walletSub := &WalletSubscription{
		walletAddress: walletAddress,
		subscription:  sub,
		ctx:           ctx,
		cancel:        cancel,
		logsChan:      logsChan,
	}

	logger.Info("created independent subscription for wallet: %s", walletAddress.Hex())

	return walletSub, nil
}

// forwardWalletSubscriptionToSharedChannel 将单个订阅的事件转发到共享通道（简单的转发，不做处理）
func (s *ListenerService) forwardWalletSubscriptionToSharedChannel(
	walletAddress common.Address,
	walletSub *WalletSubscription,
) {
	defer s.wg.Done()

	logger.Debug("started forwarding events for wallet: %s", walletAddress.Hex())

	sub := walletSub.subscription
	logsChan := walletSub.logsChan
	ctx := walletSub.ctx

	for {
		select {
		case <-ctx.Done():
			logger.Debug("stopped forwarding events for wallet: %s", walletAddress.Hex())
			return

		case err := <-sub.Err():
			logger.Error("subscription error for wallet %s: %v", walletAddress.Hex(), err)

			// WebSocket 连接错误，需要重连
			logger.Warn("wallet subscription disconnected for wallet %s, attempting to reconnect...", walletAddress.Hex())

			// 取消旧订阅
			sub.Unsubscribe()

			// 重连逻辑（带退避策略）
			maxRetries := 10
			retryDelay := 5 // 秒
			reconnected := false

			for attempt := 1; attempt <= maxRetries; attempt++ {
				// 检查服务是否还在运行
				if !s.isRunning {
					logger.Info("listener service stopped, aborting wallet subscription reconnection")
					return
				}

				// 检查钱包是否还在监听列表中
				s.walletSubscriptionsMu.RLock()
				_, stillExists := s.walletSubscriptions[walletAddress]
				s.walletSubscriptionsMu.RUnlock()

				if !stillExists {
					logger.Info("wallet %s removed from monitoring list, aborting reconnection", walletAddress.Hex())
					return
				}

				// 等待一段时间后重试（第一次立即重试）
				if attempt > 1 {
					select {
					case <-ctx.Done():
						logger.Info("wallet subscription stopped by context during reconnection")
						return
					case <-time.After(time.Duration(retryDelay) * time.Second):
						// 继续重连
					}
				}

				logger.Info("reconnecting wallet subscription for %s (attempt %d/%d)...", walletAddress.Hex(), attempt, maxRetries)

				// 重新创建订阅
				newSub, newErr := s.createWalletSubscription(walletAddress)
				if newErr != nil {
					logger.Error("failed to reconnect wallet subscription for %s (attempt %d/%d): %v", walletAddress.Hex(), attempt, maxRetries, newErr)
					// 指数退避：每次重试延迟时间翻倍，但不超过60秒
					retryDelay = retryDelay * 2
					if retryDelay > 60 {
						retryDelay = 60
					}
					continue
				}

				// 重连成功
				logger.Info("successfully reconnected wallet subscription for %s", walletAddress.Hex())

				// 更新订阅引用
				s.walletSubscriptionsMu.Lock()
				if existingSub, stillExists := s.walletSubscriptions[walletAddress]; stillExists {
					existingSub.subscription = newSub.subscription
					existingSub.logsChan = newSub.logsChan
					existingSub.ctx = newSub.ctx
					existingSub.cancel = newSub.cancel
					sub = newSub.subscription
					logsChan = newSub.logsChan
					ctx = newSub.ctx
				}
				s.walletSubscriptionsMu.Unlock()

				reconnected = true
				break
			}

			// 如果重连失败，记录错误并返回
			if !reconnected {
				logger.Error("failed to reconnect wallet subscription for %s after %d attempts", walletAddress.Hex(), maxRetries)
				return
			}

		case log, ok := <-logsChan:
			if !ok {
				logger.Debug("logs channel closed for wallet: %s", walletAddress.Hex())
				return
			}

			// 直接转发到共享通道（不做任何处理）
			select {
			case s.walletApprovalLogsChan <- log:
				// 成功转发
			case <-ctx.Done():
				return
			case <-s.walletApprovalHandlerCtx.Done():
				return
			}
		}
	}
}

// handleAllWalletApprovalEvents 处理所有钱包地址的 Approval 事件（共享的处理 goroutine）
func (s *ListenerService) handleAllWalletApprovalEvents() {
	defer s.walletApprovalHandlerWg.Done()

	logger.Info("started shared wallet approval event handler goroutine")

	for {
		select {
		case <-s.walletApprovalHandlerCtx.Done():
			logger.Info("shared wallet approval event handler stopped")
			return

		case log, ok := <-s.walletApprovalLogsChan:
			if !ok {
				logger.Info("wallet approval logs channel closed, stopping handler")
				return
			}

			// 收到新的 Approval 事件
			logger.Debug("received wallet approval event: block=%d, tx=%s, nftContract=%s",
				log.BlockNumber, log.TxHash.Hex(), log.Address.Hex())

			// 检查确认数
			// TODO: 暂时注释掉，后续根据需要启用
			// if err := s.checkConfirmations(log.BlockNumber); err != nil {
			// 	logger.Debug("wallet approval log not confirmed yet: block=%d, error=%v",
			// 		log.BlockNumber, err)
			// 	continue
			// }

			// 处理 Approval 事件（这里是唯一的业务处理逻辑）
			if err := s.processWalletApprovalLog(log); err != nil {
				logger.Error("failed to process wallet approval log at block %d, tx: %s: %v",
					log.BlockNumber, log.TxHash.Hex(), err)
				continue
			}
		}
	}
}

// processWalletApprovalLog 处理钱包授权 Approval 事件日志
func (s *ListenerService) processWalletApprovalLog(log types.Log) error {
	if len(log.Topics) < 4 {
		logger.Debug("invalid wallet approval log: insufficient topics, block=%d, tx=%s",
			log.BlockNumber, log.TxHash.Hex())
		return fmt.Errorf("invalid wallet approval log: insufficient topics")
	}

	// 解析事件参数
	owner := common.BytesToAddress(log.Topics[1][12:])
	approved := common.BytesToAddress(log.Topics[2][12:])
	tokenId := new(big.Int).SetBytes(log.Topics[3].Bytes())
	nftContractAddress := log.Address

	// 验证 approved 地址是否为拍卖合约
	if approved != s.auctionContractAddress {
		logger.Debug("wallet approval is not for auction contract: approved=%s, expected=%s",
			approved.Hex(), s.auctionContractAddress.Hex())
		return nil
	}

	// 验证 owner 是否在监听列表中
	s.walletSubscriptionsMu.RLock()
	_, isMonitored := s.walletSubscriptions[owner]
	s.walletSubscriptionsMu.RUnlock()

	if !isMonitored {
		logger.Debug("wallet approval owner is not in monitoring list: owner=%s", owner.Hex())
		return nil
	}

	logger.Info("Wallet ERC721 Approval event: owner=%s, nftContract=%s, tokenId=%s, approved=%s",
		owner.Hex(), nftContractAddress.Hex(), tokenId.String(), approved.Hex())

	// 调用处理函数（唯一的业务处理逻辑）
	return s.handleWalletERC721Approval(owner, nftContractAddress, tokenId.Uint64())
}

// handleWalletERC721Approval 处理钱包 ERC721 Approval 事件
func (s *ListenerService) handleWalletERC721Approval(
	owner common.Address,
	nftContractAddress common.Address,
	tokenId uint64,
) error {
	// TODO: 实现具体的业务逻辑
	// 例如：
	// - 更新数据库中的授权状态
	// - 通知用户授权成功
	// - 触发后续的业务流程

	logger.Info("processing wallet ERC721 approval: owner=%s, nftContract=%s, tokenId=%d", owner.Hex(), nftContractAddress.Hex(), tokenId)
	ownerAddress := strings.ToLower(owner.Hex())
	nftContractAddressStr := strings.ToLower(nftContractAddress.Hex())
	nftId, err := s.serviceManager.NFTService.OnNFTApproved(ownerAddress, nftContractAddressStr, tokenId)
	if err != nil {
		logger.Error("failed to process wallet ERC721 approval: owner=%s, nftContract=%s, tokenId=%d: %v",
			owner.Hex(), nftContractAddress.Hex(), tokenId, err)
		return err
	}
	// 向前端推送消息
	if s.wsHub != nil {
		message := websocket.NewMessage(websocket.MessageTypeNFTApproved, map[string]interface{}{
			"ownerAddress":    ownerAddress,
			"nftId":           nftId,
			"contractAddress": nftContractAddressStr,
			"tokenId":         tokenId,
		})
		s.wsHub.BroadcastMessage(message)
	}

	// 示例：可以调用服务管理器的方法
	// if s.serviceManager != nil && s.serviceManager.NFTService != nil {
	//     return s.serviceManager.NFTService.OnNFTApproved(
	//         owner.Hex(),
	//         nftContractAddress.Hex(),
	//         tokenId,
	//     )
	// }

	return nil
}

// ClearWalletAddresses 清空所有钱包地址订阅
func (s *ListenerService) ClearWalletAddresses() error {
	s.walletSubscriptionsMu.Lock()
	defer s.walletSubscriptionsMu.Unlock()

	count := len(s.walletSubscriptions)
	logger.Info("clearing all wallet address subscriptions (count: %d)", count)

	// 取消所有订阅
	for addr, walletSub := range s.walletSubscriptions {
		logger.Debug("stopping subscription for wallet: %s", addr.Hex())
		if walletSub.subscription != nil {
			walletSub.subscription.Unsubscribe()
		}
		walletSub.cancel()
	}

	// 清空映射
	s.walletSubscriptions = make(map[common.Address]*WalletSubscription)

	// 停止处理 goroutine
	if s.walletApprovalHandlerCancel != nil {
		s.walletApprovalHandlerCancel()
		s.walletApprovalHandlerWg.Wait()
	}
	if s.walletApprovalLogsChan != nil {
		close(s.walletApprovalLogsChan)
		s.walletApprovalLogsChan = nil
	}

	logger.Info("cleared all wallet address subscriptions")

	return nil
}

// ========== 启动时加载用户钱包地址 ==========

// loadAllUserWalletAddresses 从数据库加载所有用户的钱包地址并纳入监听
func (s *ListenerService) loadAllUserWalletAddresses() {
	defer s.wg.Done()

	logger.Info("loading all user wallet addresses from database...")

	// 获取所有用户的钱包地址
	walletAddresses, err := s.serviceManager.UserService.GetAllWalletAddresses()
	if err != nil {
		logger.Error("failed to load wallet addresses from database: %v", err)
		return
	}

	if len(walletAddresses) == 0 {
		logger.Info("no wallet addresses found in database")
		return
	}

	logger.Info("found %d wallet address(es) in database, adding to approval listener...", len(walletAddresses))

	// 为每个钱包地址添加监听
	successCount := 0
	failedCount := 0

	for _, walletAddrStr := range walletAddresses {
		// 规范化地址（去除空格，转换为小写）
		walletAddrStr = strings.TrimSpace(walletAddrStr)
		if walletAddrStr == "" {
			continue
		}

		// 转换为 common.Address
		walletAddr := common.HexToAddress(walletAddrStr)
		if walletAddr == (common.Address{}) {
			logger.Warn("invalid wallet address format: %s", walletAddrStr)
			failedCount++
			continue
		}

		// 添加监听
		if err := s.AddWalletAddress(walletAddr); err != nil {
			logger.Error("failed to add wallet address %s to listener: %v", walletAddrStr, err)
			failedCount++
			continue
		}

		successCount++
	}

	logger.Info("finished loading wallet addresses: %d successful, %d failed", successCount, failedCount)
}

// ========== WebSocket 心跳机制 ==========

// startHeartbeat 启动 WebSocket 心跳机制，定期发送 ping 保持连接活跃
func (s *ListenerService) startHeartbeat() {
	defer s.wg.Done()

	logger.Info("starting WebSocket heartbeat mechanism (interval: %v)", s.heartbeatInterval)

	ticker := time.NewTicker(s.heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.heartbeatCtx.Done():
			logger.Info("WebSocket heartbeat stopped")
			return

		case <-ticker.C:
			// 发送心跳：调用一个轻量级的 RPC 方法保持连接活跃
			// 使用 eth_blockNumber 是一个很好的选择，因为它很轻量且能验证连接
			// 心跳 RPC 调用的超时时间设置为心跳间隔的 50%，确保不会阻塞太久
			heartbeatTimeout := s.heartbeatInterval / 2
			if heartbeatTimeout < 3*time.Second {
				heartbeatTimeout = 3 * time.Second // 最小3秒
			}
			if heartbeatTimeout > 10*time.Second {
				heartbeatTimeout = 10 * time.Second // 最大10秒
			}

			ctx, cancel := context.WithTimeout(context.Background(), heartbeatTimeout)
			_, err := s.client.HeaderByNumber(ctx, nil)
			cancel()

			if err != nil {
				logger.Warn("WebSocket heartbeat failed: %v (connection may be unstable)", err)
				// 心跳失败不中断服务，只是记录警告
				// 如果连接真的断开，订阅的错误处理会触发重连
			} else {
				logger.Debug("WebSocket heartbeat sent successfully")
			}
		}
	}
}

// SetHeartbeatInterval 设置心跳间隔
func (s *ListenerService) SetHeartbeatInterval(interval time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.heartbeatInterval = interval
	logger.Info("heartbeat interval set to: %v", interval)
}

// ========== 共享工具方法 ==========

// checkConfirmations 检查日志是否有足够的确认数（拍卖合约和钱包授权共用）
func (s *ListenerService) checkConfirmations(blockNumber uint64) error {
	// currentBlock, err := s.getCurrentBlockNumber()
	// if err != nil {
	// 	return fmt.Errorf("failed to get current block number: %w", err)
	// }

	// if currentBlock < blockNumber+s.confirmations {
	// 	return fmt.Errorf("not enough confirmations: current=%d, log block=%d, required=%d",
	// 		currentBlock, blockNumber, s.confirmations)
	// }

	return nil
}

// getCurrentBlockNumber 获取当前区块号（共享方法）
func (s *ListenerService) getCurrentBlockNumber() (uint64, error) {
	header, err := s.client.HeaderByNumber(s.ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block header: %w", err)
	}
	return header.Number.Uint64(), nil
}

// Close 关闭服务并释放资源
func (s *ListenerService) Close() error {
	if s.IsRunning() {
		if err := s.Stop(); err != nil {
			return err
		}
	}

	if s.ethClient != nil {
		s.ethClient.Close()
	}

	return nil
}

// ============ 拍卖合约事件处理函数 ============

// handleAuctionBidPlaced 处理出价事件
func (s *ListenerService) handleAuctionBidPlaced(event *my_auction.MyXAuctionV2BidPlaced, log *types.Log) error {
	logger.Info("Auction BidPlaced event: auctionId=%d, bidder=%s, amount=%s, paymentToken=%s, bidCount=%s, block=%d, tx=%s",
		event.AuctionId, event.Bidder.Hex(), event.Amount.String(), event.PaymentToken.Hex(), event.BidCount.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现出价事件处理逻辑
	// 例如：更新数据库中的出价记录、更新拍卖的最高出价等

	// 调用 BidService 处理出价事件
	bid, err := s.serviceManager.BidService.OnEventBidPlaced(event, log)
	if err != nil {
		logger.Error("failed to process bid placed event: %v", err)
		return err
	}
	// 向前端推送消息 - 只推送给订阅了该拍卖的客户端
	if s.wsHub != nil && bid != nil {
		// 使用 BidService 的转换方法，统一数据格式
		bidResponse := s.serviceManager.BidService.ConvertBidToResponse(bid)
		message := websocket.NewMessage(websocket.MessageTypeAuctionBidPlaced, bidResponse)

		// 构建房间ID：auction:{auctionID}
		roomID := fmt.Sprintf("auction:%s", bid.AuctionID)

		// 只推送给订阅了该拍卖的客户端
		if err := s.wsHub.BroadcastToRoom(roomID, message); err != nil {
			logger.Error("failed to broadcast bid message to room %s: %v", roomID, err)
		}
	}

	return nil
}

// handleAuctionCreated 处理拍卖创建事件
func (s *ListenerService) handleAuctionCreated(event *my_auction.MyXAuctionV2AuctionCreated, log *types.Log) error {
	logger.Info("Auction Created event: auctionId=%d, creator=%s, nftAddress=%s, tokenId=%s, block=%d, tx=%s",
		event.AuctionId, event.Creator.Hex(), event.NftAddress.Hex(), event.TokenId.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现拍卖创建事件处理逻辑
	// 例如：在数据库中创建新的拍卖记录
	// 这里更新拍卖状态为待上架
	auctionContractId := event.AuctionId.Uint64()
	nftAddress := strings.ToLower(event.NftAddress.Hex())
	ownerAddress := strings.ToLower(event.Creator.Hex())
	tokenId := event.TokenId.Uint64()
	s.serviceManager.AuctionService.OnEventAuctionCreated(auctionContractId, ownerAddress, nftAddress, tokenId)
	// 向前端推送消息
	if s.wsHub != nil {
		message := websocket.NewMessage(websocket.MessageTypeAuctionCreated, map[string]interface{}{
			"auctionContractId": auctionContractId,
			"ownerAddress":      ownerAddress,
			"nftAddress":        nftAddress,
			"tokenId":           tokenId,
		})
		s.wsHub.BroadcastMessage(message)
	}

	return nil
}

// handleAuctionEnded 处理拍卖结束事件
func (s *ListenerService) handleAuctionEnded(event *my_auction.MyXAuctionV2AuctionEnded, log *types.Log) error {
	logger.Info("Auction Ended event: auctionId=%d, winner=%s, finalBid=%s, seller=%s, paymentToken=%s, block=%d, tx=%s",
		event.AuctionId, event.Winner.Hex(), event.FinalBid.String(), event.Seller.Hex(), event.PaymentToken.Hex(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现拍卖结束事件处理逻辑
	// 例如：更新拍卖状态为已结束、记录获胜者信息等

	contractAuctionId := event.AuctionId.Uint64()
	winner := strings.ToLower(event.Winner.Hex())
	finalBid := event.FinalBid.Uint64()
	seller := strings.ToLower(event.Seller.Hex())
	paymentToken := strings.ToLower(event.PaymentToken.Hex())
	bidValue := event.BidValue.Uint64()
	minBidValue := event.MinBidValue.Uint64()
	// 向前端推送消息
	if s.wsHub != nil {
		auction, err := s.serviceManager.AuctionService.GetByContractID(contractAuctionId)
		if err != nil {
			logger.Error("failed to get auction by contract id: %v", err)
			return err
		}
		if auction == nil {
			logger.Error("auction not found by contract id: %d", contractAuctionId)
			return fmt.Errorf("auction not found by contract id: %d", contractAuctionId)
		}
		if event.Winner == common.BigToAddress(big.NewInt(0)) {
			message := websocket.NewMessage(websocket.MessageTypeAuctionEnded, map[string]interface{}{
				"auctionId":    contractAuctionId,
				"winner":       winner,
				"finalBid":     finalBid,
				"seller":       seller,
				"paymentToken": paymentToken,
				"bidValue":     bidValue,
				"minBidValue":  minBidValue,
				"nftName":      auction.NftName,
				"nftId":        auction.NFTID,
			})
			s.wsHub.BroadcastMessage(message)
		} else {

			usdValue, err := s.serviceManager.AuctionService.ConvertToUSDFromTokenUnit(paymentToken, big.NewInt(int64(finalBid)))
			if err != nil {
				logger.Error("failed to convert to USD: %v", err)
				return err
			}
			message := websocket.NewMessage(websocket.MessageTypeAuctionEnded, map[string]interface{}{
				"auctionId":    contractAuctionId,
				"winner":       winner,
				"finalBid":     finalBid,
				"seller":       seller,
				"paymentToken": paymentToken,
				"bidValue":     bidValue,
				"minBidValue":  minBidValue,
				"usdValueStr":  usdValue.AmountUSDStr,
				"usdValue":     usdValue.AmountUSD,
				"nftName":      auction.NftName,
				"nftId":        auction.NFTID,
			})
			s.wsHub.BroadcastMessage(message)
		}

	}

	return nil
}

// handleAuctionCancelled 处理拍卖取消事件
func (s *ListenerService) handleAuctionCancelled(event *my_auction.MyXAuctionV2AuctionCancelled, log *types.Log) error {
	logger.Info("Auction Cancelled event: auctionId=%d, cancelledBy=%s, bidder=%s, refundAmount=%s, block=%d, tx=%s",
		event.AuctionId, event.CancelledBy.Hex(), event.Bidder.Hex(), event.RefundAmount.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现拍卖取消事件处理逻辑
	// 例如：更新拍卖状态为已取消、处理退款等
	auctionContractId := event.AuctionId.Uint64()
	cancelledBy := strings.ToLower(event.CancelledBy.Hex())
	bidder := strings.ToLower(event.Bidder.Hex())
	paymentToken := strings.ToLower(event.PaymentToken.Hex())
	refundAmount := event.RefundAmount.Uint64()
	refundAmountValue := event.RefundAmountValue.Uint64()
	auctionId, err := s.serviceManager.AuctionService.OnEventAuctionCancelled(auctionContractId, cancelledBy, bidder, paymentToken, refundAmount, refundAmountValue)
	if err != nil {
		logger.Error("failed to process auction cancelled event: %v", err)
		return err
	}

	// 删除调度器上的时间调度
	if s.serviceManager.AuctionTaskScheduler != nil && auctionId != "" {
		if err := s.serviceManager.AuctionTaskScheduler.CancelAuctionEndTask(auctionId); err != nil {
			logger.Error("failed to cancel auction end task for auction %s: %v", auctionId, err)
			// 不返回错误，因为取消调度失败不应该阻止拍卖取消流程
		} else {
			logger.Info("cancelled auction end task for auction: %s", auctionId)
		}
	}

	// 向前端推送消息
	if s.wsHub != nil {
		message := websocket.NewMessage(websocket.MessageTypeAuctionCancelled, map[string]interface{}{
			"auctionId":         auctionContractId,
			"cancelledBy":       cancelledBy,
			"bidder":            bidder,
			"paymentToken":      paymentToken,
			"refundAmount":      refundAmount,
			"refundAmountValue": refundAmountValue,
		})
		s.wsHub.BroadcastMessage(message)
	}

	return nil
}

// handleAuctionForceEnded 处理强制结束拍卖事件
func (s *ListenerService) handleAuctionForceEnded(event *my_auction.MyXAuctionV2AuctionForceEnded, log *types.Log) error {
	logger.Info("Auction ForceEnded event: auctionId=%d, endedBy=%s, block=%d, tx=%s",
		event.AuctionId, event.EndedBy.Hex(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现强制结束拍卖事件处理逻辑

	// 向前端推送消息
	if s.wsHub != nil {
		message := websocket.NewMessage(websocket.MessageTypeAuctionForceEnded, map[string]interface{}{
			"auctionId": event.AuctionId.String(),
			"endedBy":   event.EndedBy.Hex(),
		})
		s.wsHub.BroadcastMessage(message)
	}

	return nil
}

// handleAuctionPlatformFeeUpdated 处理平台手续费更新事件
func (s *ListenerService) handleAuctionPlatformFeeUpdated(event *my_auction.MyXAuctionV2PlatformFeeUpdated, log *types.Log) error {
	logger.Info("Auction PlatformFeeUpdated event: oldFee=%s, newFee=%s, block=%d, tx=%s",
		event.OldFee.String(), event.NewFee.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现平台手续费更新事件处理逻辑
	return nil
}

// handleAuctionFeeTierUpdated 处理手续费档次更新事件
func (s *ListenerService) handleAuctionFeeTierUpdated(event *my_auction.MyXAuctionV2FeeTierUpdated, log *types.Log) error {
	logger.Info("Auction FeeTierUpdated event: tierIndex=%s, threshold=%s, feeRate=%s, block=%d, tx=%s",
		event.TierIndex.String(), event.Threshold.String(), event.FeeRate.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现手续费档次更新事件处理逻辑
	return nil
}

// handleAuctionDynamicFeeEnabled 处理动态手续费启用/禁用事件
func (s *ListenerService) handleAuctionDynamicFeeEnabled(event *my_auction.MyXAuctionV2DynamicFeeEnabled, log *types.Log) error {
	logger.Info("Auction DynamicFeeEnabled event: enabled=%v, block=%d, tx=%s",
		event.Enabled, log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现动态手续费启用/禁用事件处理逻辑
	return nil
}

// handleAuctionPaused 处理合约暂停事件
func (s *ListenerService) handleAuctionPaused(event *my_auction.MyXAuctionV2Paused, log *types.Log) error {
	logger.Info("Auction Paused event: account=%s, block=%d, tx=%s",
		event.Account.Hex(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现合约暂停事件处理逻辑
	return nil
}

// handleAuctionUnpaused 处理合约取消暂停事件
func (s *ListenerService) handleAuctionUnpaused(event *my_auction.MyXAuctionV2Unpaused, log *types.Log) error {
	logger.Info("Auction Unpaused event: account=%s, block=%d, tx=%s",
		event.Account.Hex(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现合约取消暂停事件处理逻辑
	return nil
}

// handleAuctionTotalValueLockedUpdated 处理TVL更新事件
func (s *ListenerService) handleAuctionTotalValueLockedUpdated(event *my_auction.MyXAuctionV2TotalValueLockedUpdated, log *types.Log) error {
	logger.Info("Auction TotalValueLockedUpdated event: newTVL=%s, totalBidsPlaced=%s, change=%s, isIncrease=%v, block=%d, tx=%s",
		event.NewTVL.String(), event.TotalBidsPlaced.String(), event.Change.String(), event.IsIncrease, log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现TVL更新事件处理逻辑
	return nil
}

// handleAuctionNFTApproved 处理NFT批准事件（拍卖合约发出的自定义事件）
func (s *ListenerService) handleAuctionNFTApproved(event *my_auction.MyXAuctionV2NFTApproved, log *types.Log) error {
	logger.Info("Auction NFTApproved event: owner=%s, nftAddress=%s, tokenId=%s, block=%d, tx=%s",
		event.Owner.Hex(), event.NftAddress.Hex(), event.TokenId.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现NFT批准事件处理逻辑
	return nil
}

// handleAuctionNFTApprovalCancelled 处理NFT取消批准事件（拍卖合约发出的自定义事件）
func (s *ListenerService) handleAuctionNFTApprovalCancelled(event *my_auction.MyXAuctionV2NFTApprovalCancelled, log *types.Log) error {
	logger.Info("Auction NFTApprovalCancelled event: owner=%s, nftAddress=%s, tokenId=%s, block=%d, tx=%s",
		event.Owner.Hex(), event.NftAddress.Hex(), event.TokenId.String(), log.BlockNumber, log.TxHash.Hex())
	// TODO: 实现NFT取消批准事件处理逻辑
	return nil
}
