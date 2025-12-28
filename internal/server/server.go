package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/database"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/middleware"
	"my-auction-market-api/internal/router"
	"my-auction-market-api/internal/services"
)

type Server struct {
	engine         *gin.Engine
	cfg            config.Config
	serviceManager *services.ServiceManager
}

func New(cfg config.Config) (*Server, error) {
	if err := database.Init(cfg); err != nil {
		return nil, fmt.Errorf("database init failed: %w", err)
	}

	// Auto migrate database
	// if err := database.AutoMigrate(
	// 	&models.User{},
	// 	&models.Auction{},
	// 	&models.Bid{},
	// ); err != nil {
	// 	return nil, fmt.Errorf("database migration failed: %w", err)
	// }

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestLogger())
	engine.Use(middleware.ValidationErrorHandler())

	// 初始化服务管理器（统一管理所有业务服务）
	serviceManager, err := services.NewServiceManager(cfg)
	if err != nil {
		return nil, fmt.Errorf("service manager initialization failed: %w", err)
	}

	// 注册路由（传递服务管理器）
	if err := router.Register(engine.Group("/api"), cfg, serviceManager); err != nil {
		return nil, fmt.Errorf("router registration failed: %w", err)
	}

	return &Server{
		engine:         engine,
		cfg:            cfg,
		serviceManager: serviceManager,
	}, nil
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.cfg.HTTPPort)

	// 创建上下文用于优雅关闭
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动拍卖任务调度器
	s.serviceManager.StartAuctionTaskScheduler(ctx)

	// 启动区块链事件监听服务（如果已初始化）
	if err := s.serviceManager.StartListenerService(); err != nil {
		logger.Warn("failed to start listener service: %v", err)
		logger.Warn("continuing without blockchain event listener")
	}

	// 确保在服务器关闭时清理所有服务
	defer func() {
		if err := s.serviceManager.Close(); err != nil {
			logger.Error("error closing services: %v", err)
		}
	}()

	server := &http.Server{
		Addr:         addr,
		Handler:      s.engine,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	logger.Info("starting %s in %s on %s", s.cfg.AppName, s.cfg.Environment, addr)

	return server.ListenAndServe()
}
