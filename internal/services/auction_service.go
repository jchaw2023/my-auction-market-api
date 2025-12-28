package services

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"

	"my-auction-market-api/internal/config"
	erc721_nft "my-auction-market-api/internal/contracts/erc721_nft"
	"my-auction-market-api/internal/contracts/my_auction"
	"my-auction-market-api/internal/database"
	"my-auction-market-api/internal/errors"
	"my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/page"
	"my-auction-market-api/internal/utils"
)

var (
	// snowflakeGenerator 全局 snowflake ID 生成器
	snowflakeGenerator *sonyflake.Sonyflake
)

// 拍卖状态常量
const (
	AuctionStatusPending   = "pending"   // 待上架
	AuctionStatusActive    = "active"    // 已上架/进行中
	AuctionStatusEnded     = "ended"     // 已结束
	AuctionStatusCancelled = "cancelled" // 已取消
)

func init() {
	// 初始化 snowflake 生成器
	snowflakeGenerator = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), // 设置起始时间
	})
	if snowflakeGenerator == nil {
		logger.Error("failed to initialize Sonyflake generator")
	}
}

type AuctionService struct {
	ethClient     *ethereum.Client
	config        config.EthereumConfig
	taskScheduler *AuctionTaskScheduler
}

func (s *AuctionService) OnEventAuctionCancelled(auctionContractId uint64,
	cancelledBy string, bidder string, paymentToken string, refundAmount uint64, refundAmountValue uint64) (string, error) {
	//修改nft_ownerships
	var auctionId string
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var auction models.Auction
		if err := tx.Where("contract_auction_id = ?", auctionContractId).First(&auction).Error; err != nil {
			logger.Error("failed to get auction: %v", err)
			return err
		}
		var nftOwnership models.NFTOwnership
		if err := tx.Where("nft_id = ? AND user_id = ?", auction.NFTID, auction.UserID).First(&nftOwnership).Error; err != nil {
			logger.Error("failed to get nft_ownership: %v", err)
			return err
		}

		if err := tx.Model(&models.NFTOwnership{}).Where("nft_id = ? AND user_id = ?", auction.NFTID, auction.UserID).
			Updates(map[string]interface{}{"status": models.NFTOwnershipStatusHolding,
				"owner_address": auction.OwnerAddress,
				"approved":      0}).Error; err != nil {
			logger.Error("failed to update nft_ownership: %v", err)
			return err
		}
		nftOnlineLock := fmt.Sprintf("%s:%s", auction.NFTID, auction.AuctionID)
		if err := tx.Model(&models.Auction{}).Where("contract_auction_id = ?", auctionContractId).
			Updates(map[string]interface{}{"status": AuctionStatusCancelled,
				"updated_at":  time.Now(),
				"online":      0,
				"online_lock": nftOnlineLock}).Error; err != nil {
			logger.Error("failed to update auction: %v", err)
			return err
		}
		//
		auctionId = auction.AuctionID
		return nil
	})
	return auctionId, err
}

func NewAuctionService(ethCfg config.EthereumConfig) (*AuctionService, error) {
	// 初始化以太坊客户端
	ethClient, err := ethereum.NewClient(ethCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Ethereum client: %w", err)
	}

	return &AuctionService{
		ethClient: ethClient,
		config:    ethCfg,
	}, nil
}

// GetEthClient 获取以太坊客户端（用于链上查询）
func (s *AuctionService) GetEthClient() *ethereum.Client {
	return s.ethClient
}

// GetConfig 获取以太坊配置
func (s *AuctionService) GetConfig() config.EthereumConfig {
	return s.config
}

// SetTaskScheduler 设置任务调度器
func (s *AuctionService) SetTaskScheduler(scheduler *AuctionTaskScheduler) {
	s.taskScheduler = scheduler
}

// ConvertTokenAmountToUSD 将代币金额转换为美元价值（服务方法，封装配置和客户端）
func (s *AuctionService) ConvertTokenAmountToUSD(tokenAddress string, amount float64) (*models.ConvertToUSDResponse, error) {
	return ConvertTokenAmountToUSD(&s.config, tokenAddress, amount, s.ethClient.GetClient())
}

// ConvertToUSDFromTokenUnit 将代币最小单位金额转换为美元价值（服务方法，封装配置和客户端）
// tokenAddress: 代币地址
// amountBigInt: 代币数量（最小单位，例如 wei）
func (s *AuctionService) ConvertToUSDFromTokenUnit(tokenAddress string, amountBigInt *big.Int) (*models.ConvertToUSDResponse, error) {
	return ConvertToUSDFromTokenUnit(&s.config, tokenAddress, amountBigInt, s.ethClient.GetClient())
}

// GetTokenPrice 获取代币价格（服务方法，封装配置和客户端）
func (s *AuctionService) GetTokenPrice(tokenAddress string) (*models.TokenPriceResponse, error) {
	return GetTokenPrice(&s.config, tokenAddress, s.ethClient.GetClient())
}

// ConvertUnitUSDToToken 将整数美元金额（使用价格精度格式）转换为代币金额
// tokenAddress: 代币地址
// amountUSDUnit: 美元金额（uint64，使用价格精度格式，精度从链上获取）
// 返回：USD金额（小数）、代币金额（小数）、代币精度
func (s *AuctionService) ConvertUnitUSDToToken(tokenAddress string, amountUSDUnit uint64) (*models.ConvertUSDToTokenResponse, error) {
	return ConvertUnitUSDToToken(&s.config, tokenAddress, amountUSDUnit, s.ethClient.GetClient())
}

