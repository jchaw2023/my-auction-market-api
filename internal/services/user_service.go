package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"my-auction-market-api/internal/database"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(payload models.RegisterPayload) (*models.RegisterResult, error) {
	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already exists")
	}

	// Check if username already exists
	if err := database.DB.Where("username = ?", payload.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := models.User{
		Username:      payload.Username,
		Email:         payload.Email,
		Password:      string(hashedPassword),
		WalletAddress: payload.WalletAddress,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.RegisterResult{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		WalletAddress: user.WalletAddress,
	}, nil
}

func (s *UserService) Login(payload models.LoginPayload) (*models.User, error) {
	// Find user by email or username
	var user models.User
	if err := database.DB.Where("email = ? OR username = ?", payload.Account, payload.Account).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (s *UserService) GetProfile(userID uint64) (*models.UserProfile, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &models.UserProfile{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		WalletAddress: user.WalletAddress,
	}, nil
}

// UpdateProfile 更新用户信息（用户名和邮箱）
func (s *UserService) UpdateProfile(userID uint64, payload models.UpdateProfilePayload) (*models.UserProfile, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 构建更新字段
	updates := make(map[string]interface{})

	// 如果提供了用户名，检查是否与其他用户重复
	if payload.Username != "" && payload.Username != user.Username {
		var existingUser models.User
		if err := database.DB.Where("username = ? AND id != ?", payload.Username, userID).First(&existingUser).Error; err == nil {
			return nil, errors.New("username already exists")
		}
		updates["username"] = payload.Username
	}

	// 如果提供了邮箱，检查是否与其他用户重复
	if payload.Email != "" && payload.Email != user.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ? AND id != ?", payload.Email, userID).First(&existingUser).Error; err == nil {
			return nil, errors.New("email already exists")
		}
		updates["email"] = payload.Email
	}

	// 如果没有要更新的字段，直接返回当前用户信息
	if len(updates) == 0 {
		return &models.UserProfile{
			ID:            user.ID,
			Username:      user.Username,
			Email:         user.Email,
			WalletAddress: user.WalletAddress,
		}, nil
	}

	// 更新用户信息
	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// 重新查询更新后的用户信息
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated user: %w", err)
	}

	return &models.UserProfile{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		WalletAddress: user.WalletAddress,
	}, nil
}

// RequestNonce 为钱包地址生成 nonce 和消息
func (s *UserService) RequestNonce(walletAddress string) (*models.WalletLoginRequestNonceResult, error) {
	// 规范化地址（小写）
	walletAddress = strings.ToLower(strings.TrimPrefix(walletAddress, "0x"))
	if len(walletAddress) != 40 {
		return nil, errors.New("invalid wallet address")
	}
	walletAddress = "0x" + walletAddress

	// 生成随机 nonce
	nonceBytes := make([]byte, 32)
	if _, err := rand.Read(nonceBytes); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}
	nonce := hex.EncodeToString(nonceBytes)

	// 构建消息（EIP-191 格式）
	message := fmt.Sprintf("Sign this message to authenticate with My Auction Market: %s", nonce)
	logger.Info("message: %s", message)

	// 查找或创建用户
	var user models.User
	err := database.DB.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		// 用户不存在，创建新用户
		// 生成唯一的用户名和邮箱（使用钱包地址的后8位）
		addrSuffix := walletAddress[2:10] // 去掉 0x 前缀，取前8位
		username := fmt.Sprintf("user_%s", addrSuffix)
		email := fmt.Sprintf("%s@wallet.local", addrSuffix)

		// 确保用户名和邮箱唯一（如果冲突，添加时间戳）
		var count int64
		database.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
		if count > 0 {
			username = fmt.Sprintf("user_%s_%d", addrSuffix, time.Now().Unix())
		}
		database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
		if count > 0 {
			email = fmt.Sprintf("%s_%d@wallet.local", addrSuffix, time.Now().Unix())
		}

		user = models.User{
			Username:      username,
			Email:         email,
			Password:      "", // 钱包登录不需要密码，但数据库要求非空，使用空字符串
			WalletAddress: walletAddress,
			Nonce:         nonce,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		// 更新 nonce - 使用 Update 方法明确更新字段
		if err := database.DB.Model(&user).Update("nonce", nonce).Error; err != nil {
			return nil, fmt.Errorf("failed to update nonce: %w", err)
		}
		// 同时更新内存中的 user 对象，以便后续使用
		user.Nonce = nonce
	}

	return &models.WalletLoginRequestNonceResult{
		Nonce:   nonce,
		Message: message,
	}, nil
}

