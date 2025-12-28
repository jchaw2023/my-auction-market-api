package models

import (
	"time"

	"my-auction-market-api/internal/utils"

	"github.com/shopspring/decimal"
)

type Auction struct {
	ID                     uint64           `json:"-" gorm:"primaryKey;autoIncrement;type:bigint(20) unsigned"`
	AuctionID              string           `json:"auctionId" gorm:"type:varchar(50);uniqueIndex:auction_id;comment:拍卖ID(使用snowflake生成ID，避免自增ID)"`
	UserID                 uint64           `json:"userId" gorm:"type:bigint(20) unsigned;index:idx_auctions_user_id;comment:创建者ID"`
	OwnerAddress           string           `json:"ownerAddress" gorm:"type:varchar(42);index:idx_auctions_owner;comment:当前拥有者地址"`
	ContractAuctionID      uint64           `json:"contractAuctionId" gorm:"type:bigint(20) unsigned;not null;index:idx_auctions_contract_id;comment:合约里面拍卖列表索引"`
	NFTID                  string           `json:"nftId" gorm:"type:varchar(64);index:idx_auctions_nft_id;comment:NFT唯一标识"`
	NFTAddress             string           `json:"nftAddress" gorm:"type:varchar(42);not null;index:idx_auctions_nft_address;comment:NFT合约地址"`
	TokenID                uint64           `json:"tokenId" gorm:"type:bigint(20) unsigned;not null;index:idx_auctions_token_id;comment:NFT的Token ID"`
	ContractName           string           `json:"contractName" gorm:"type:varchar(255);comment:合约名称"`
	ContractSymbol         string           `json:"contractSymbol" gorm:"type:varchar(64);comment:合约符号"`
	TokenURI               string           `json:"tokenURI" gorm:"type:text;comment:Token URI"`
	NftName                string           `json:"nftName" gorm:"type:varchar(255);comment:NFT名称"`
	Image                  string           `json:"image" gorm:"type:text;comment:NFT图片URL"`
	Description            string           `json:"description" gorm:"type:text;comment:NFT描述"`
	Metadata               string           `json:"metadata" gorm:"type:longtext;comment:完整元数据JSON"`
	Status                 string           `json:"status" gorm:"type:varchar(20);not null;default:'pending';index:idx_auctions_status;comment:状态(pending,active,ended,cancelled)"`
	OnlineLock             string           `json:"onlineLock" gorm:"column:online_lock;type:varchar(76);uniqueIndex:nft_online_id;comment:NFT在线标志 nft_id:1,也作为一个锁字段，解锁就改成其他值"`
	Online                 uint64           `json:"online" gorm:"type:bigint(20);index:online;comment:1表示在线 其他值表示下线"`
	StartTime              *time.Time       `json:"startTime" gorm:"type:datetime;not null;comment:开始时间"`
	StartTimestamp         uint64           `json:"startTimestamp" gorm:"type:bigint(20);not null;default:0;index:start_timestamp;comment:开始时间时间戳"`
	EndTime                *time.Time       `json:"endTime" gorm:"type:datetime;not null;comment:结束时间"`
	EndTimestamp           uint64           `json:"endTimestamp" gorm:"type:bigint(20);not null;default:0;index:end_timestamp;comment:结束时间时间戳"`
	StartPrice             *decimal.Decimal `json:"startPrice" gorm:"type:decimal(65,30);comment:起拍价(单位由PaymentToken指定:0x0=ETH,其他=ERC20代币)"`
	PaymentToken           string           `json:"paymentToken" gorm:"type:varchar(42);comment:起拍价链上交易代币地址(0x0表示ETH,其他地址表示ERC20代币)"`
	StartPriceUSD          *decimal.Decimal `json:"startPriceUSD" gorm:"type:decimal(65,30);comment:起拍价USD"`
	StartPriceUnitUSD      uint64           `json:"startPriceUnitUSD" gorm:"column:start_price_unit_usd;type:bigint(20);comment:起拍价USD预言机价格（小数点起拍价USD*10**8）"`
	HighestBidder          string           `json:"highestBidder" gorm:"type:varchar(42);comment:最高出价者地址"`
	HighestBidPaymentToken string           `json:"highestBidPaymentToken" gorm:"type:varchar(42);comment:最高出价使用链上交易代币地址(0x0=ETH,其他=ERC20代币,可能与拍卖PaymentToken不同)"`
	HighestBid             *decimal.Decimal `json:"highestBid" gorm:"type:decimal(65,30) unsigned;comment:最高出价金额(单位由HighestBidPaymentToken指定,可能与拍卖PaymentToken不同)"`
	HighestBidUSD          *decimal.Decimal `json:"highestBidUSD" gorm:"type:decimal(65,30);comment:最高出价USD(用于比较不同代币的出价)"`
	HighestBidUnitUSD      uint64           `json:"highestBidUnitUSD" gorm:"column:highest_bid_unit_usd;type:bigint(20);comment:最高出价USD预言机价格（小数点最高价USD*10**8）"`
	BidCount               uint64           `json:"bidCount" gorm:"type:bigint(20) unsigned;default:0;comment:出价次数"`
	CreatedAt              *time.Time       `json:"createdAt" gorm:"type:datetime;not null;default:current_timestamp;comment:创建时间"`
	UpdatedAt              *time.Time       `json:"updatedAt" gorm:"type:datetime;not null;default:current_timestamp on update current_timestamp;comment:更新时间"`

	User User  `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	Bids []Bid `json:"bids,omitempty" gorm:"foreignKey:AuctionID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

// IsETH 判断支付代币是否为ETH
func (a *Auction) IsETH() bool {
	return a.PaymentToken == utils.ETHAddress ||
		a.PaymentToken == "0x0" ||
		a.PaymentToken == ""
}

// GetPaymentTokenSymbol 获取支付代币符号(用于显示)
func (a *Auction) GetPaymentTokenSymbol() string {
	if a.IsETH() {
		return "ETH"
	}
	// 如果是ERC20代币，可以后续通过合约查询获取symbol
	// 这里先返回通用标识
	return "ERC20"
}

// IsBidTokenSameAsAuction 判断出价代币是否与拍卖指定代币相同
func (a *Auction) IsBidTokenSameAsAuction(bidPaymentToken string) bool {
	if a.PaymentToken == "" {
		return bidPaymentToken == "" || bidPaymentToken == "0x0" || bidPaymentToken == utils.ETHAddress
	}
	return a.PaymentToken == bidPaymentToken
}

// CompareBids 比较两个出价的USD价值(用于判断哪个出价更高)
// 返回: 1表示bid1更高, -1表示bid2更高, 0表示相等
func CompareBids(bid1AmountUSD, bid2AmountUSD decimal.Decimal) int {
	return bid1AmountUSD.Cmp(bid2AmountUSD)
}

type AuctionPayload struct {
	NFTID        string          `json:"nftId" binding:"required"` // NFT唯一标识（从前端传入）
	NFTAddress   string          `json:"nftAddress" binding:"required"`
	TokenID      uint64          `json:"tokenId" binding:"required"`
	PaymentToken string          `json:"paymentToken" binding:"required"` // 支付代币地址(0x0表示ETH,其他表示ERC20代币)
	StartPrice   decimal.Decimal `json:"startPrice" binding:"required"`   // 起拍价(单位由PaymentToken指定)
	StartTime    *time.Time      `json:"startTime" binding:"required"`    // ISO 8601 格式
	EndTime      *time.Time      `json:"endTime" binding:"required"`      // ISO 8601 格式
}

type Bid struct {
	ID                uint64           `json:"id" gorm:"primaryKey;autoIncrement;type:bigint(20) unsigned;comment:出价ID"`
	AuctionID         string           `json:"auctionId" gorm:"type:varchar(50);comment:拍卖ID"`
	ContractAuctionID uint64           `json:"contractAuctionId" gorm:"type:bigint(20) unsigned;not null;comment:拍卖合约里面的拍卖ID"`
	UserID            uint64           `json:"userId" gorm:"type:bigint(20) unsigned;not null;index:idx_bids_user_id;comment:出价者ID"`
	WalletAddress     string           `json:"walletAddress" gorm:"type:varchar(50);comment:出价者钱包地址"`
	Winner            bool             `json:"winner" gorm:"type:tinyint(1);default:0;comment:竞拍获胜者"`
	Amount            *decimal.Decimal `json:"amount" gorm:"type:decimal(20,8);not null;default:0.00000000;comment:出价金额(ETH,USDC)"`
	AmountUnit        uint64           `json:"amountUnit" gorm:"column:amount_unit;type:bigint(20);not null;default:0;comment:出价金额(wei、usdc最小单位等)"`
	AmountUSD         *decimal.Decimal `json:"amountUSD" gorm:"type:decimal(20,8);comment:出价金额USD"`
	AmountUnitUSD     uint64           `json:"amountUnitUSD" gorm:"column:amount_unit_usd;type:bigint(20);comment:出价金额8位"`
	PaymentToken      string           `json:"paymentToken" gorm:"type:varchar(42);comment:支付代币地址"`
	TransactionHash   string           `json:"transactionHash" gorm:"type:varchar(66);index:idx_bids_transaction_hash;comment:交易哈希"`
	BlockNumber       uint64           `json:"blockNumber" gorm:"type:bigint(20) unsigned;comment:区块号"`
	Timestamp         uint64           `json:"timestamp" gorm:"type:bigint(20) unsigned;comment:链上时间"`
	BidCount          uint64           `json:"bidCount" gorm:"type:bigint(20);comment:出价总数"`
	IsHighest         bool             `json:"isHighest" gorm:"type:tinyint(1);default:0;index:idx_bids_is_highest;comment:是否为最高出价"`
	MinBidder         string           `json:"minBidder" gorm:"type:varchar(42);comment:上一个最高出价值地址（当前出价起码要超过的最小金额地址）"`
	MinBidUnitUSD     uint64           `json:"minBidUnitUSD" gorm:"column:min_bid_unit_usd;type:bigint(20);comment:上一个最高出价值（当前出价起码要超过的最小金额数）"`
	CreatedAt         *time.Time       `json:"createdAt" gorm:"type:datetime;not null;default:current_timestamp;index:idx_bids_created_at;comment:创建时间"`

	Auction Auction `json:"auction,omitempty" gorm:"foreignKey:AuctionID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

// IsETH 判断支付代币是否为ETH
func (b *Bid) IsETH() bool {
	return b.PaymentToken == utils.ETHAddress ||
		b.PaymentToken == "0x0" ||
		b.PaymentToken == ""
}

type BidPayload struct {
	AuctionID    uint64 `json:"auctionId" binding:"required"`
	Amount       string `json:"amount" binding:"required"`       // 出价金额
	PaymentToken string `json:"paymentToken" binding:"required"` // 出价使用的代币地址(可以与拍卖PaymentToken不同)
}

// UpdateAuctionPayload 更新拍卖信息的请求体（只包含可更新的字段）
type UpdateAuctionPayload struct {
	PaymentToken string          `json:"paymentToken" binding:"required"` // 支付代币地址(0x0表示ETH,其他表示ERC20代币)
	StartPrice   decimal.Decimal `json:"startPrice" binding:"required"`   // 起拍价(单位由PaymentToken指定)
	StartTime    *time.Time      `json:"startTime" binding:"required"`    // ISO 8601 格式
	EndTime      *time.Time      `json:"endTime" binding:"required"`      // ISO 8601 格式
}

// ConvertToUSDPayload 转换金额为美元的请求体
type ConvertToUSDPayload struct {
	Token  string  `json:"token" binding:"required"`  // 代币地址(0x0表示ETH,其他表示ERC20代币)
	Amount float64 `json:"amount" binding:"required"` // 金额(例如 1.0 表示 1 ETH)
}

// CheckNFTApprovalPayload 检查NFT授权状态的请求体
type CheckNFTApprovalPayload struct {
	NFTAddress string `json:"nftAddress" binding:"required"` // NFT合约地址
	TokenID    uint64 `json:"tokenId" binding:"required"`    // Token ID
}

// ConvertToUSDResponse 转换金额为美元的响应
type ConvertToUSDResponse struct {
	Token         string  `json:"token"`         // 代币地址
	Amount        float64 `json:"amount"`        // 原始金额
	AmountUSD     float64 `json:"amountUSD"`     // 美元金额
	AmountUnitUSD uint64  `json:"amountUnitUSD"` // 美元金额(小数点后8位)
	AmountUSDStr  string  `json:"amountUSDStr"`  // 美元金额(字符串格式，便于前端显示)
}

// TokenPriceResponse 代币价格响应
type TokenPriceResponse struct {
	Token        string  `json:"token"`        // 代币地址
	Price        float64 `json:"price"`        // 代币价格（1个代币 = X USD）
	PriceUnitUSD uint64  `json:"priceUnitUSD"` // 代币价格(小数点后8位，Chainlink格式)
	PriceUSDStr  string  `json:"priceUSDStr"`  // 代币价格(字符串格式，便于前端显示)
	Decimals     uint8   `json:"decimals"`     // 价格精度（通常是8，Chainlink标准）
}

// ConvertUSDToTokenResponse 将USD金额转换为代币金额的响应
type ConvertUSDToTokenResponse struct {
	TokenAddress  string  `json:"tokenAddress"`  // 代币地址
	USDAmount     float64 `json:"usdAmount"`     // 美元金额（小数格式）
	TokenAmount   float64 `json:"tokenAmount"`   // 代币金额（小数格式）
	TokenDecimals uint8   `json:"tokenDecimals"` // 代币精度
	USDUnitAmount uint64  `json:"usdUnitAmount"` // 美元金额（8位小数格式，Chainlink格式）
}

// AuctionSimpleStatsResponse 拍卖简单统计响应
type AuctionSimpleStatsResponse struct {
	TotalAuctionsCreated uint64  `json:"totalAuctionsCreated"` // 拍卖总数
	TotalBidsPlaced      uint64  `json:"totalBidsPlaced"`      // 出价总数
	PlatformFee          uint64  `json:"platformFee"`          // 平台费用
	TotalValueLocked     float64 `json:"totalValueLocked"`     // 总锁定价值（USD，8位小数转换为float64）
	TotalValueLockedStr  string  `json:"totalValueLockedStr"`   // 总锁定价值（字符串格式，便于前端显示）
}

// AuctionNFTItem 拍卖中的 NFT 信息（用于 NFT 列表接口）
type AuctionNFTItem struct {
	NFTAddress     string `json:"nftAddress"`     // NFT合约地址
	TokenID        uint64 `json:"tokenId"`        // Token ID
	NFTID          string `json:"nftId"`          // NFT唯一标识
	Name           string `json:"name"`           // NFT名称
	Image          string `json:"image"`          // NFT图片URL
	ContractName   string `json:"contractName"`   // 合约名称
	ContractSymbol string `json:"contractSymbol"` // 合约符号
	TokenURI       string `json:"tokenURI"`       // Token URI
	Description    string `json:"description"`    // NFT描述
	AuctionCount   int64  `json:"auctionCount"`   // 该NFT的拍卖数量
}

// AuctionDetailResponse 拍卖详情响应（只包含钱包地址，不包含完整User信息）
type AuctionDetailResponse struct {
	Auction
	SellerWalletAddress string `json:"sellerWalletAddress"` // 卖家钱包地址（只返回钱包地址，不返回完整User信息）
}

// BidDetailResponse 出价详情响应（只包含钱包地址，不包含完整User信息）
type BidDetailResponse struct {
	Bid
	BidderWalletAddress string `json:"bidderWalletAddress"` // 出价者钱包地址（只返回钱包地址，不返回完整User信息）
}

// BidResponse 出价响应（用于 WebSocket 和 API，统一格式）
type BidResponse struct {
	ID                 uint64     `json:"id"`                 // 出价ID
	AuctionID          string     `json:"auctionId"`         // 拍卖ID
	ContractAuctionID  uint64     `json:"contractAuctionId"` // 合约中的拍卖ID
	UserID             uint64     `json:"userId"`            // 出价者用户ID
	Bidder             string     `json:"bidder"`            // 出价者钱包地址
	Amount             float64    `json:"amount"`            // 出价金额(ETH,USDC) - 数值类型
	AmountUnit         uint64     `json:"amountUnit"`        // 出价金额(wei、usdc最小单位等)
	AmountUSD          float64    `json:"amountUSD"`         // 出价金额USD - 数值类型
	AmountUnitUSD      uint64     `json:"amountUnitUSD"`     // 出价金额USD最小单位（8位小数）
	PaymentToken       string     `json:"paymentToken"`       // 支付代币地址
	PaymentTokenSymbol string     `json:"paymentTokenSymbol"` // 支付代币符号
	TransactionHash    string     `json:"transactionHash"`   // 交易哈希
	BlockNumber        uint64     `json:"blockNumber"`        // 区块号
	Timestamp          uint64     `json:"timestamp"`         // 链上时间
	BidCount           uint64     `json:"bidCount"`         // 出价总数
	IsHighest          bool       `json:"isHighest"`         // 是否为最高出价
	MinBidder          string     `json:"minBidder"`        // 上一个最高出价者地址
	MinBidUnitUSD      uint64     `json:"minBidUnitUSD"`    // 上一个最高出价值（USD最小单位）
	CreatedAt          *time.Time `json:"createdAt"`         // 创建时间
}

// AuctionHistoryResponse 拍卖历史记录响应（精简字段）
type AuctionHistoryResponse struct {
	AuctionID              string           `json:"auctionId"`                        // 拍卖ID
	NFTID                  string           `json:"nftId"`                            // NFT唯一标识
	NFTAddress             string           `json:"nftAddress"`                       // NFT合约地址
	TokenID                uint64           `json:"tokenId"`                          // Token ID
	NftName                string           `json:"nftName"`                          // NFT名称
	Image                  string           `json:"image"`                            // NFT图片URL
	ContractName           string           `json:"contractName"`                     // 合约名称
	ContractSymbol         string           `json:"contractSymbol"`                   // 合约符号
	Status                 string           `json:"status"`                           // 状态
	PaymentToken           string           `json:"paymentToken"`                     // 支付代币地址（用于地板价）
	FloorPrice             *decimal.Decimal `json:"floorPrice"`                       // 地板价格（起拍价）
	FloorPriceUSD          *decimal.Decimal `json:"floorPriceUSD"`                    // 地板价格USD
	EndTime                *time.Time       `json:"endTime"`                          // 结束时间
	EndTimestamp           uint64           `json:"endTimestamp"`                     // 结束时间戳
	BidCount               uint64           `json:"bidCount"`                         // 出价次数
	HighestBid             *decimal.Decimal `json:"highestBid,omitempty"`             // 当前最高价（如果有出价）
	HighestBidUSD          *decimal.Decimal `json:"highestBidUSD,omitempty"`          // 当前最高价USD（如果有出价）
	HighestBidder          string           `json:"highestBidder,omitempty"`          // 出价人地址（如果有出价）
	HighestBidPaymentToken string           `json:"highestBidPaymentToken,omitempty"` // 最高出价使用的代币地址
}