// GetAuctionSimpleStats 获取拍卖简单统计信息（从合约获取）
func (s *AuctionService) GetAuctionSimpleStats() (*models.AuctionSimpleStatsResponse, error) {
	// 创建拍卖合约实例
	auctionContract, err := my_auction.NewMyXAuctionV2(common.HexToAddress(s.config.AuctionContractAddress), s.ethClient.GetClient())
	if err != nil {
		logger.Error("failed to create auction contract: %s", err.Error())
		return nil, fmt.Errorf("failed to create auction contract: %w", err)
	}

	// 调用合约的 GetAuctionSimpleStats 方法
	totalAuctionsCreated, totalBidsPlaced, platformFee, totalValueLocked, err := auctionContract.GetAuctionSimpleStats(&bind.CallOpts{})
	if err != nil {
		logger.Error("failed to get auction simple stats: %s", err.Error())
		return nil, fmt.Errorf("failed to get auction simple stats: %w", err)
	}

	// 将 totalValueLocked 从 8 位小数格式转换为 float64
	// Chainlink 使用 8 位小数格式
	totalValueLockedDecimal := decimal.NewFromBigInt(totalValueLocked, -8)
	totalValueLockedFloat := totalValueLockedDecimal.InexactFloat64()

	return &models.AuctionSimpleStatsResponse{
		TotalAuctionsCreated: totalAuctionsCreated.Uint64(),
		TotalBidsPlaced:      totalBidsPlaced.Uint64(),
		PlatformFee:          platformFee.Uint64(),
		TotalValueLocked:     totalValueLockedFloat,
		TotalValueLockedStr:  totalValueLockedDecimal.StringFixed(2), // 保留2位小数
	}, nil
}

// CheckNFTApproval 检查指定TokenID的NFT是否已授权给平台合约
func (s *AuctionService) CheckNFTApproval(nftAddress string, tokenID uint64) (bool, error) {
	platformContractAddress := common.HexToAddress(s.config.AuctionContractAddress)
	nftContractAddress := common.HexToAddress(nftAddress)

	_ethClient := s.ethClient.GetClient()
	nftContract, err := erc721_nft.NewMyNFT(nftContractAddress, _ethClient)
	if err != nil {
		return false, fmt.Errorf("failed to create NFT contract instance: %w", err)
	}

	opts := &bind.CallOpts{Context: context.Background()}
	tokenIDBigInt := big.NewInt(int64(tokenID))
	approvedAddress, err := nftContract.GetApproved(opts, tokenIDBigInt)
	if err != nil {
		return false, fmt.Errorf("failed to check GetApproved for token %d: %w", tokenID, err)
	}

	return approvedAddress == platformContractAddress, nil
}

// Close 关闭以太坊客户端连接
func (s *AuctionService) Close() {
	if s.ethClient != nil {
		s.ethClient.Close()
	}
}

func (s *AuctionService) List(query page.PageQuery) ([]models.Auction, int64, error) {
	var auctions []models.Auction
	var total int64

	// 查询条件：online == 1
	baseQuery := database.DB.Model(&models.Auction{}).Where("online = ?", 1)

	// 统计总数
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	if err := baseQuery.
		Preload("User").
		Order("created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&auctions).Error; err != nil {
		return nil, 0, err
	}

	return auctions, total, nil
}

// ListPublic 获取公开拍卖列表（用于首页展示）
// 排序规则：active 状态的排在前面，ended 状态的排在后面，每个状态内部按时间倒序
// statusFilter: 可选的状态筛选，如果为空或 "all"，则返回所有 active 和 ended 状态的数据
func (s *AuctionService) ListPublic(query page.PageQuery, statusFilter string) ([]models.Auction, int64, error) {
	var auctions []models.Auction
	var total int64

	// 查询条件：online == 1，且只包含 active 和 ended 状态
	baseQuery := database.DB.Model(&models.Auction{}).
		Where("online = ? AND status IN ?", 1, []string{AuctionStatusActive, AuctionStatusEnded})

	// 如果指定了状态筛选，进一步过滤
	if statusFilter != "" && statusFilter != "all" {
		if statusFilter == AuctionStatusActive || statusFilter == AuctionStatusEnded {
			baseQuery = baseQuery.Where("status = ?", statusFilter)
		}
	}

	// 统计总数
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	// 排序：使用 CASE WHEN 确保 active 排在前面，ended 排在后面，然后按时间倒序
	if err := baseQuery.
		Preload("User").
		Order("CASE WHEN status = 'active' THEN 0 WHEN status = 'ended' THEN 1 ELSE 2 END, created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&auctions).Error; err != nil {
		return nil, 0, err
	}

	return auctions, total, nil
}

// ListNFTs 获取全站拍卖中的 NFT 列表（去重，支持分页）
func (s *AuctionService) ListNFTs(query page.PageQuery) ([]models.AuctionNFTItem, int64, error) {
	var nfts []models.AuctionNFTItem
	var total int64

	// 先统计去重后的 NFT 总数
	// 使用 DISTINCT 去重 nft_address 和 token_id 的组合
	var countResult struct {
		Count int64
	}
	countSQL := `SELECT COUNT(DISTINCT CONCAT(nft_address, '-', token_id)) as count FROM auctions`
	if err := database.DB.Raw(countSQL).Scan(&countResult).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count NFTs: %w", err)
	}
	total = countResult.Count

	// 查询去重后的 NFT 列表，并统计每个 NFT 的拍卖数量
	// 使用 GROUP BY 去重，并获取每个 NFT 的最新信息
	querySQL := `
		SELECT 
			nft_address as nft_address,
			token_id as token_id,
			MAX(nft_id) as nft_id,
			MAX(name) as name,
			MAX(image) as image,
			MAX(contract_name) as contract_name,
			MAX(contract_symbol) as contract_symbol,
			MAX(token_uri) as token_uri,
			MAX(description) as description,
			COUNT(*) as auction_count
		FROM auctions
		GROUP BY nft_address, token_id
		ORDER BY MAX(created_at) DESC
		LIMIT ? OFFSET ?
	`

	if err := database.DB.Raw(querySQL, query.Limit(), query.Offset()).
		Scan(&nfts).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list NFTs: %w", err)
	}

	return nfts, total, nil
}

