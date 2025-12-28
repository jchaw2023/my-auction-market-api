package models

import (
	"time"
)

// NFT Ownership 状态常量
const (
	NFTOwnershipStatusHolding    = "holding"    // 持有中
	NFTOwnershipStatusSelling   = "selling"    // 出售中
	NFTOwnershipStatusSold      = "sold"       // 已出售
	NFTOwnershipStatusTransfered = "transfered" // 已转移
)

// NFTOwnership NFT用户关系表
type NFTOwnership struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement;type:int(11)"`
	NFTID        string     `json:"nftId" gorm:"type:varchar(50);index;comment:NFTID"`
	UserID       int64      `json:"userId" gorm:"type:bigint(20);comment:用户ID"`
	OwnerAddress string     `json:"ownerAddress" gorm:"type:varchar(50);comment:用户钱包地址"`
	Status       string     `json:"status" gorm:"type:varchar(50);comment:状态（holding,selling,sold,transfered）"`
	Approved     bool       `json:"approved" gorm:"type:tinyint(1);comment:是否授权给平台合约"`
	Timestamp    int64      `json:"timestamp" gorm:"type:bigint(20);comment:交易时间"`
	BlockNumber  uint64     `json:"blockNumber" gorm:"type:bigint(20);comment:区块数"`
	LastSyncedAt *time.Time `json:"lastSyncedAt" gorm:"type:datetime;comment:上次同步时间"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`

	// 一对一关系：关联到 NFT 模型（通过 NFTID 字段）
	NFT NFT `json:"nft,omitempty" gorm:"foreignKey:NFTID;references:NFTID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}
