package handlers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	appJWT "my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/response"
	"my-auction-market-api/internal/services"
)

type UserHandler struct {
	service         *services.UserService
	listenerService *services.ListenerService
}

func NewUserHandler(userService *services.UserService, listenerService *services.ListenerService) *UserHandler {
	return &UserHandler{
		service:         userService,
		listenerService: listenerService,
	}
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get current user profile
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Security     BearerAuth
// @Router       /users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	profile, err := h.service.GetProfile(user.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, profile)
}

// RequestNonce godoc
// @Summary      Request a nonce for wallet login
// @Description  Generates a unique nonce for a given wallet address to be signed by the user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        walletAddress  body      models.WalletLoginRequestNoncePayload  true  "Wallet address"
// @Success      200            {object}  response.Response{data=models.WalletLoginRequestNonceResult}
// @Failure      400            {object}  response.Response
// @Failure      500            {object}  response.Response
// @Router       /auth/wallet/request-nonce [post]
func (h *UserHandler) RequestNonce(c *gin.Context) {
	var payload models.WalletLoginRequestNoncePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Error(err)
		return
	}

	result, err := h.service.RequestNonce(payload.WalletAddress)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, result)
}

// VerifyWalletLogin godoc
// @Summary      Verify wallet signature and login
// @Description  Verifies the signed message and logs in the user, returning a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        walletLogin  body      models.WalletLoginVerifyPayload  true  "Wallet login payload"
// @Success      200          {object}  response.Response{data=models.LoginResult}
// @Failure      400          {object}  response.Response
// @Failure      401          {object}  response.Response
// @Failure      500          {object}  response.Response
// @Router       /auth/wallet/verify [post]
func (h *UserHandler) VerifyWalletLogin(c *gin.Context) {
	var payload models.WalletLoginVerifyPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Error(err)
		return
	}

	user, err := h.service.VerifyWalletLogin(payload)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	// 验证签名成功后，立即将钱包地址添加到监听服务
	if h.listenerService != nil && h.listenerService.IsRunning() {
		walletAddr := common.HexToAddress(user.WalletAddress)
		if walletAddr != (common.Address{}) {
			if err := h.listenerService.AddWalletAddress(walletAddr); err != nil {
				// 添加监听失败不影响登录，只记录日志
				logger.Warn("failed to add wallet address to listener after login: %v", err)
			} else {
				logger.Info("wallet address added to listener after login: %s", user.WalletAddress)
			}
		}
	}

	// Generate JWT token
	token, err := appJWT.GenerateToken(
		user.ID,
		user.Username,
		user.Email,
	)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, models.LoginResult{
		Token: token,
		User: models.UserProfile{
			ID:            user.ID,
			Username:      user.Username,
			Email:         user.Email,
			WalletAddress: user.WalletAddress,
		},
	})
}

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Update user profile information (username and/or email)
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        profile  body      models.UpdateProfilePayload  true  "Update profile payload"
// @Success      200      {object}  response.Response{data=models.UserProfile}
// @Failure      400      {object}  response.Response
// @Failure      401      {object}  response.Response
// @Failure      500      {object}  response.Response
// @Security     BearerAuth
// @Router       /users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	user, err := appJWT.ExtractUserFromContext(c)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	var payload models.UpdateProfilePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// 验证至少提供一个字段
	if payload.Username == "" && payload.Email == "" {
		response.BadRequest(c, "at least one field (username or email) must be provided")
		return
	}

	profile, err := h.service.UpdateProfile(user.ID, payload)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, profile)
}

// GetPlatformStats godoc
// @Summary      Get platform statistics
// @Description  Get platform statistics including total users, total auctions, and total bids
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response{data=models.PlatformStatsResponse}
// @Failure      500  {object}  response.Response
// @Router       /users/stats [get]
func (h *UserHandler) GetPlatformStats(c *gin.Context) {
	stats, err := h.service.GetPlatformStats()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, stats)
}
