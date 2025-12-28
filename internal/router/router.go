package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/handlers"
	"my-auction-market-api/internal/middleware"
	"my-auction-market-api/internal/services"
	"my-auction-market-api/internal/websocket"
)

func Register(rg *gin.RouterGroup, cfg config.Config, smr *services.ServiceManager) error {
	// Swagger documentation
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rg.GET("/health", handlers.HealthCheck)
	rg.GET("/config/ethereum", handlers.GetEthereumConfig(cfg))

	// 使用服务管理器中的服务创建handlers
	userHandler := handlers.NewUserHandler(smr.UserService, smr.ListenerService)
	auctionHandler := handlers.NewAuctionHandler(smr.AuctionService)
	bidHandler := handlers.NewBidHandler(smr.BidService, smr.AuctionService)
	nftHandler := handlers.NewNFTHandler(smr.NFTService)
	auctionTaskHandler := handlers.NewAuctionTaskHandler(smr.AuctionTaskScheduler)

	// Auth routes (no authentication required) - Wallet login only
	auth := rg.Group("/auth")
	{
		// Wallet login routes
		auth.POST("/wallet/request-nonce", userHandler.RequestNonce)
		auth.POST("/wallet/verify", userHandler.VerifyWalletLogin)
	}

	// Users routes
	users := rg.Group("/users")
	{
		// 公开接口：平台统计数据（不需要认证）
		users.GET("/stats", userHandler.GetPlatformStats)
		
		// 需要认证的接口
		usersAuth := users.Group("")
		usersAuth.Use(middleware.AuthMiddleware())
		{
			usersAuth.GET("/profile", userHandler.GetProfile)
			usersAuth.PUT("/profile", userHandler.UpdateProfile)
		}
	}

	// Auctions routes
	auctions := rg.Group("/auctions")
	{
		auctions.GET("", auctionHandler.List)
		auctions.GET("/public", auctionHandler.ListPublic) // 公开拍卖列表（首页专用，带排序）
		// 静态路由必须在动态路由之前
		auctions.GET("/stats", auctionHandler.GetAuctionSimpleStats) // 拍卖简单统计
		auctions.GET("/nfts", auctionHandler.ListNFTs)
		auctions.GET("/supported-tokens", auctionHandler.GetSupportedTokens)
		auctions.GET("/token-price/:token", auctionHandler.GetTokenPrice)
		auctions.POST("/convert-to-usd", auctionHandler.ConvertToUSD)
		auctions.POST("/check-nft-approval", auctionHandler.CheckNFTApproval)
		// More specific routes must come before wildcard routes
		auctions.GET("/:id/detail", auctionHandler.GetDetailByID)
		auctions.GET("/:id/bids", bidHandler.GetBidsByAuctionID)
		auctions.GET("/:id", auctionHandler.GetByID)

		auctionsAuth := auctions.Group("")
		auctionsAuth.Use(middleware.AuthMiddleware())
		{
			auctionsAuth.POST("", auctionHandler.Create)
			auctionsAuth.PUT("/:id", auctionHandler.Update)
			auctionsAuth.POST("/:id/cancel", auctionHandler.Cancel)
			// 更具体的路由必须在通用路由之前
			auctionsAuth.GET("/my/history", auctionHandler.GetUserAuctionHistory)
			auctionsAuth.GET("/my", auctionHandler.GetUserAuctions)
		}
	}

	// Bids routes
	bids := rg.Group("/bids")
	{
		bids.GET("/:id", bidHandler.GetByID)
	}

	// Auction Task routes (require authentication)
	auctionTasks := rg.Group("/auction-tasks")
	auctionTasks.Use(middleware.AuthMiddleware())
	{
		auctionTasks.GET("/:auctionId", auctionTaskHandler.GetTaskStatus)
		auctionTasks.DELETE("/:auctionId", auctionTaskHandler.CancelAuctionTask)
	}

	// NFT routes (require authentication)
	nfts := rg.Group("/nfts")
	nfts.Use(middleware.AuthMiddleware())
	{
		nfts.GET("/owned", nftHandler.GetOwnedNFTs)
		nfts.POST("/sync", nftHandler.SyncNFTs)
		nfts.GET("/sync/status", nftHandler.GetSyncStatus)
		nfts.GET("/my", nftHandler.GetMyNFTs)
		nfts.GET("/my/list", nftHandler.GetMyNFTsList)
		nfts.GET("/my/ownership/:nftId", nftHandler.GetMyNFTOwnershipByNFTID)
		nfts.GET("/:id", nftHandler.GetNFTByID)
		nfts.POST("/verify", nftHandler.VerifyOwnership)
	}

	// WebSocket 路由
	rg.GET("/ws", func(c *gin.Context) {
		websocket.ServeWS(smr.WSHub, c)
	})

	// 可选：需要认证的 WebSocket 路由
	wsAuth := rg.Group("/ws")
	wsAuth.Use(middleware.AuthMiddleware())
	{
		wsAuth.GET("/auth", func(c *gin.Context) {
			websocket.ServeWS(smr.WSHub, c)
		})
	}

	return nil
}
