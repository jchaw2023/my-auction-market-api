package models

import (
	"time"
)

// NFT NFT信息模型
type NFT struct {
	ID              uint64    `json:"id" gorm:"primaryKey;autoIncrement;type:bigint(20) unsigned"`
	NFTID           string    `json:"nftId" gorm:"type:varchar(64);not null;uniqueIndex:idx_nft_id;comment:NFT唯一标识"`
	UserID          uint64    `json:"userId" gorm:"type:bigint(20) unsigned;not null;index:idx_nfts_user_id;comment:用户ID"`
	ContractAddress string    `json:"contractAddress" gorm:"type:varchar(42);not null;index:idx_nfts_contract;comment:NFT合约地址"`
	TokenID         uint64    `json:"tokenId" gorm:"type:bigint(20) unsigned;not null;index:idx_nfts_token_id;comment:Token ID"`
	TokenURI        string    `json:"tokenURI" gorm:"type:text;comment:Token URI"`
	ContractName    string    `json:"contractName" gorm:"type:varchar(255);comment:合约名称"`
	ContractSymbol  string    `json:"contractSymbol" gorm:"type:varchar(64);comment:合约符号"`
	NftOwnerAddress string    `json:"nftOwnerAddress" gorm:"type:varchar(42);index:idx_nfts_owner;comment:当前拥有者地址"`
	NftName         string    `json:"nftName" gorm:"type:varchar(255);comment:NFT名称"`
	Image           string    `json:"image" gorm:"type:text;comment:NFT图片URL"`
	Description     string    `json:"description" gorm:"type:text;comment:NFT描述"`
	Metadata        string    `json:"metadata" gorm:"type:json;comment:完整元数据JSON"`
	LastSyncedAt    time.Time `json:"lastSyncedAt" gorm:"type:datetime;comment:上次同步时间"`
	CreatedAt       time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

// NFTOwnershipVerifyPayload NFT所有权验证请求
type NFTOwnershipVerifyPayload struct {
	ContractAddress string `json:"contractAddress" binding:"required"`
	TokenID         uint64 `json:"tokenId" binding:"required"`
}

// NFTSyncPayload NFT同步请求
type NFTSyncPayload struct {
	ContractAddress string `json:"contractAddress" binding:"omitempty"` // 可选，如果不提供则同步所有合约
}

// NFTSyncResult NFT同步结果
type NFTSyncResult struct {
	TotalFound  int `json:"totalFound"`  // 链上找到的NFT数量
	TotalSynced int `json:"totalSynced"` // 成功同步的数量
	TotalFailed int `json:"totalFailed"` // 同步失败的数量
}

// NFTSyncStatus NFT同步状态
type NFTSyncStatus struct {
	LastSyncAt time.Time `json:"lastSyncAt"` // 上次同步时间
	TotalNFTs  int       `json:"totalNFTs"`  // 数据库中NFT总数
	IsSyncing  bool      `json:"isSyncing"`  // 是否正在同步中
}
