package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/contracts/my_auction"
	"my-auction-market-api/internal/database"
	ethclientwrapper "my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/page"
	"my-auction-market-api/internal/utils"
)

type BidService struct {
	config    config.EthereumConfig
	ethClient *ethclientwrapper.Client
}

func NewBidService(ethCfg config.EthereumConfig, ethClient *ethclientwrapper.Client) *BidService {
	return &BidService{
		config:    ethCfg,
		ethClient: ethClient,
	}
}

// ConvertBidToResponse 将 Bid 转换为 BidResponse（统一格式，用于 WebSocket 和 API）
// 这是一个公开方法，供其他服务（如 ListenerService）调用
func (s *BidService) ConvertBidToResponse(bid *models.Bid) models.BidResponse {
	// 准备出价金额（代币金额）- 使用数值类型
	var amount float64
	if bid.Amount != nil {
		amount = bid.Amount.InexactFloat64()
	}

	// 准备USD金额 - 使用数值类型
	var amountUSD float64
	if bid.AmountUSD != nil {
		amountUSD = bid.AmountUSD.InexactFloat64()
	}

	// 获取支付代币符号
	paymentTokenSymbol := "UNKNOWN"
	if s.ethClient != nil {
		_, symbol, _, _, err := utils.ERC20Token(s.ethClient.GetClient(), bid.PaymentToken, s.config.ChainID)
		if err != nil {
			logger.Warn("failed to get payment token symbol for %s: %v", bid.PaymentToken, err)
			// 如果获取失败，尝试根据地址判断
			if strings.EqualFold(bid.PaymentToken, utils.ETHAddress) {
				paymentTokenSymbol = "ETH"
			}
		} else {
			paymentTokenSymbol = symbol
		}
	}

	return models.BidResponse{
		ID:                 bid.ID,
		AuctionID:          bid.AuctionID,
		ContractAuctionID:  bid.ContractAuctionID,
		UserID:             bid.UserID,
		Bidder:             bid.WalletAddress,
		Amount:             amount,
		AmountUnit:         bid.AmountUnit,
		AmountUSD:          amountUSD,
		AmountUnitUSD:      bid.AmountUnitUSD,
		PaymentToken:       bid.PaymentToken,
		PaymentTokenSymbol: paymentTokenSymbol,
		TransactionHash:    bid.TransactionHash,
		BlockNumber:        bid.BlockNumber,
		Timestamp:          bid.Timestamp,
		BidCount:           bid.BidCount,
		IsHighest:          bid.IsHighest,
		MinBidder:          bid.MinBidder,
		MinBidUnitUSD:      bid.MinBidUnitUSD,
		CreatedAt:          bid.CreatedAt,
	}
}

