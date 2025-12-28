package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"my-auction-market-api/internal/redisdb"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName      string        `yaml:"app_name"`
	Environment  string        `yaml:"environment"`
	HTTPPort     int           `yaml:"http_port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	LogLevel     string        `yaml:"log_level"`

	Database  DatabaseConfig  `yaml:"database"`
	JWT       JWTConfig       `yaml:"jwt"`
	Ethereum  EthereumConfig  `yaml:"ethereum"`
	Etherscan EtherscanConfig `yaml:"etherscan"`
	Redis     RedisConfig     `yaml:"redis"`
}

type DatabaseConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Name            string        `yaml:"name"`
	Charset         string        `yaml:"charset"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
}

type JWTConfig struct {
	Secret     string        `yaml:"secret"`
	Expiration time.Duration `yaml:"expiration"`
}

type EthereumConfig struct {
	RPCURL                 string        `yaml:"rpc_url"`
	WssURL                 string        `yaml:"wss_url"`
	AuctionContractAddress string        `yaml:"auction_contract_address"`
	PlatformPrivateKey     string        `yaml:"platform_private_key"` // 平台私钥，用于签名合约交易
	ChainID                int64         `yaml:"chain_id"`
	WebSocketTimeout       time.Duration `yaml:"websocket_timeout"` // WebSocket 连接超时时间（默认60秒）
}

type EtherscanConfig struct {
	APIKey  string `yaml:"api_key"`
	ChainID int64  `yaml:"chain_id"`
}

type RedisConfig struct {
	Addr         string        `yaml:"addr"`
	Password     string        `yaml:"password"`
	DB           int           `yaml:"db"`
	PoolSize     int           `yaml:"pool_size"`
	MinIdleConns int           `yaml:"min_idle_conns"`
	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

func MustLoad() Config {
	configPath := getEnv("CONFIG_PATH", "config.yaml")

	cfg, err := LoadFromFile(configPath)
	if err != nil {
		cfg = getDefaultConfig()
	}

	return cfg
}

func LoadFromFile(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config file: %w", err)
	}

	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 10 * time.Second
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 15 * time.Second
	}

	if cfg.Database.MaxOpenConns == 0 {
		cfg.Database.MaxOpenConns = 25
	}
	if cfg.Database.MaxIdleConns == 0 {
		cfg.Database.MaxIdleConns = 10
	}
	if cfg.Database.ConnMaxLifetime == 0 {
		cfg.Database.ConnMaxLifetime = 5 * time.Minute
	}
	if cfg.Database.ConnMaxIdleTime == 0 {
		cfg.Database.ConnMaxIdleTime = 10 * time.Minute
	}

	if cfg.JWT.Secret == "" {
		cfg.JWT.Secret = "your-secret-key-change-in-production"
	}
	if cfg.JWT.Expiration == 0 {
		cfg.JWT.Expiration = 24 * time.Hour
	}

	// 设置默认 WebSocket 超时时间（60秒，常见值）
	if cfg.Ethereum.WebSocketTimeout == 0 {
		cfg.Ethereum.WebSocketTimeout = 60 * time.Second
	}

	// 设置 Redis 默认值
	if cfg.Redis.Addr == "" {
		cfg.Redis.Addr = "localhost:6379"
	}
	if cfg.Redis.PoolSize == 0 {
		cfg.Redis.PoolSize = 10
	}
	if cfg.Redis.MinIdleConns == 0 {
		cfg.Redis.MinIdleConns = 5
	}
	if cfg.Redis.DialTimeout == 0 {
		cfg.Redis.DialTimeout = 5 * time.Second
	}
	if cfg.Redis.ReadTimeout == 0 {
		cfg.Redis.ReadTimeout = 3 * time.Second
	}
	if cfg.Redis.WriteTimeout == 0 {
		cfg.Redis.WriteTimeout = 3 * time.Second
	}

	return cfg, nil
}

func getDefaultConfig() Config {
	return Config{
		AppName:      "my-auction-market-api",
		Environment:  "development",
		HTTPPort:     8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		LogLevel:     "info",
		Database: DatabaseConfig{
			Host:            "localhost",
			Port:            3306,
			User:            "root",
			Password:        "",
			Name:            "auction_market_db",
			Charset:         "utf8mb4",
			MaxOpenConns:    25,
			MaxIdleConns:    10,
			ConnMaxLifetime: 5 * time.Minute,
			ConnMaxIdleTime: 10 * time.Minute,
		},
		JWT: JWTConfig{
			Secret:     "your-secret-key-change-in-production",
			Expiration: 24 * time.Hour,
		},
		Ethereum: EthereumConfig{
			RPCURL:                 "https://sepolia.infura.io/v3/your-api-key",
			AuctionContractAddress: "",
			PlatformPrivateKey:     "", // 平台私钥，用于签名合约交易
			ChainID:                11155111,
			WebSocketTimeout:       60 * time.Second, // 默认60秒超时
		},
		Etherscan: EtherscanConfig{
			APIKey:  "",
			ChainID: 11155111,
		},
		Redis: RedisConfig{
			Addr:         "localhost:6379",
			Password:     "",
			DB:           0,
			PoolSize:     10,
			MinIdleConns: 5,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
		},
	}
}

func (c Config) DSN() string {
	return c.Database.User + ":" + c.Database.Password + "@tcp(" + c.Database.Host + ":" + strconv.Itoa(c.Database.Port) + ")/" + c.Database.Name + "?charset=" + c.Database.Charset + "&parseTime=True&loc=Local&allowNativePasswords=true"
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}

// ToRedisConfig 将 RedisConfig 转换为 redisdb.Config
func (r RedisConfig) ToRedisConfig() *redisdb.Config {
	return &redisdb.Config{
		Addr:         r.Addr,
		Password:     r.Password,
		DB:           r.DB,
		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
		DialTimeout:  r.DialTimeout,
		ReadTimeout:  r.ReadTimeout,
		WriteTimeout: r.WriteTimeout,
	}
}