// VerifyWalletLogin 验证钱包签名并登录
func (s *UserService) VerifyWalletLogin(payload models.WalletLoginVerifyPayload) (*models.User, error) {
	// 规范化地址
	walletAddress := strings.ToLower(strings.TrimPrefix(payload.WalletAddress, "0x"))
	if len(walletAddress) != 40 {
		return nil, errors.New("invalid wallet address")
	}
	walletAddress = "0x" + walletAddress

	// 查找用户
	var user models.User
	if err := database.DB.Where("wallet_address = ?", walletAddress).First(&user).Error; err != nil {
		return nil, errors.New("user not found or invalid wallet address")
	}

	// 验证 nonce（从消息中提取）
	expectedNonce := strings.TrimPrefix(payload.Message, "Sign this message to authenticate with My Auction Market: ")
	if user.Nonce == "" || user.Nonce != expectedNonce {
		return nil, errors.New("invalid or expired nonce")
	}

	// 验证签名
	isValid, err := utils.VerifySignature(payload.Message, payload.Signature, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}
	if !isValid {
		return nil, errors.New("invalid signature")
	}

	// 清除 nonce（防止重放攻击）- 使用 Update 方法明确更新字段
	if err := database.DB.Model(&user).Update("nonce", "").Error; err != nil {
		// 即使清除 nonce 失败，也允许登录（记录日志即可）
		// 但为了安全，最好还是返回错误
		return nil, fmt.Errorf("failed to clear nonce: %w", err)
	}

	// 可选：设置 nonce 过期时间（例如 5 分钟）
	// 这里简化处理，nonce 使用后立即清除

	return &user, nil
}

// GetOrCreateUserByWalletAddress 根据钱包地址获取或创建用户
func (s *UserService) GetOrCreateUserByWalletAddress(walletAddress string) (*models.User, error) {
	walletAddress = strings.ToLower(strings.TrimPrefix(walletAddress, "0x"))
	if len(walletAddress) != 40 {
		return nil, errors.New("invalid wallet address")
	}
	walletAddress = "0x" + walletAddress

	var user models.User
	err := database.DB.Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		// 用户不存在，创建新用户
		user = models.User{
			Username:      fmt.Sprintf("user_%s", walletAddress[2:10]),
			Email:         fmt.Sprintf("%s@wallet.local", walletAddress[2:10]),
			Password:      "",
			WalletAddress: walletAddress,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	return &user, nil
}

// GetPlatformStats 获取平台统计数据
func (s *UserService) GetPlatformStats() (*models.PlatformStatsResponse, error) {
	var totalUsers int64
	var totalAuctions int64
	var totalBids int64

	// 统计总用户数
	if err := database.DB.Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	// 统计拍卖总数
	if err := database.DB.Model(&models.Auction{}).Count(&totalAuctions).Error; err != nil {
		return nil, fmt.Errorf("failed to count auctions: %w", err)
	}

	// 统计出价总数
	if err := database.DB.Model(&models.Bid{}).Count(&totalBids).Error; err != nil {
		return nil, fmt.Errorf("failed to count bids: %w", err)
	}

	return &models.PlatformStatsResponse{
		TotalUsers:   uint64(totalUsers),
		TotalAuctions: uint64(totalAuctions),
		TotalBids:    uint64(totalBids),
	}, nil
}

func (s *UserService) GetAllWalletAddresses() ([]string, error) {
	var users []models.User
	// 只查询有钱包地址的用户，并且钱包地址不为空
	if err := database.DB.Where("wallet_address != ? AND wallet_address != ''", "").
		Select("wallet_address").
		Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get wallet addresses: %w", err)
	}

	walletAddresses := make([]string, 0, len(users))
	for _, user := range users {
		if user.WalletAddress != "" {
			walletAddresses = append(walletAddresses, user.WalletAddress)
		}
	}

	return walletAddresses, nil
}