// GetBidsByAuctionID 获取拍卖的出价列表，返回格式与 WebSocket 消息相同
func (s *BidService) GetBidsByAuctionID(auctionID string, query page.PageQuery) ([]models.BidResponse, int64, error) {
	var total int64
	var bids []models.Bid

	// 统计总数
	if err := database.DB.Model(&models.Bid{}).
		Where("auction_id = ?", auctionID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询出价记录
	if err := database.DB.
		Where("auction_id = ?", auctionID).
		Order("created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&bids).Error; err != nil {
		return nil, 0, err
	}

	// 转换为统一格式
	bidList := make([]models.BidResponse, 0, len(bids))
	for _, bid := range bids {
		bidList = append(bidList, s.ConvertBidToResponse(&bid))
	}

	return bidList, total, nil
}

func (s *BidService) GetByID(id uint64) (*models.Bid, error) {
	var bid models.Bid
	if err := database.DB.Preload("User").
		Preload("Auction").
		First(&bid, id).Error; err != nil {
		return nil, fmt.Errorf("bid not found: %w", err)
	}
	return &bid, nil
}

// GetBidsByTransactionHash 根据交易哈希获取出价记录列表（支持分页）
func (s *BidService) GetBidsByTransactionHash(transactionHash string, query page.PageQuery) ([]models.BidDetailResponse, int64, error) {
	// 规范化交易哈希（转为小写）
	transactionHash = strings.ToLower(transactionHash)

	var total int64
	var results []struct {
		models.Bid
		BidderWalletAddress string `gorm:"column:bidder_wallet_address"`
	}

	// 统计总数
	if err := database.DB.Model(&models.Bid{}).
		Where("transaction_hash = ?", transactionHash).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 使用 JOIN 一次性查询出价和钱包地址
	if err := database.DB.Table("bids").
		Select("bids.*, users.wallet_address as bidder_wallet_address").
		Joins("LEFT JOIN users ON bids.user_id = users.id").
		Where("bids.transaction_hash = ?", transactionHash).
		Order("bids.created_at DESC").
		Offset(query.Offset()).
		Limit(query.Limit()).
		Find(&results).Error; err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	bidDetails := make([]models.BidDetailResponse, 0, len(results))
	for _, result := range results {
		bidDetails = append(bidDetails, models.BidDetailResponse{
			Bid:                 result.Bid,
			BidderWalletAddress: result.BidderWalletAddress,
		})
	}

	return bidDetails, total, nil
}

// Close 关闭服务并释放资源
func (s *BidService) Close() error {
	if s.ethClient != nil {
		s.ethClient.Close()
	}
	return nil
}

// OnEventBidPlaced 处理出价事件，创建出价记录
func (s *BidService) OnEventBidPlaced(event *my_auction.MyXAuctionV2BidPlaced, log *types.Log) (*models.Bid, error) {
	contractAuctionId := event.AuctionId.Uint64()

	// 直接查询数据库获取拍卖信息
	var auction models.Auction
	if err := database.DB.Where("contract_auction_id = ?", contractAuctionId).
		First(&auction).Error; err != nil {
		logger.Warn("Failed to get auction by contract_auction_id=%d: %s", contractAuctionId, err.Error())
		return nil, fmt.Errorf("failed to get auction by contract_auction_id: %w", err)
	}

	auctionId := auction.AuctionID

	bidder := strings.ToLower(event.Bidder.Hex())

	// 根据出价者钱包地址查询用户（必须是已登录系统的用户才能出价）
	var user models.User
	if err := database.DB.Where("wallet_address = ?", bidder).First(&user).Error; err != nil {
		logger.Warn("bidder wallet address not found in system: wallet=%s, auctionId=%d", bidder, contractAuctionId)
		return nil, fmt.Errorf("bidder wallet address %s is not registered in the system", bidder)
	}
	userID := user.ID
	paymentToken := strings.ToLower(event.PaymentToken.Hex())
	bidCount := event.BidCount.Uint64()

	transactionHash := strings.ToLower(log.TxHash.Hex())
	blockNumber := log.BlockNumber
	timestamp := event.Timestamp.Uint64() // 链上时间
	createdAt := time.Now()

	// 获取出价金额的最小单位（wei、usdc最小单位等）
	amountUnit := event.Amount.Uint64()
	// 获取USD金额的最小单位（8位小数）
	amountUnitUSD := event.BidValue.Uint64()
	// 获取最低出价钱包地址 也是链上最高出价者
	minBidder := strings.ToLower(event.MinBidder.Hex())
	// 获取最低出价美元价值x8
	minBidValueUSD := event.MinBidValue.Uint64()

	// 转换为USD（使用独立的函数，不依赖 AuctionService）
	// 注意：这里应该传入代币最小单位 event.Amount，而不是 event.BidValue（USD单位）
	amountUSDResponse, err := ConvertToUSDFromTokenUnit(&s.config, paymentToken, event.Amount, s.ethClient.GetClient())
	if err != nil {
		logger.Error("failed to convert to USD: %v", err)
		return nil, fmt.Errorf("failed to convert to USD: %w", err)
	}

	// 转换为decimal类型
	amountToken := decimal.NewFromFloat(amountUSDResponse.Amount) // ETH,USDC（代币金额）
	// 使用转换后的USD金额，确保精度正确
	amountUSD := decimal.NewFromFloat(amountUSDResponse.AmountUSD)

	// 创建出价记录
	bid := models.Bid{
		AuctionID:         auctionId,
		ContractAuctionID: contractAuctionId,
		UserID:            userID,
		WalletAddress:     bidder,
		Winner:            false,         // 出价时无法确定是否为获胜者，默认为false
		Amount:            &amountToken,  // 出价金额(ETH,USDC)
		AmountUnit:        amountUnit,    // 出价金额(wei、usdc最小单位等)
		AmountUSD:         &amountUSD,    // 出价金额USD
		AmountUnitUSD:     amountUnitUSD, // 出价金额USD最小单位（8位小数）
		PaymentToken:      paymentToken,
		TransactionHash:   transactionHash,
		BlockNumber:       blockNumber,
		Timestamp:         timestamp,
		BidCount:          bidCount,            // 出价总数
		IsHighest:         minBidder == bidder, // 是否为最高出价
		MinBidder:         minBidder,
		MinBidUnitUSD:     minBidValueUSD,
		CreatedAt:         &createdAt,
	}
	// 保存到数据库
	if err := database.DB.Create(&bid).Error; err != nil {
		logger.Error("failed to create bid record: %v", err)
		return nil, fmt.Errorf("failed to create bid record: %w", err)
	}

	// 更新拍卖的 bid_count 字段
	if err := database.DB.Model(&auction).
		Where("auction_id = ?", auctionId).
		Updates(map[string]interface{}{"bid_count": bidCount,
			"updated_at":                time.Now(),
			"highest_bidder":            bidder,
			"highest_bid_payment_token": paymentToken,
			"highest_bid":               amountToken,   // 最高出价金额
			"highest_bid_usd":           amountUSD,     // 最高出价USD
			"highest_bid_unit_usd":      amountUnitUSD, // 最高出价USD最小单位（8位小数）

		}).Error; err != nil {
		logger.Warn("failed to update auction bid_count: %v", err)
		return nil, fmt.Errorf("failed to update auction bid_count: %w", err)
	}

	return &bid, nil
}