// ConvertToUSDFromTokenUnit 将代币最小单位金额转换为美元价值（公有函数）
// ethConfig: 以太坊配置指针（包含 ChainID 和 AuctionContractAddress）
// tokenAddress: 代币地址
// amountBigInt: 代币数量（最小单位，例如 wei）
// client: 可选参数，以太坊客户端（如果为 nil 或不提供则根据 ethConfig 创建新客户端）
func ConvertToUSDFromTokenUnit(ethConfig *config.EthereumConfig, tokenAddress string, amountBigInt *big.Int, client ...*ethclient.Client) (*models.ConvertToUSDResponse, error) {
	// 处理可选的 client 参数
	var ethClient *ethclient.Client
	var ethClientWrapper *ethereum.Client
	var err error

	if len(client) > 0 && client[0] != nil {
		// 使用传入的 client
		ethClient = client[0]
	} else {
		// 根据配置创建新的客户端
		ethClientWrapper, err = ethereum.NewClient(*ethConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
		}
		ethClient = ethClientWrapper.GetClient()
		defer ethClientWrapper.Close()
	}

	// 检查金额是否大于0
	if amountBigInt == nil || amountBigInt.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	// 获取支付代币精度（用于计算原始金额）
	_, _, tokenDecimals, _, err := utils.ERC20Token(ethClient, tokenAddress, ethConfig.ChainID)
	if err != nil {
		logger.Error("failed to get ERC20 token: %s", err.Error())
		return nil, fmt.Errorf("failed to get ERC20 token: %w", err)
	}

	// 创建拍卖合约实例
	auctionContract, err := my_auction.NewMyXAuctionV2(common.HexToAddress(ethConfig.AuctionContractAddress), ethClient)
	if err != nil {
		logger.Error("failed to create auction contract: %s", err.Error())
		return nil, fmt.Errorf("failed to create auction contract: %w", err)
	}

	tokenAddressAddress := common.HexToAddress(tokenAddress)

	// 获取代币价格（用于获取价格精度）
	price, err := auctionContract.GetTokenPrice(&bind.CallOpts{}, tokenAddressAddress)
	if err != nil {
		logger.Error("failed to get token price: %s", err.Error())
		return nil, fmt.Errorf("failed to get token price: %w", err)
	}

	// 调用合约的 ConvertToUSDValue 方法
	usdValue, err := auctionContract.ConvertToUSDValue(
		&bind.CallOpts{},
		tokenAddressAddress,
		amountBigInt,
	)
	if err != nil {
		logger.Error("failed to convert to USD value: %s", err.Error())
		return nil, fmt.Errorf("failed to convert to USD value: %w", err)
	}

	// 将最小单位金额转换为原始金额（考虑代币精度）
	// 例如：1000000000000000000 wei = 1 ETH (精度18)
	amountDecimal := decimal.NewFromBigInt(amountBigInt, int32(tokenDecimals)*-1)
	amountFloat := amountDecimal.InexactFloat64()

	// 将 USD 值从价格精度格式转换为实际美元金额
	// Chainlink 返回的是 8 位小数格式，需要除以 10^price.Decimals
	usdValueDecimal := decimal.NewFromBigInt(usdValue, int32(price.Decimals)*-1)

	return &models.ConvertToUSDResponse{
		Token:         tokenAddress,
		Amount:        amountFloat,
		AmountUSD:     usdValueDecimal.InexactFloat64(),
		AmountUnitUSD: usdValue.Uint64(),
		AmountUSDStr:  usdValueDecimal.StringFixed(int32(price.Decimals)), // 保留价格精度位小数
	}, nil
}

// ConvertTokenAmountToUSD 将代币金额转换为美元价值（公有函数）
// ethConfig: 以太坊配置指针（包含 ChainID 和 AuctionContractAddress）
// tokenAddress: 代币地址
// amount: 代币数量（float64）
// client: 可选参数，以太坊客户端（如果为 nil 或不提供则根据 ethConfig 创建新客户端）
func ConvertTokenAmountToUSD(ethConfig *config.EthereumConfig, tokenAddress string, amount float64, client ...*ethclient.Client) (*models.ConvertToUSDResponse, error) {
	// 处理可选的 client 参数
	var ethClient *ethclient.Client
	var ethClientWrapper *ethereum.Client
	var err error

	if len(client) > 0 && client[0] != nil {
		// 使用传入的 client
		ethClient = client[0]
	} else {
		// 根据配置创建新的客户端
		ethClientWrapper, err = ethereum.NewClient(*ethConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
		}
		ethClient = ethClientWrapper.GetClient()
		defer ethClientWrapper.Close()
	}

	// 将 float64 转换为 decimal
	amountDecimal := decimal.NewFromFloat(amount)
	if amountDecimal.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	// 获取支付代币精度
	_, _, tokenDecimals, _, err := utils.ERC20Token(ethClient, tokenAddress, ethConfig.ChainID)
	if err != nil {
		logger.Error("failed to get ERC20 token: %s", err.Error())
		return nil, fmt.Errorf("failed to get ERC20 token: %w", err)
	}

	// 将金额转换为最小单位（考虑精度）
	amountBigInt := amountDecimal.Mul(decimal.NewFromInt(int64(math.Pow10(int(tokenDecimals))))).BigInt()

	// 创建拍卖合约实例
	auctionContract, err := my_auction.NewMyXAuctionV2(common.HexToAddress(ethConfig.AuctionContractAddress), ethClient)
	if err != nil {
		logger.Error("failed to create auction contract: %s", err.Error())
		return nil, fmt.Errorf("failed to create auction contract: %w", err)
	}

	tokenAddressAddress := common.HexToAddress(tokenAddress)

	// 获取代币价格（用于获取价格精度）
	price, err := auctionContract.GetTokenPrice(&bind.CallOpts{}, tokenAddressAddress)
	if err != nil {
		logger.Error("failed to get token price: %s", err.Error())
		return nil, fmt.Errorf("failed to get token price: %w", err)
	}

	// 调用合约的 ConvertToUSDValue 方法
	usdValue, err := auctionContract.ConvertToUSDValue(
		&bind.CallOpts{},
		tokenAddressAddress,
		amountBigInt,
	)
	if err != nil {
		logger.Error("failed to convert to USD value: %s", err.Error())
		return nil, fmt.Errorf("failed to convert to USD value: %w", err)
	}

	// 将最小单位金额转换为原始金额（考虑代币精度）
	// 例如：1000000000000000000 wei = 1 ETH (精度18)
	amountFloat := amountDecimal.InexactFloat64()

	// 将 USD 值从价格精度格式转换为实际美元金额
	// Chainlink 返回的是 8 位小数格式，需要除以 10^price.Decimals
	usdValueDecimal := decimal.NewFromBigInt(usdValue, int32(price.Decimals)*-1)

	return &models.ConvertToUSDResponse{
		Token:         tokenAddress,
		Amount:        amountFloat,
		AmountUSD:     usdValueDecimal.InexactFloat64(),
		AmountUnitUSD: usdValue.Uint64(),
		AmountUSDStr:  usdValueDecimal.StringFixed(int32(price.Decimals)), // 保留价格精度位小数
	}, nil
}

