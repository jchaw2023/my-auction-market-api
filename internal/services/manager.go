package services

import (
	"context"
	"fmt"

	"my-auction-market-api/internal/config"
	ethclientwrapper "my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/websocket"
)

// ServiceManager 统一管理所有业务服务
// 包括：拍卖服务、出价服务、NFT服务、用户服务、监听服务等
type ServiceManager struct {
	AuctionService       *AuctionService
	BidService           *BidService
	NFTService           *NFTService
	UserService          *UserService
	ListenerService      *ListenerService
	AuctionTaskScheduler *AuctionTaskScheduler
	WSHub                *websocket.Hub
}

// NewServiceManager 创建服务管理器并初始化所有服务
func NewServiceManager(cfg config.Config) (*ServiceManager, error) {
	manager := &ServiceManager{}

	// 初始化 WebSocket Hub
	manager.WSHub = websocket.NewHub()
	go manager.WSHub.Run()
	logger.Info("websocket hub initialized and running")

	// 初始化用户服务（不需要外部依赖）
	manager.UserService = NewUserService()

	// 初始化拍卖任务调度器（需要 Redis）
	manager.AuctionTaskScheduler = GetAuctionTaskScheduler(&cfg)

	// 初始化拍卖服务（需要以太坊客户端）
	auctionService, err := NewAuctionService(cfg.Ethereum)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize auction service: %w", err)
	}
	manager.AuctionService = auctionService

	// 初始化出价服务（需要以太坊配置和客户端）
	// 创建以太坊客户端用于 USD 转换
	bidEthClient, err := ethclientwrapper.NewClient(cfg.Ethereum)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client for bid service: %w", err)
	}
	manager.BidService = NewBidService(cfg.Ethereum, bidEthClient)

	// 将任务调度器传递给拍卖服务
	manager.AuctionService.SetTaskScheduler(manager.AuctionTaskScheduler)

	// 初始化NFT服务（需要以太坊客户端和Etherscan配置）
	nftService, err := NewNFTService(cfg.Ethereum, cfg.Etherscan)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NFT service: %w", err)
	}
	manager.NFTService = nftService

	// 初始化区块链事件监听服务（可选，失败不影响其他服务）
	// 检查是否配置了合约地址和 WSS URL
	logger.Info("checking listener service configuration - contract_address: %s, wss_url: %s, rpc_url: %s",
		cfg.Ethereum.AuctionContractAddress, cfg.Ethereum.WssURL, cfg.Ethereum.RPCURL)

	if cfg.Ethereum.AuctionContractAddress == "" {
		logger.Warn("auction contract address not configured, listener service will not be initialized")
	} else if cfg.Ethereum.WssURL == "" && cfg.Ethereum.RPCURL == "" {
		logger.Warn("neither WSS URL nor RPC URL configured, listener service will not be initialized")
	} else {
		logger.Info("attempting to initialize listener service...")
		// 将服务管理器和 WebSocket Hub 传递给监听服务，以便在事件处理时访问其他业务服务和推送消息
		listenerService, err := NewListenerService(cfg.Ethereum, manager, manager.WSHub)
		if err != nil {
			logger.Error("failed to initialize listener service: %v", err)
			logger.Warn("blockchain event listener will not be available")
			logger.Info("this is a non-critical service, application will continue without it")
			// 不返回错误，允许在没有监听服务的情况下继续
		} else {
			manager.ListenerService = listenerService
			logger.Info("blockchain event listener service initialized successfully")
			logger.Info("listener service ready to start - contract: %s, wss_url: %s",
				cfg.Ethereum.AuctionContractAddress, cfg.Ethereum.WssURL)
		}
	}

	return manager, nil
}

// StartListenerService 启动监听服务（如果已初始化）
func (sm *ServiceManager) StartListenerService() error {
	if sm.ListenerService == nil {
		logger.Info("listener service is not initialized, skipping startup")
		logger.Info("this may be due to missing configuration or initialization failure (check logs above)")
		return nil // 返回 nil 而不是错误，因为这是可选服务
	}

	if err := sm.ListenerService.Start(); err != nil {
		return fmt.Errorf("failed to start listener service: %w", err)
	}

	logger.Info("blockchain event listener service started")
	return nil
}

// StopListenerService 停止监听服务（如果已初始化）
func (sm *ServiceManager) StopListenerService() error {
	if sm.ListenerService == nil {
		return nil // 服务未初始化，无需停止
	}

	if err := sm.ListenerService.Stop(); err != nil {
		return fmt.Errorf("failed to stop listener service: %w", err)
	}

	logger.Info("blockchain event listener service stopped")
	return nil
}

// StartAuctionTaskScheduler 启动拍卖任务调度器
func (sm *ServiceManager) StartAuctionTaskScheduler(ctx context.Context) {
	if sm.AuctionTaskScheduler != nil {
		sm.AuctionTaskScheduler.StartAsync(ctx)
		logger.Info("Auction task scheduler started")
	}
}

// Close 关闭所有服务并释放资源
func (sm *ServiceManager) Close() error {
	// 关闭拍卖任务调度器
	if sm.AuctionTaskScheduler != nil {
		sm.AuctionTaskScheduler.Shutdown()
	}

	// 停止监听服务
	if err := sm.StopListenerService(); err != nil {
		logger.Error("error stopping listener service: %v", err)
	}

	// 关闭拍卖服务的以太坊客户端
	if sm.AuctionService != nil {
		sm.AuctionService.Close()
	}

	// 关闭NFT服务的以太坊客户端
	if sm.NFTService != nil {
		sm.NFTService.Close()
	}

	// 关闭出价服务的以太坊客户端
	if sm.BidService != nil {
		sm.BidService.Close()
	}

	return nil
}
