// @title           My Auction Market API
// @version         1.0
// @description     A RESTful API for NFT auction marketplace using Gin and GORM
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @schemes   http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
package main

import (
	_ "my-auction-market-api/docs"
	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/server"
)

func main() {
	cfg := config.MustLoad()
	logger.Init(cfg.LogLevel)
	jwt.Init(cfg.JWT)

	srv, err := server.New(cfg)
	if err != nil {
		logger.Fatalf("server initialization failed: %v", err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatalf("server exited with error: %v", err)
	}
}