// GetTokenPrice 获取代币价格（公有函数）
// ethConfig: 以太坊配置指针（包含 ChainID 和 AuctionContractAddress）
// tokenAddress: 代币地址
// client: 可选参数，以太坊客户端（如果为 nil 或不提供则根据 ethConfig 创建新客户端）
func GetTokenPrice(ethConfig *config.EthereumConfig, tokenAddress string, client ...*ethclient.Client) (*models.TokenPriceResponse, error) {
	// 处理可选的 client 参数
	var ethClient *ethclient.Client
	var ethClientWrapper *ethereum.Client
	var err error

	if len(client) > 0 && client[0] != nil {
		// 使用传入的 client
		ethClient = client[0]
	} else {
		// 根据配置创建新的客户端
		ethClientWrapper, err = ethereum.NewClient(*ethConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
		}
		ethClient = ethClientWrapper.GetClient()
		defer ethClientWrapper.Close()
	}

	auctionContract, err := my_auction.NewMyXAuctionV2(common.HexToAddress(ethConfig.AuctionContractAddress), ethClient)
	if err != nil {
		logger.Error("failed to create auction contract: %s", err.Error())
		return nil, fmt.Errorf("failed to create auction contract: %w", err)
	}

	tokenAddressAddress := common.HexToAddress(tokenAddress)

	// 获取代币价格（从 Chainlink price feed）
	price, err := auctionContract.GetTokenPrice(&bind.CallOpts{}, tokenAddressAddress)
	if err != nil {
		logger.Error("failed to get token price: %s", err.Error())
		return nil, fmt.Errorf("failed to get token price: %w", err)
	}

	// 将价格从 8 位小数格式转换为实际美元金额
	// Chainlink 返回的是 8 位小数格式，需要除以 1e8
	priceDecimal := decimal.NewFromBigInt(price.Price, int32(price.Decimals)*-1)

	return &models.TokenPriceResponse{
		Token:        tokenAddress,
		Price:        priceDecimal.InexactFloat64(),
		PriceUnitUSD: price.Price.Uint64(),
		PriceUSDStr:  priceDecimal.StringFixed(int32(price.Decimals)), // 保留8位小数
		Decimals:     price.Decimals,
	}, nil
}

// ConvertUnitUSDToToken 将整数美元金额（使用价格精度格式）转换为代币金额（公有函数）
// ethConfig: 以太坊配置指针（包含 ChainID 和 AuctionContractAddress）
// tokenAddress: 代币地址
// amountUSDUnit: 美元金额（uint64，使用价格精度格式，精度从链上获取）
// client: 可选参数，以太坊客户端（如果为 nil 或不提供则根据 ethConfig 创建新客户端）
func ConvertUnitUSDToToken(ethConfig *config.EthereumConfig, tokenAddress string, amountUSDUnit uint64, client ...*ethclient.Client) (*models.ConvertUSDToTokenResponse, error) {
	// 处理可选的 client 参数
	var ethClient *ethclient.Client
	var ethClientWrapper *ethereum.Client
	var err error

	if len(client) > 0 && client[0] != nil {
		// 使用传入的 client
		ethClient = client[0]
	} else {
		// 根据配置创建新的客户端
		ethClientWrapper, err = ethereum.NewClient(*ethConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create Ethereum client: %w", err)
		}
		ethClient = ethClientWrapper.GetClient()
		defer ethClientWrapper.Close()
	}

	// 步骤1: 获取代币价格（从链上获取真实的价格和精度）
	// GetTokenPrice 返回：代币价格（1个代币 = X USD）和价格的精度（Decimals）
	priceResponse, err := GetTokenPrice(ethConfig, tokenAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get token price: %w", err)
	}

	// 步骤2: 获取代币精度（用于代币金额显示）
	_, _, tokenDecimals, _, err := utils.ERC20Token(ethClient, tokenAddress, ethConfig.ChainID)
	if err != nil {
		logger.Error("failed to get ERC20 token decimals: %s", err.Error())
		return nil, fmt.Errorf("failed to get ERC20 token decimals: %w", err)
	}

	// 步骤3: 将 USD 金额从链上精度格式转换为实际美元金额（用于返回）
	// amountUSDUnit 的精度应该和价格精度一致（priceResponse.Decimals，从 Chainlink oracle 获取）
	priceDecimals := priceResponse.Decimals
	divisor := decimal.NewFromInt(int64(math.Pow10(int(priceDecimals))))
	amountUSDDecimal := decimal.NewFromInt(int64(amountUSDUnit)).Div(divisor)
	amountUSDFloat := amountUSDDecimal.InexactFloat64()

	// 步骤4: 根据 USD 金额和代币价格计算需要多少代币
	// 使用整数价格（PriceUnitUSD）和整数美元金额（amountUSDUnit）进行整数运算，避免精度丢失
	// 由于 amountUSDUnit 和 priceResponse.PriceUnitUSD 的精度相同（都是 priceDecimals），
	// 可以直接相除得到代币数量（小数形式）
	if priceResponse.PriceUnitUSD == 0 {
		return nil, fmt.Errorf("token price must be greater than 0")
	}

	// 使用整数运算：amountUSDUnit / priceResponse.PriceUnitUSD
	// 例如：amountUSDUnit = 298228000000 (表示 $2982.28，精度8)
	//      priceResponse.PriceUnitUSD = 298228000000 (表示 $2982.28/代币，精度8)
	//      结果 = 1.0 代币
	amountUSDDecimalBig := decimal.NewFromInt(int64(amountUSDUnit))
	priceUnitUSDDecimal := decimal.NewFromInt(int64(priceResponse.PriceUnitUSD))
	tokenAmountDecimal := amountUSDDecimalBig.Div(priceUnitUSDDecimal)
	tokenAmountFloat := tokenAmountDecimal.InexactFloat64()

	return &models.ConvertUSDToTokenResponse{
		TokenAddress:  tokenAddress,
		USDAmount:     amountUSDFloat,
		TokenAmount:   tokenAmountFloat,
		TokenDecimals: tokenDecimals,
		USDUnitAmount: amountUSDUnit,
	}, nil
}

