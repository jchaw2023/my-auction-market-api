package models

import (
	"time"
)

type User struct {
	ID            uint64    `json:"-" gorm:"primaryKey;autoIncrement;type:bigint(20) unsigned;comment:用户ID"`
	Username      string    `json:"username" gorm:"type:varchar(64);not null;uniqueIndex:idx_users_username;comment:用户名"`
	Email         string    `json:"email" gorm:"type:varchar(128);not null;uniqueIndex:idx_users_email;comment:邮箱"`
	Password      string    `json:"-" gorm:"type:varchar(255);not null;comment:密码哈希"`
	WalletAddress string    `json:"walletAddress" gorm:"type:varchar(42);index:idx_users_wallet_address;comment:钱包地址"`
	Nonce         string    `json:"-" gorm:"type:varchar(64);comment:登录Nonce"`
	CreatedAt     time.Time `json:"createdAt" gorm:"type:datetime;not null;default:current_timestamp;comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"type:datetime;not null;default:current_timestamp on update current_timestamp;comment:更新时间"`

	Auctions []Auction `json:"auctions,omitempty" gorm:"foreignKey:UserID"`
	Bids     []Bid     `json:"bids,omitempty" gorm:"foreignKey:UserID"`
}

type RegisterPayload struct {
	Username      string `json:"username" binding:"required,min=1,max=64"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=6"`
	WalletAddress string `json:"walletAddress" binding:"omitempty"`
}

type RegisterResult struct {
	ID            uint64 `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	WalletAddress string `json:"walletAddress"`
}

type LoginPayload struct {
	Account  string `json:"account" binding:"required"` // can be email or username
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResult struct {
	Token string      `json:"token"`
	User  UserProfile `json:"user"`
}

type UserProfile struct {
	ID            uint64 `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	WalletAddress string `json:"walletAddress"`
}

// WalletLoginRequestNoncePayload 钱包登录请求 nonce
type WalletLoginRequestNoncePayload struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

// WalletLoginRequestNonceResult 返回 nonce 和消息
type WalletLoginRequestNonceResult struct {
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

// WalletLoginVerifyPayload 钱包登录验证签名
type WalletLoginVerifyPayload struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Message       string `json:"message" binding:"required"`
	Signature     string `json:"signature" binding:"required"`
}

// PlatformStatsResponse 平台统计数据响应
type PlatformStatsResponse struct {
	TotalUsers    uint64 `json:"totalUsers"`    // 总用户数
	TotalAuctions uint64 `json:"totalAuctions"` // 拍卖总数
	TotalBids     uint64 `json:"totalBids"`     // 出价总数
}

// UpdateProfilePayload 更新用户信息请求
type UpdateProfilePayload struct {
	Username string `json:"username" binding:"omitempty,min=1,max=64"` // 用户名（可选）
	Email    string `json:"email" binding:"omitempty,email"`           // 邮箱（可选）
}
