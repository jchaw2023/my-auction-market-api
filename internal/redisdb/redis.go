package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client Redis 客户端封装
type Client struct {
	*redis.Client
}

// Config Redis 配置
type Config struct {
	Addr         string        // 地址，例如: "localhost:6379"
	Password     string        // 密码
	DB           int           // 数据库编号
	PoolSize     int           // 连接池大小
	MinIdleConns int           // 最小空闲连接数
	DialTimeout  time.Duration // 连接超时
	ReadTimeout  time.Duration // 读取超时
	WriteTimeout time.Duration // 写入超时
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}

// NewClient 创建新的 Redis 客户端
func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &Client{Client: rdb}, nil
}

// Close 关闭 Redis 连接
func (c *Client) Close() error {
	return c.Client.Close()
}

// Ping 测试连接
func (c *Client) Ping(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}