// GetByID 根据AuctionID（字符串）获取单个拍卖记录（基本信息）
func (s *AuctionService) GetByID(auctionID string) (*models.Auction, error) {
	var auction models.Auction
	if err := database.DB.Where("auction_id = ?", auctionID).First(&auction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound.WithMessage("auction not found")
		}
		return nil, fmt.Errorf("failed to get auction: %w", err)
	}
	return &auction, nil
}

// GetDetailByID 获取拍卖详情（只返回钱包地址，不返回完整User信息）
func (s *AuctionService) GetDetailByID(id uint64) (*models.AuctionDetailResponse, error) {
	var result struct {
		models.Auction
		SellerWalletAddress string `gorm:"column:seller_wallet_address"`
	}

	if err := database.DB.Table("auctions").
		Select("auctions.*, users.wallet_address as seller_wallet_address").
		Joins("LEFT JOIN users ON auctions.user_id = users.id").
		Where("auctions.id = ?", id).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound.WithMessage("auction not found")
		}
		return nil, fmt.Errorf("failed to get auction detail: %w", err)
	}

	return &models.AuctionDetailResponse{
		Auction:             result.Auction,
		SellerWalletAddress: result.SellerWalletAddress,
	}, nil
}

func (s *AuctionService) GetByContractID(contractAuctionID uint64) (*models.Auction, error) {
	var auction models.Auction
	if err := database.DB.Where("contract_auction_id = ?", contractAuctionID).
		Preload("User").
		Preload("Bids").
		Preload("Bids.User").
		First(&auction).Error; err != nil {
		return nil, fmt.Errorf("auction not found: %w", err)
	}
	return &auction, nil
}

// Create 创建拍卖
// 流程：
// 1. 验证NFT所有权（数据库）
// 2. 链上验证NFT授权状态
// 3. 验证时间参数
// 4. 检查NFT是否已有在线拍卖（防止重复上架）
// 5. 计算起拍价USD价值
// 6. 生成拍卖ID和OnlineLock
// 7. 构建并保存拍卖记录
func (s *AuctionService) Create(userID uint64, payload models.AuctionPayload) (*models.Auction, error) {
	// ========== 步骤1: 验证NFT所有权 ==========
	// 通过 nft_ownerships 表验证用户是否拥有该NFT（状态为 holding）
	var ownership models.NFTOwnership
	if err := database.DB.Where("nft_id = ? AND user_id = ? AND status = ?",
		payload.NFTID, userID, models.NFTOwnershipStatusHolding).
		Preload("NFT").
		First(&ownership).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(
				fmt.Sprintf("NFT not found or you don't own it: NFT ID %s. Please sync your NFTs first", payload.NFTID))
		}
		return nil, fmt.Errorf("failed to verify NFT ownership: %w", err)
	}

	// 从关联的 NFT 获取元数据
	nft := ownership.NFT
	if nft.NFTID == "" {
		return nil, fmt.Errorf("NFT data not found for NFT ID: %s", payload.NFTID)
	}
	// ========== 步骤4: 检查NFT是否已有在线拍卖 ==========
	// OnlineLock 格式: nft_id:1，用于锁定NFT的唯一性，防止同一NFT同时存在多个在线拍卖
	onlineLock := fmt.Sprintf("%s:1", payload.NFTID)
	var existingAuction models.Auction
	if err := database.DB.Where("online_lock = ?", onlineLock).First(&existingAuction).Error; err == nil {
		return nil, errors.BadRequest(
			fmt.Sprintf("该 NFT 已经在拍卖中，拍卖ID: %s", existingAuction.AuctionID))
	} else if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to check existing auction: %w", err)
	}
	// ========== 步骤2: 链上验证NFT所有权和授权状态 ==========
	// 获取当前用户的钱包地址
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// 创建NFT合约实例
	platformContractAddress := common.HexToAddress(s.config.AuctionContractAddress)
	nftContractAddress := common.HexToAddress(payload.NFTAddress)

	_ethClient := s.ethClient.GetClient()
	nftContract, err := erc721_nft.NewMyNFT(nftContractAddress, _ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create NFT contract instance: %w", err)
	}

	opts := &bind.CallOpts{Context: context.Background()}
	tokenIDBigInt := big.NewInt(int64(payload.TokenID))

	// 检查链上NFT的实际owner是否是当前用户
	chainOwner, err := nftContract.OwnerOf(opts, tokenIDBigInt)
	if err != nil {
		return nil, fmt.Errorf("failed to get NFT owner for token %d: %w", payload.TokenID, err)
	}

	// 规范化地址（转为小写）进行比较
	normalizedChainOwner := strings.ToLower(chainOwner.Hex())
	// 验证链上owner是否是当前用户
	if normalizedChainOwner != user.WalletAddress {
		return nil, errors.BadRequest(
			fmt.Sprintf("NFT token %d is not owned by your wallet address. Chain owner: %s, Your wallet: %s. NFT ID: %s",
				payload.TokenID, chainOwner.Hex(), user.WalletAddress, payload.NFTID))
	}

	// 检查指定TokenID的NFT是否已授权给平台合约
	approvedAddress, err := nftContract.GetApproved(opts, tokenIDBigInt)
	if err != nil {
		return nil, fmt.Errorf("failed to check GetApproved for token %d: %w", payload.TokenID, err)
	}

	if approvedAddress != platformContractAddress {
		return nil, errors.BadRequest(
			fmt.Sprintf("NFT token %d has not been approved for the platform contract. Please approve the NFT first. NFT ID: %s, Contract: %s",
				payload.TokenID, payload.NFTID, payload.NFTAddress))
	}

	// ========== 步骤3: 验证时间参数 ==========
	if payload.StartTime == nil || payload.EndTime == nil {
		return nil, errors.BadRequest("start time and end time are required")
	}

	// 验证结束时间必须大于开始时间
	if payload.EndTime.Before(*payload.StartTime) || payload.EndTime.Equal(*payload.StartTime) {
		return nil, errors.BadRequest(
			fmt.Sprintf("end time must be after start time, start time: %s, end time: %s",
				payload.StartTime.Format("2006-01-02 15:04:05"),
				payload.EndTime.Format("2006-01-02 15:04:05")))
	}

	// 验证拍卖持续时间至少1分钟
	minDuration := 1 * time.Minute
	if payload.EndTime.Sub(*payload.StartTime) < minDuration {
		return nil, errors.BadRequest(
			fmt.Sprintf("auction duration must be at least 1 minute, current duration: %v",
				payload.EndTime.Sub(*payload.StartTime)))
	}

	// ========== 步骤5: 计算起拍价USD价值 ==========
	startPriceFloat, _ := payload.StartPrice.Float64()
	usdResponse, err := ConvertTokenAmountToUSD(&s.config, payload.PaymentToken, startPriceFloat, _ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to convert start price to USD: %w", err)
	}
	startPriceUSD := decimal.NewFromFloat(usdResponse.AmountUSD)
	startPriceUnitUSD := usdResponse.AmountUnitUSD

	// ========== 步骤6: 生成拍卖ID和时间戳 ==========
	// 使用 snowflake 算法生成唯一拍卖ID
	auctionID := GenerateID()

	// 计算时间戳（Unix 时间戳，秒）
	startTimestamp := uint64(payload.StartTime.Unix())
	endTimestamp := uint64(payload.EndTime.Unix())
	now := time.Now()
	onlineTimestamp := uint64(now.Unix()) // 用于 Online 字段，表示创建时间

	// ========== 步骤7: 构建并保存拍卖记录 ==========
	auction := models.Auction{
		// 基本信息
		AuctionID:         auctionID,
		UserID:            userID,
		Status:            AuctionStatusPending, // 新创建的拍卖默认为待上架状态
		ContractAuctionID: 0,                    // 将在链上创建后设置

		// NFT基本信息（从关联的NFT获取）
		NFTID:          nft.NFTID,
		NFTAddress:     payload.NFTAddress,
		TokenID:        payload.TokenID,
		OnlineLock:     onlineLock,      // 格式: nft_id:1，用于锁定NFT唯一性
		Online:         onlineTimestamp, // 创建时间戳（表示未上线，上架后会改为1）
		TokenURI:       nft.TokenURI,
		ContractName:   nft.ContractName,
		ContractSymbol: nft.ContractSymbol,
		NftName:        nft.NftName,
		OwnerAddress:   nft.NftOwnerAddress,
		Image:          nft.Image,
		Description:    nft.Description,
		Metadata:       nft.Metadata,

		// 拍卖信息
		PaymentToken:      payload.PaymentToken,
		StartPrice:        &payload.StartPrice,
		StartPriceUSD:     &startPriceUSD,
		StartPriceUnitUSD: startPriceUnitUSD,
		StartTime:         payload.StartTime,
		EndTime:           payload.EndTime,
		StartTimestamp:    startTimestamp,
		EndTimestamp:      endTimestamp,

		// 出价信息（初始值）
		HighestBid:    &decimal.Zero,
		HighestBidUSD: &decimal.Zero,
		BidCount:      0,
	}

	// 保存到数据库
	if err := database.DB.Create(&auction).Error; err != nil {
		return nil, fmt.Errorf("failed to create auction: %w", err)
	}
	return &auction, nil
}

