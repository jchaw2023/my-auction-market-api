package database

import (
	"context"
	"errors"
	"time"

	gormLogger "gorm.io/gorm/logger"

	appLogger "my-auction-market-api/internal/logger"
)

type GormLogger struct {
	LogLevel gormLogger.LogLevel
}

func NewGormLogger(level gormLogger.LogLevel) *GormLogger {
	return &GormLogger{
		LogLevel: level,
	}
}

func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return &GormLogger{
		LogLevel: level,
	}
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		appLogger.GormInfo("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		appLogger.GormWarn("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		appLogger.GormError("[GORM] "+msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.LogLevel >= gormLogger.Error && (!errors.Is(err, gormLogger.ErrRecordNotFound)):
		appLogger.GormError("[GORM] %s [%s] [rows:%d] %s", err, elapsed, rows, sql)
	case elapsed > 200*time.Millisecond && l.LogLevel >= gormLogger.Warn:
		appLogger.GormWarn("[GORM] [slow query] [%s] [rows:%d] %s", elapsed, rows, sql)
	case l.LogLevel == gormLogger.Info:
		appLogger.GormInfo("[GORM] [%s] [rows:%d] %s", elapsed, rows, sql)
	}
}

