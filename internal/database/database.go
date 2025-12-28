package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"my-auction-market-api/internal/config"
	appLogger "my-auction-market-api/internal/logger"
)

var DB *gorm.DB

func Init(cfg config.Config) error {
	var err error

	dsn := cfg.DSN()

	var gormLogLevel gormLogger.LogLevel
	if cfg.Environment == "development" {
		gormLogLevel = gormLogger.Info
	} else {
		gormLogLevel = gormLogger.Warn
	}

	customLogger := NewGormLogger(gormLogLevel)

	// 配置 MySQL 驱动，解决 auth_gssapi_client 问题
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: customLogger,
	})

	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)

	appLogger.Info("database connected: %s@%s:%d/%s", cfg.Database.User, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	appLogger.Info("database pool configured: max_open=%d, max_idle=%d, max_lifetime=%s, max_idle_time=%s",
		cfg.Database.MaxOpenConns,
		cfg.Database.MaxIdleConns,
		cfg.Database.ConnMaxLifetime,
		cfg.Database.ConnMaxIdleTime,
	)

	return nil
}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