// GenerateID 使用 snowflake 算法生成唯一 ID（返回字符串格式）
func GenerateID() string {
	if snowflakeGenerator == nil {
		logger.Error("snowflake generator is not initialized, using timestamp as fallback")
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	id, err := snowflakeGenerator.NextID()
	if err != nil {
		logger.Error("failed to generate snowflake ID: %s, using timestamp as fallback", err.Error())
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	return fmt.Sprintf("%d", id)
}

func (s *AuctionService) UpdateContractID(id uint64, contractAuctionID uint64) error {
	if err := database.DB.Model(&models.Auction{}).
		Where("id = ?", id).
		Update("contract_auction_id", contractAuctionID).Error; err != nil {
		return fmt.Errorf("failed to update auction contract id: %w", err)
	}
	return nil
}

func (s *AuctionService) UpdateStatus(id uint64, status string) error {
	if err := database.DB.Model(&models.Auction{}).
		Where("id = ?", id).
		Update("status", status).Error; err != nil {
		return fmt.Errorf("failed to update auction status: %w", err)
	}
	return nil
}

func (s *AuctionService) UpdateHighestBid(id uint64, highestBidPaymentToken string, highestBid decimal.Decimal, highestBidUSD decimal.Decimal, highestBidder string) error {
	updates := map[string]interface{}{
		"highest_bid_payment_token": highestBidPaymentToken,
		"highest_bid":               highestBid,
		"highest_bid_usd":           highestBidUSD,
		"highest_bidder":            highestBidder,
	}

	if err := database.DB.Model(&models.Auction{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update highest bid: %w", err)
	}

	// Increment bid count
	if err := database.DB.Model(&models.Auction{}).
		Where("id = ?", id).
		Update("bid_count", gorm.Expr("bid_count + 1")).Error; err != nil {
		return fmt.Errorf("failed to increment bid count: %w", err)
	}

	return nil
}

func (s *AuctionService) GetUserAuctions(userID uint64, query page.PageQuery, statuses []string) ([]models.Auction, int64, error) {
	var auctions []models.Auction
	var total int64

	// 构建查询条件
	db := database.DB.Model(&models.Auction{}).Where("user_id = ?", userID)

	// 如果指定了状态过滤，添加状态条件
	if len(statuses) > 0 {
		db = db.Where("status IN ?", statuses)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询数据
	queryDB := database.DB.Where("user_id = ?", userID)
	if len(statuses) > 0 {
		queryDB = queryDB.Where("status IN ?", statuses)
	}

	if err := queryDB.
		Preload("User").
		Order("created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&auctions).Error; err != nil {
		return nil, 0, err
	}

	return auctions, total, nil
}

// GetUserAuctionHistory 获取用户拍卖历史记录（精简字段）
func (s *AuctionService) GetUserAuctionHistory(userID uint64, query page.PageQuery, statuses []string) ([]models.AuctionHistoryResponse, int64, error) {
	var auctions []models.Auction
	var total int64

	// 构建查询条件
	db := database.DB.Model(&models.Auction{}).Where("user_id = ?", userID)

	// 如果指定了状态过滤，添加状态条件
	if len(statuses) > 0 {
		db = db.Where("status IN ?", statuses)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询数据（只选择需要的字段）
	queryDB := database.DB.Where("user_id = ?", userID)
	if len(statuses) > 0 {
		queryDB = queryDB.Where("status IN ?", statuses)
	}

	if err := queryDB.
		Select("auction_id, nft_id, nft_address, token_id, nft_name, image, contract_name, contract_symbol, status, payment_token, start_price, start_price_usd, end_time, end_timestamp, bid_count, highest_bid, highest_bid_usd, highest_bidder, highest_bid_payment_token").
		Order("created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&auctions).Error; err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	historyList := make([]models.AuctionHistoryResponse, 0, len(auctions))
	for _, auction := range auctions {
		history := models.AuctionHistoryResponse{
			AuctionID:      auction.AuctionID,
			NFTID:          auction.NFTID,
			NFTAddress:     auction.NFTAddress,
			TokenID:        auction.TokenID,
			NftName:        auction.NftName,
			Image:          auction.Image,
			ContractName:   auction.ContractName,
			ContractSymbol: auction.ContractSymbol,
			Status:         auction.Status,
			PaymentToken:   auction.PaymentToken,  // 支付代币地址
			FloorPrice:     auction.StartPrice,    // 地板价格使用起拍价
			FloorPriceUSD:  auction.StartPriceUSD, // 地板价格USD
			EndTime:        auction.EndTime,
			EndTimestamp:   auction.EndTimestamp,
			BidCount:       auction.BidCount,
		}

		// 如果有出价，才显示出价信息
		if auction.BidCount > 0 && auction.HighestBidder != "" {
			history.HighestBid = auction.HighestBid
			history.HighestBidUSD = auction.HighestBidUSD
			history.HighestBidder = auction.HighestBidder
			history.HighestBidPaymentToken = auction.HighestBidPaymentToken
		}

		historyList = append(historyList, history)
	}

	return historyList, total, nil
}

// Update 更新拍卖信息（只能更新待上架状态的拍卖）
func (s *AuctionService) Update(userID uint64, auctionID string, payload models.UpdateAuctionPayload) (*models.Auction, error) {
	// 先查询拍卖是否存在且属于当前用户
	var auction models.Auction
	if err := database.DB.Where("auction_id = ? AND user_id = ?", auctionID, userID).First(&auction).Error; err != nil {
		return nil, errors.Forbidden("auction not found or access denied")
	}

	// 只有待上架（pending）状态的拍卖可以更新
	if auction.Status != AuctionStatusPending {
		return nil, errors.Forbidden("cannot update auction: only pending auctions can be updated")
	}

	// 计算起始价的 USD 价值
	startPriceFloat, _ := payload.StartPrice.Float64()
	usdResponse, err := ConvertTokenAmountToUSD(&s.config, payload.PaymentToken, startPriceFloat, s.ethClient.GetClient())
	if err != nil {
		logger.Error("failed to convert start price to USD: %s", err.Error())
		// 如果转换失败，使用零值，但不阻止更新
		return nil, fmt.Errorf("failed to convert start price to USD: %w", err)
	}
	startPriceUSD := usdResponse.AmountUSD
	startPriceUnitUSD := usdResponse.AmountUnitUSD

	// 计算时间戳（Unix 时间戳，秒）
	startTimestamp := uint64(payload.StartTime.Unix())
	endTimestamp := uint64(payload.EndTime.Unix())

	// 更新拍卖信息
	updates := map[string]interface{}{
		"payment_token":        payload.PaymentToken,
		"start_price":          payload.StartPrice,
		"start_price_usd":      startPriceUSD,
		"start_price_unit_usd": startPriceUnitUSD,
		"start_time":           payload.StartTime,
		"end_time":             payload.EndTime,
		"start_timestamp":      startTimestamp,
		"end_timestamp":        endTimestamp,
	}

	if err := database.DB.Model(&auction).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update auction: %w", err)
	}

	// 重新加载更新后的数据
	if err := database.DB.Preload("User").Where("auction_id = ?", auctionID).First(&auction).Error; err != nil {
		return nil, fmt.Errorf("failed to reload auction: %w", err)
	}

	return &auction, nil
}

// Cancel 取消拍卖
// TODO: 实现具体的取消逻辑
func (s *AuctionService) Cancel(userID uint64, auctionID string) (*models.Auction, error) {
	// 先查询拍卖是否存在且属于当前用户
	var auction models.Auction
	if err := database.DB.Where("auction_id = ? AND user_id = ?", auctionID, userID).First(&auction).Error; err != nil {
		return nil, errors.Forbidden("auction not found or access denied")
	}
	if auction.Status == AuctionStatusPending { //直接执行取消
		if err := database.DB.Model(&auction).Updates(map[string]interface{}{
			"online_lock": fmt.Sprintf("%s:%d", auction.NFTID, auction.TokenID), //解锁nft唯一性锁定
			"status":      AuctionStatusCancelled,
		}).Error; err != nil {
			//合约执行取消
			client := s.ethClient.GetClient()
			myAuction, err := my_auction.NewMyXAuctionV2(common.HexToAddress(s.config.AuctionContractAddress), client)
			if err != nil {
				logger.Error("failed to create my auction: %v", err)
				return nil, fmt.Errorf("failed to create my auction: %w", err)
			}
			ctx := context.Background()
			auth, err := s.ethClient.GetAuth(ctx, s.config.PlatformPrivateKey)
			if err != nil {
				logger.Error("failed to get auth: %v", err)
				return nil, fmt.Errorf("failed to execute cancel auction")
			}
			if tx, err := myAuction.CancelAuction(auth, big.NewInt(int64(auction.ContractAuctionID))); err != nil {
				logger.Error("failed to wait for transaction: %v", err)
				return nil, fmt.Errorf("failed to wait for transaction")
			} else {
				logger.Info("cancel auction transaction sent: %v", tx.Hash())
			}
		}
	}
	return &auction, nil
}

// Publish 上架拍卖（将状态从 pending 改为 active）
func (s *AuctionService) Publish(userID uint64, auctionID uint64) (*models.Auction, error) {
	// 先查询拍卖是否存在且属于当前用户
	var auction models.Auction
	if err := database.DB.Where("id = ? AND user_id = ?", auctionID, userID).First(&auction).Error; err != nil {
		return nil, errors.Forbidden("auction not found or access denied")
	}

	// 只有待上架（pending）状态的拍卖可以上架
	if auction.Status != AuctionStatusPending {
		return nil, errors.Forbidden("cannot publish auction: only pending auctions can be published")
	}

	// 校验拍卖时间
	now := time.Now()

	// 开始时间可以小于当前时间（允许已经开始）
	// 结束时间必须大于当前时间（拍卖必须还未结束）
	if auction.EndTime.Before(now) || auction.EndTime.Equal(now) {
		return nil, errors.BadRequest(
			fmt.Sprintf("end time must be in the future, current end time: %s, now: %s",
				auction.EndTime.Format("2006-01-02 15:04:05"),
				now.Format("2006-01-02 15:04:05")))
	}

	// 检查结束时间是否在开始时间之后
	if auction.StartTime == nil || auction.EndTime == nil {
		return nil, errors.BadRequest("start time and end time are required")
	}
	if auction.EndTime.Before(*auction.StartTime) || auction.EndTime.Equal(*auction.StartTime) {
		return nil, errors.BadRequest(
			fmt.Sprintf("end time must be after start time, start time: %s, end time: %s",
				auction.StartTime.Format("2006-01-02 15:04:05"),
				auction.EndTime.Format("2006-01-02 15:04:05")))
	}

	// 检查拍卖持续时间是否合理（至少1分钟）
	minDuration := 1 * time.Minute
	if auction.EndTime.Sub(*auction.StartTime) < minDuration {
		return nil, errors.BadRequest(
			fmt.Sprintf("auction duration must be at least 1 minute, current duration: %v",
				auction.EndTime.Sub(*auction.StartTime)))
	}

	// 不修改 status（保持为 pending），只更新 nft_online_id 为 nft_id:1（表示已上线），online 为 1
	// pending 状态表示准备开卖了
	nftOnlineID := fmt.Sprintf("%s:1", auction.NFTID)
	online := uint64(1)
	updates := map[string]interface{}{
		"nft_online_id": nftOnlineID,
		"online":        online,
	}
	if err := database.DB.Model(&auction).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to publish auction: %w", err)
	}

	// 重新加载更新后的数据
	if err := database.DB.Preload("User").First(&auction, auctionID).Error; err != nil {
		return nil, fmt.Errorf("failed to reload auction: %w", err)
	}

	return &auction, nil
}

// GetSupportedTokens 获取平台支持的代币列表（根据当前网络配置）
func (s *AuctionService) GetSupportedTokens() ([]map[string]interface{}, error) {
	// 获取当前网络的 USDC 地址
	usdcAddress, err := utils.GetUSDCAddress(s.config.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get USDC address for chain ID %d: %w", s.config.ChainID, err)
	}

	tokens := []map[string]interface{}{
		{
			"address": utils.ETHAddress,
			"symbol":  "ETH",
			"name":    "Ethereum",
			"default": true,
		},
		{
			"address": usdcAddress,
			"symbol":  "USDC",
			"name":    "USD Coin",
			"default": false,
		},
	}

	return tokens, nil
}

func (s *AuctionService) OnEventAuctionCreated(auctionContractId uint64, ownerAddress string, nftAddress string, tokenId uint64) error {
	// 开启事务
	// 先查询拍卖是否存在且属于当前用户

	nftID := GenerateNFTID(nftAddress, tokenId)
	nftOnlineLock := fmt.Sprintf("%s:1", nftID)
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var auction models.Auction
		if err := tx.Where("online_lock = ? and status = ?", nftOnlineLock, AuctionStatusPending).First(&auction).Error; err != nil {
			logger.Error("failed to get auction: %v", err)
			return fmt.Errorf("failed to get auction: %w", err)
		}
		// 更新拍卖状态为 active
		// 更新拍卖的合同拍卖ID
		if err := tx.Model(&auction).
			Updates(map[string]interface{}{
				"status":              AuctionStatusActive,
				"contract_auction_id": auctionContractId,
				"owner_address":       ownerAddress, //拍卖合约地址
				"online":              1,            //上线
			}).Error; err != nil {
			logger.Error("failed to update auction status: %v", err)
			return fmt.Errorf("failed to update auction status: %w", err)
		}
		//更新nft_ownerships
		if err := tx.Model(&models.NFTOwnership{}).
			Where("nft_id = ? and user_id = ?", nftID, auction.UserID).
			Updates(map[string]interface{}{
				"owner_address": ownerAddress,                     //委托给拍卖合约	ownerAddress
				"status":        models.NFTOwnershipStatusSelling, // 在售
			}).Error; err != nil {
			logger.Error("failed to update nft ownership: %v", err)
			return fmt.Errorf("failed to update nft ownership: %w", err)
		}
		return nil
	})

	if err == nil {
		// 获取最新拍卖信息开始调度拍卖结束任务
		// 如果 end_time 改变了，重新调度任务
		var auction models.Auction
		if err := database.DB.Where("online_lock = ? and status = ?", nftOnlineLock, AuctionStatusActive).First(&auction).Error; err != nil {
			logger.Error("failed to get auction: %v", err)
			return fmt.Errorf("failed to get auction: %w", err)
		} else {
			logger.Info("Auction end task rescheduled: auctionID=%s, endTime=%v", auction.AuctionID, auction.EndTime)
			if s.taskScheduler != nil && auction.EndTime != nil {
				// 注意：Asynq 不支持直接取消延迟任务，但可以在 handler 中检查
				// 这里重新调度，如果旧任务执行时会检查状态
				if err := s.taskScheduler.ScheduleAuctionEndTask(
					&auction,
				); err != nil {
					logger.Error("Failed to reschedule auction end task: auctionID=%s, error=%v", auction.AuctionID, err)
				} else {
					logger.Info("Auction end task rescheduled: auctionID=%s, endTime=%v", auction.AuctionID, auction.EndTime)
				}
			}
			logger.Info("OnEventAuctionCreated: auction updated successfully, contractAuctionId=%d, nftAddress=%s, tokenId=%d", auctionContractId, nftAddress, tokenId)
		}
	}
	return err
}
