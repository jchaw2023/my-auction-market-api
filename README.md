# My Auction Market API

NFT 拍卖商城后端 API，基于 Go + Gin 框架构建，整合以太坊客户端用于与智能合约交互。提供完整的 RESTful API 和 WebSocket 实时推送服务。

## 📋 项目概述

这是一个去中心化 NFT 拍卖平台的后端服务，支持用户通过钱包登录、创建和管理 NFT 拍卖、链上出价等功能。系统通过监听区块链事件实现链上数据同步，使用 Redis 和任务队列处理定时任务，通过 WebSocket 实时推送拍卖状态更新。

### 核心特性

- 🔐 **钱包签名认证**：基于 MetaMask 钱包签名的 JWT 认证机制
- 🎨 **NFT 拍卖管理**：创建、查询、更新、取消 NFT 拍卖
- 💰 **链上出价**：支持链上出价交易，自动同步到数据库
- 📊 **实时数据同步**：WebSocket 实时推送拍卖状态和出价更新
- 🔗 **区块链集成**：实时监听智能合约事件并同步到数据库
- ⏰ **任务调度**：基于 Redis 和 Asynq 的可靠任务调度系统
- 💱 **价格转换**：集成 Chainlink 价格预言机，支持 USD 价值计算
- 📡 **事件监听**：监听合约事件（拍卖创建、出价、结束等）
- 📄 **API 文档**：自动生成的 Swagger API 文档
- 🔒 **安全设计**：JWT 认证、请求验证、敏感信息保护

---

## 🏗️ 技术栈

### 核心框架

- **语言**: Go 1.25
- **Web 框架**: Gin v1.11.0
- **ORM**: GORM v1.31.1
- **数据库驱动**: MySQL Driver v1.6.0

### 数据库与缓存

- **数据库**: MySQL 8.0+
- **缓存/队列**: Redis (go-redis/v9)
- **任务队列**: Asynq v0.25.1 (基于 Redis)

### 区块链相关

- **以太坊客户端**: go-ethereum v1.14.0
- **Web3 交互**: ethclient, accounts/abi
- **智能合约**: ERC721 NFT、ERC20 Token、拍卖合约

### 认证与安全

- **JWT**: golang-jwt/jwt/v5 v5.3.0
- **密码加密**: golang.org/x/crypto (bcrypt)
- **请求验证**: go-playground/validator/v10

### 其他组件

- **日志**: zerolog v1.34.0
- **配置管理**: gopkg.in/yaml.v3
- **WebSocket**: gorilla/websocket v1.4.2
- **HTTP 客户端**: go-resty/resty/v2 (用于调用外部 API)
- **数值计算**: shopspring/decimal (精确金额计算)
- **ID 生成**: sony/sonyflake (分布式唯一 ID)
- **API 文档**: swaggo/gin-swagger, swaggo/swag

### 开发工具

- **代码生成**: swag (Swagger 文档生成)
- **依赖管理**: Go Modules

---

## 📁 项目结构

```
my-auction-market-api/
├── cmd/
│   └── server/
│       └── main.go                    # 应用入口点，初始化配置和启动服务器
│
├── internal/                          # 内部代码包（不对外暴露）
│   ├── config/                        # 配置管理
│   │   └── config.go                  # 配置文件加载、解析和验证
│   │
│   ├── database/                      # 数据库相关
│   │   ├── database.go                # 数据库连接初始化、GORM 配置
│   │   └── gorm_logger.go             # GORM 日志适配器（集成 zerolog）
│   │
│   ├── handlers/                      # HTTP 请求处理器（Controller 层）
│   │   ├── auction_handler.go         # 拍卖相关接口处理
│   │   ├── bid_handler.go             # 出价相关接口处理
│   │   ├── user_handler.go            # 用户相关接口处理
│   │   ├── nft_handler.go             # NFT 相关接口处理
│   │   ├── auction_task_handler.go    # 拍卖任务调度接口处理
│   │   ├── config_handler.go          # 配置接口处理（返回公开配置）
│   │   └── health_handler.go          # 健康检查接口
│   │
│   ├── middleware/                    # HTTP 中间件
│   │   ├── auth.go                    # JWT 认证中间件
│   │   ├── request_logger.go          # 请求日志中间件
│   │   └── validation.go              # 请求验证中间件
│   │
│   ├── models/                        # 数据模型（数据库实体）
│   │   ├── user.go                    # 用户模型
│   │   ├── auction.go                 # 拍卖模型
│   │   ├── nft.go                     # NFT 模型
│   │   └── nft_ownership.go           # NFT 所有权关系模型
│   │
│   ├── services/                      # 业务逻辑层（Service 层）
│   │   ├── manager.go                 # 服务管理器（统一管理所有服务）
│   │   ├── user_service.go            # 用户服务（登录、注册、资料管理）
│   │   ├── auction_service.go         # 拍卖服务（CRUD、状态管理）
│   │   ├── bid_service.go             # 出价服务（出价记录、价格转换）
│   │   ├── nft_service.go             # NFT 服务（同步、查询、验证）
│   │   ├── auction_task_scheduler.go  # 拍卖任务调度器（定时结束拍卖）
│   │   └── listener_service.go        # 区块链事件监听服务（监听合约事件）
│   │
│   ├── ethereum/                      # 以太坊相关封装
│   │   ├── client.go                  # 以太坊客户端封装（RPC 调用）
│   │   └── etherscan.go               # Etherscan API 客户端（用于查询交易）
│   │
│   ├── contracts/                     # 智能合约相关（ABI、绑定代码）
│   │   ├── my_auction/                # 拍卖合约
│   │   │   ├── my_auction.go          # Go 绑定代码（自动生成）
│   │   │   ├── MyXAuctionV2.abi       # 合约 ABI
│   │   │   └── MyXAuctionV2.bin       # 合约字节码
│   │   ├── erc721_nft/                # ERC721 NFT 合约绑定
│   │   └── erc20_metadata/            # ERC20 代币合约绑定
│   │
│   ├── router/                        # 路由配置
│   │   └── router.go                  # 路由注册、中间件配置
│   │
│   ├── websocket/                     # WebSocket 服务
│   │   ├── hub.go                     # WebSocket Hub（管理连接）
│   │   ├── handler.go                 # WebSocket 处理器
│   │   └── message.go                 # 消息类型定义
│   │
│   ├── jwt/                           # JWT 认证
│   │   └── jwt.go                     # JWT 生成、验证、解析
│   │
│   ├── logger/                        # 日志工具
│   │   └── logger.go                  # 日志初始化（zerolog 配置）
│   │
│   ├── response/                      # 统一响应格式
│   │   └── response.go                # 响应封装函数
│   │
│   ├── page/                          # 分页工具
│   │   └── pagination.go              # 分页查询参数绑定和处理
│   │
│   ├── utils/                         # 工具函数
│   │   ├── ethereum.go                # 以太坊相关工具（地址验证、金额转换）
│   │   └── signature.go               # 签名验证工具（钱包签名验证）
│   │
│   ├── validator/                     # 验证器
│   │   └── validator.go               # 自定义验证规则
│   │
│   └── errors/                        # 错误定义
│       └── errors.go                  # 自定义错误类型
│
├── sql/                               # SQL 脚本
│   └── db.sql                         # 数据库初始化脚本（表结构、索引）
│
├── docs/                              # API 文档（Swagger 生成）
│   ├── docs.go                        # Swagger 注解代码
│   ├── swagger.json                   # Swagger JSON 格式文档
│   ├── swagger.yaml                   # Swagger YAML 格式文档
│   ├── AMOUNT_STORAGE_DESIGN.md       # 金额存储设计文档
│   └── AUCTION_TASK_SCHEDULER.md      # 拍卖任务调度器设计文档
│
├── test/                              # 测试文件
│   └── main.go                        # 手动测试脚本（合约调用测试）
│
├── config.yaml.example                # 配置文件模板
├── config.yaml                        # 配置文件（本地，不提交到 Git）
├── go.mod                             # Go 模块定义
├── go.sum                             # 依赖版本锁定文件
├── .gitignore                         # Git 忽略文件配置
└── README.md                          # 项目说明文档（本文件）
```

---

## ⚙️ 核心服务说明

### ServiceManager（服务管理器）

统一管理所有业务服务，负责服务的初始化和依赖注入：

- **AuctionService**: 拍卖业务逻辑
- **BidService**: 出价业务逻辑
- **NFTService**: NFT 管理业务逻辑
- **UserService**: 用户业务逻辑
- **ListenerService**: 区块链事件监听服务
- **AuctionTaskScheduler**: 拍卖任务调度器
- **WSHub**: WebSocket Hub（实时消息推送）

### 主要服务功能

#### UserService（用户服务）
- 钱包登录：生成 nonce、验证签名、生成 JWT token
- 用户资料管理：查询、更新用户信息
- 平台统计：获取用户、拍卖、出价统计数据

#### AuctionService（拍卖服务）
- 拍卖 CRUD：创建、查询、更新、取消拍卖
- 状态管理：pending → active → ended/cancelled
- 链上交互：调用智能合约创建拍卖
- 任务调度：创建拍卖结束定时任务

#### BidService（出价服务）
- 出价记录：查询出价历史
- 价格转换：代币数量转 USD（使用 Chainlink 价格预言机）
- 出价验证：验证出价金额是否有效

#### NFTService（NFT 服务）
- NFT 同步：从区块链扫描用户 NFT 并保存到数据库
- NFT 查询：查询用户拥有的 NFT
- 所有权验证：验证用户是否拥有指定 NFT
- 元数据获取：从 IPFS/HTTP 获取 NFT 元数据

#### ListenerService（事件监听服务）
- 监听合约事件：AuctionCreated、BidPlaced、AuctionEnded 等
- 数据同步：将链上事件同步到数据库
- 消息推送：通过 WebSocket 推送事件通知
- 重连机制：自动重连 WebSocket 连接

#### AuctionTaskScheduler（任务调度器）
- 任务调度：使用 Asynq 调度拍卖结束任务
- 定时执行：在拍卖结束时间执行结算逻辑
- 任务管理：查询、取消任务状态
- 可靠性保证：基于 Redis 的可靠任务队列

---

## 🚀 快速开始

### 环境要求

- **Go**: 1.25 或更高版本
- **MySQL**: 8.0 或更高版本
- **Redis**: 6.0 或更高版本
- **以太坊节点**: 本地节点或 RPC 服务（如 Infura、Alchemy）

### 1. 安装依赖

```bash
go mod download
```

### 2. 数据库初始化

#### 方式一：使用 SQL 脚本（推荐）

```bash
# 创建数据库
mysql -u root -p -e "CREATE DATABASE auction_market_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 导入表结构
mysql -u root -p auction_market_db < sql/db.sql
```

#### 方式二：使用 GORM 自动迁移

```bash
# 创建数据库
mysql -u root -p -e "CREATE DATABASE auction_market_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 运行应用（会自动迁移表结构）
go run cmd/server/main.go
```

**注意**：GORM 自动迁移会在 `internal/server/server.go` 中执行，首次启动时会自动创建表结构。

### 3. 配置文件

**重要提示**：项目中的 `config.yaml` 文件已加入 `.gitignore`，不会被提交到代码仓库。

首次配置请按以下步骤：

```bash
# 复制配置文件模板
cp config.yaml.example config.yaml

# 编辑配置文件，填入实际的配置信息
# Windows: notepad config.yaml
# Linux/Mac: vim config.yaml 或 nano config.yaml
```

#### 必须配置的敏感信息：

1. **数据库配置**：
```yaml
database:
  host: localhost          # 数据库主机地址
  port: 3306              # 数据库端口
  user: root              # 数据库用户名
  password: YOUR_DATABASE_PASSWORD  # 数据库密码（请替换）
  name: auction_market_db # 数据库名称
  charset: utf8mb4        # 字符集
  max_open_conns: 25      # 最大打开连接数
  max_idle_conns: 10      # 最大空闲连接数
  conn_max_lifetime: 5m   # 连接最大生存时间
  conn_max_idle_time: 10m # 连接最大空闲时间
```

2. **JWT 配置**：
```yaml
jwt:
  secret: YOUR_JWT_SECRET_KEY_CHANGE_IN_PRODUCTION  # JWT 密钥（请使用强随机字符串）
  expiration: 24h  # Token 过期时间
```

3. **以太坊配置**（核心配置）：
```yaml
ethereum:
  rpc_url: https://your-rpc-provider.com/YOUR_API_KEY    # RPC 节点 URL
  wss_url: wss://your-wss-provider.com/YOUR_API_KEY      # WebSocket 节点 URL（用于事件监听）
  auction_contract_address: 0xYOUR_AUCTION_CONTRACT_ADDRESS  # 拍卖合约地址
  platform_private_key: YOUR_PLATFORM_PRIVATE_KEY        # ⚠️ 平台私钥（用于签名交易，请妥善保管）
  chain_id: 11155111      # 链 ID（Sepolia: 11155111, Mainnet: 1）
  websocket_timeout: 60s  # WebSocket 连接超时时间
```

4. **Etherscan API Key**（可选，用于查询交易）：
```yaml
etherscan:
  api_key: YOUR_ETHERSCAN_API_KEY  # Etherscan API Key（可选）
  chain_id: 11155111
```

5. **Redis 配置**（用于缓存和任务队列）：
```yaml
redis:
  addr: localhost:6379              # Redis 地址
  password: YOUR_REDIS_PASSWORD     # Redis 密码（如果无密码可留空）
  db: 0                             # 数据库编号
  pool_size: 10                     # 连接池大小
  min_idle_conns: 5                 # 最小空闲连接数
  dial_timeout: 5s                  # 连接超时时间
  read_timeout: 3s                  # 读取超时时间
  write_timeout: 3s                 # 写入超时时间
```

### 4. 启动 Redis（如果未启动）

```bash
# 使用 Docker 启动 Redis
docker run -d -p 6379:6379 redis:latest

# 或使用本地安装的 Redis
redis-server
```

### 5. 运行应用

```bash
# 开发模式运行
go run cmd/server/main.go

# 或编译后运行
go build -o auction-api cmd/server/main.go
./auction-api
```

应用将在 `http://localhost:8080` 启动。

### 6. 验证服务

- **健康检查**: `curl http://localhost:8080/api/health`
- **Swagger 文档**: 浏览器访问 `http://localhost:8080/api/swagger/index.html`
- **以太坊配置**: `curl http://localhost:8080/api/config/ethereum`

---

## 📚 API 功能说明

### 基础接口

#### 健康检查
- `GET /api/health` - 健康检查接口，用于检查 API 服务是否正常运行

#### 配置信息
- `GET /api/config/ethereum` - 获取以太坊网络配置信息（RPC URL、合约地址、链 ID 等）
  - **返回**: 拍卖合约地址、RPC URL、链 ID（用于前端配置）

### 认证相关

#### 钱包登录（推荐）

**请求 nonce**：
- `POST /api/auth/wallet/request-nonce`
  - **请求体**: `{ "walletAddress": "0x..." }`
  - **返回**: `{ "nonce": "...", "message": "..." }`
  - **说明**: 为指定钱包地址生成 nonce，用户需要使用钱包签名此 nonce

**验证签名并登录**：
- `POST /api/auth/wallet/verify`
  - **请求体**: `{ "walletAddress": "0x...", "signature": "0x...", "nonce": "..." }`
  - **返回**: `{ "token": "...", "user": { ... } }`
  - **说明**: 验证钱包签名，验证通过后返回 JWT token 和用户信息

**注意**：所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer <token>`。

### 用户相关

#### 用户信息（需要认证）
- `GET /api/users/profile` - 获取当前用户资料信息
- `PUT /api/users/profile` - 更新用户资料（用户名、邮箱）
  - **请求体**: `{ "username": "...", "email": "..." }`

#### 平台统计（公开接口）
- `GET /api/users/stats` - 获取平台统计数据
  - **返回**: 总用户数、总拍卖数、总出价数

### 拍卖相关

#### 拍卖列表（公开接口）
- `GET /api/auctions` - 获取拍卖列表（支持分页）
  - **查询参数**: `page`, `pageSize`
  
- `GET /api/auctions/public` - 获取公开拍卖列表（首页专用，按状态和时间排序）
  - **查询参数**: `page`, `pageSize`, `status` (active/ended/all)
  - **说明**: 返回按状态排序的拍卖列表（active 在前，ended 在后）

#### 拍卖详情（公开接口）
- `GET /api/auctions/:id` - 获取拍卖基本信息（通过拍卖 ID 字符串）
- `GET /api/auctions/:id/detail` - 获取拍卖详细信息（包含卖家钱包地址，通过数字 ID）

#### 拍卖管理（需要认证）
- `POST /api/auctions` - 创建新拍卖
  - **请求体**: NFT 地址、Token ID、起拍价、支付代币、开始/结束时间等
  - **说明**: 会在链上创建拍卖，并调度结束任务
  
- `PUT /api/auctions/:id` - 更新拍卖信息（仅限 pending 状态的拍卖）
  - **请求体**: 可更新的拍卖字段
  
- `POST /api/auctions/:id/cancel` - 取消拍卖（仅限 pending 或 active 状态的拍卖）

- `GET /api/auctions/my` - 获取我创建的拍卖列表
  - **查询参数**: `page`, `pageSize`, `status` (支持多个状态筛选，如 `?status=pending&status=active`)

- `GET /api/auctions/my/history` - 获取我的拍卖历史记录（简化字段）

#### 拍卖工具接口（公开接口）
- `GET /api/auctions/stats` - 获取拍卖简单统计数据（从合约获取）
  - **返回**: 总拍卖数、总出价数、平台费用、锁定总价值（TVL）

- `GET /api/auctions/nfts` - 获取所有拍卖中的 NFT 列表（去重，支持分页）

- `GET /api/auctions/supported-tokens` - 获取平台支持的支付代币列表
  - **返回**: 代币地址、符号、名称等信息

- `GET /api/auctions/token-price/:token` - 获取代币价格（USD，通过 Chainlink 价格预言机）
  - **路径参数**: 代币合约地址

- `POST /api/auctions/convert-to-usd` - 将代币数量转换为 USD 价值
  - **请求体**: `{ "token": "0x...", "amount": "..." }`

- `POST /api/auctions/check-nft-approval` - 检查 NFT 是否已授权给平台合约
  - **请求体**: `{ "nftAddress": "0x...", "tokenId": "..." }`

### 出价相关

#### 出价查询（公开接口）
- `GET /api/auctions/:auctionId/bids` - 获取指定拍卖的所有出价记录（支持分页）
  - **查询参数**: `page`, `pageSize`
  - **返回**: 出价者钱包地址和出价信息（不包含完整用户信息）

- `GET /api/bids/:id` - 获取单个出价详情（通过出价 ID）

**注意**：出价功能通常在链上直接进行，前端调用智能合约出价，后端通过监听合约事件同步到数据库。

### NFT 相关（需要认证）

#### NFT 查询
- `GET /api/nfts/my` - 获取我拥有的 NFT 列表（从数据库查询，支持分页）
- `GET /api/nfts/my/list` - 获取我拥有的 NFT 完整列表（不分页，所有合约）

- `GET /api/nfts/:id` - 根据 NFT ID 获取 NFT 详情
- `GET /api/nfts/my/ownership/:nftId` - 获取我的 NFT 所有权记录（通过 nftId）

#### NFT 同步
- `POST /api/nfts/sync` - 同步用户 NFT（从区块链同步到数据库）
  - **功能**: 扫描用户钱包地址，获取所有 ERC721 NFT 并保存到数据库
  - **说明**: 这是一个异步操作，可能需要一些时间
  
- `GET /api/nfts/sync/status` - 获取 NFT 同步状态

#### NFT 验证
- `POST /api/nfts/verify` - 验证用户是否拥有指定 NFT
  - **请求体**: `{ "nftAddress": "0x...", "tokenId": "..." }`

### 拍卖任务调度（需要认证）

#### 任务管理
- `GET /api/auction-tasks/:auctionId` - 获取拍卖结束任务的调度状态
- `DELETE /api/auction-tasks/:auctionId` - 取消拍卖结束任务的调度

**说明**：系统会自动调度拍卖结束任务，在拍卖结束时执行结算逻辑。此接口用于查询和管理这些任务。

### WebSocket

#### 实时通信
- `GET /api/ws` - WebSocket 连接（公开）
- `GET /api/ws/auth` - WebSocket 连接（需要认证）

**功能**：用于实时推送拍卖更新、出价通知等消息。支持订阅特定拍卖的事件。

**消息类型**：
- `auction_created`: 拍卖创建
- `auction_bid_placed`: 新出价
- `auction_ended`: 拍卖结束
- `auction_cancelled`: 拍卖取消
- `nft_approved`: NFT 授权成功

**订阅机制**：
- 发送 `subscribe` 消息订阅特定拍卖
- 发送 `unsubscribe` 消息取消订阅

### API 响应格式

所有 API 响应遵循统一格式：

```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

分页接口的响应格式：

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "items": [ ... ],
    "pagination": {
      "page": 1,
      "pageSize": 10,
      "total": 100,
      "totalPages": 10
    }
  }
}
```

### 错误码说明

- `200` - 成功
- `201` - 创建成功
- `400` - 请求参数错误
- `401` - 未授权（需要登录）
- `403` - 禁止访问（权限不足）
- `404` - 资源不存在
- `500` - 服务器内部错误

---

## 📊 数据库设计

### 主要数据表

#### users (用户表)
```sql
- id: BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT
- username: VARCHAR(255) UNIQUE           # 用户名（唯一）
- email: VARCHAR(255) UNIQUE              # 邮箱（唯一）
- password: VARCHAR(255) NOT NULL         # 密码哈希（钱包登录用户为空字符串）
- wallet_address: VARCHAR(255) UNIQUE     # 钱包地址（唯一，索引）
- created_at: TIMESTAMP                   # 创建时间
- updated_at: TIMESTAMP                   # 更新时间
```

#### auctions (拍卖表)
```sql
- id: BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT
- user_id: BIGINT UNSIGNED                # 创建者 ID（外键：users.id）
- auction_id: VARCHAR(255) UNIQUE         # 拍卖 ID（字符串，唯一，索引）
- contract_auction_id: BIGINT UNSIGNED    # 合约中的拍卖 ID
- nft_address: VARCHAR(255)               # NFT 合约地址
- token_id: VARCHAR(255)                  # Token ID
- start_price: DECIMAL(65,0)              # 起拍价（Wei 单位）
- start_price_usd: DECIMAL(20,8)          # 起拍价（USD）
- payment_token: VARCHAR(255)             # 支付代币地址
- start_time: TIMESTAMP                   # 开始时间（索引）
- end_time: TIMESTAMP                     # 结束时间（索引）
- status: VARCHAR(50)                     # 状态：pending/active/ended/cancelled（索引）
- highest_bid: DECIMAL(65,0)              # 最高出价
- highest_bidder: VARCHAR(255)            # 最高出价者地址
- bid_count: INT UNSIGNED DEFAULT 0       # 出价次数
- created_at: TIMESTAMP                   # 创建时间
- updated_at: TIMESTAMP                   # 更新时间
```

#### bids (出价表)
```sql
- id: BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT
- auction_id: BIGINT UNSIGNED             # 拍卖 ID（外键：auctions.id，索引）
- user_id: BIGINT UNSIGNED                # 出价者 ID（外键：users.id，索引）
- amount: DECIMAL(65,0)                   # 出价金额（Wei 单位）
- amount_usd: DECIMAL(20,8)               # 出价金额（USD）
- payment_token: VARCHAR(255)             # 支付代币地址
- transaction_hash: VARCHAR(255)          # 交易哈希（索引）
- block_number: BIGINT UNSIGNED           # 区块号
- is_highest: BOOLEAN DEFAULT FALSE       # 是否为最高出价
- created_at: TIMESTAMP                   # 创建时间（索引，用于排序）
```

#### nfts (NFT 表)
```sql
- id: BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT
- nft_id: VARCHAR(255) UNIQUE             # NFT ID（字符串，唯一，索引）
- contract_address: VARCHAR(255)          # 合约地址（索引）
- token_id: VARCHAR(255)                  # Token ID
- name: VARCHAR(500)                      # NFT 名称
- image: TEXT                             # 图片 URL
- description: TEXT                       # 描述
- metadata: JSON                          # 元数据（JSON 格式）
- created_at: TIMESTAMP                   # 创建时间
- updated_at: TIMESTAMP                   # 更新时间
```

#### nft_ownerships (NFT 所有权表)
```sql
- id: BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT
- user_id: BIGINT UNSIGNED                # 用户 ID（外键：users.id，索引）
- nft_id: BIGINT UNSIGNED                 # NFT ID（外键：nfts.id，索引）
- token_id: VARCHAR(255)                  # Token ID
- contract_address: VARCHAR(255)          # 合约地址（索引）
- created_at: TIMESTAMP                   # 创建时间
- updated_at: TIMESTAMP                   # 更新时间
- UNIQUE KEY (user_id, nft_id)            # 唯一索引：一个用户对一个 NFT 只有一条记录
```

### 索引设计

- **用户表**: wallet_address（唯一索引）
- **拍卖表**: auction_id（唯一索引）、status、start_time、end_time
- **出价表**: auction_id、user_id、transaction_hash、created_at
- **NFT 表**: nft_id（唯一索引）、contract_address
- **所有权表**: user_id、nft_id、（user_id, nft_id）唯一索引

---

## 🔐 安全注意事项

### 配置文件安全

本项目已配置 `.gitignore` 排除敏感配置文件，但仍需注意：

1. **如果 `config.yaml` 已被提交到 Git 历史中**，需要从历史记录中删除：
   ```bash
   # 使用 git filter-branch 从历史中删除敏感文件
   git filter-branch --force --index-filter \
     "git rm --cached --ignore-unmatch config.yaml" \
     --prune-empty --tag-name-filter cat -- --all
   ```

2. **生产环境建议**：
   - 使用环境变量管理敏感配置
   - 使用密钥管理服务（如 AWS Secrets Manager、HashiCorp Vault 等）
   - 定期轮换 JWT secret 和 API keys
   - 使用不同的密钥和配置用于开发、测试、生产环境

3. **平台私钥安全**：
   - `platform_private_key` 用于签名链上交易，一旦泄露可能导致资产损失
   - 建议使用硬件钱包或多签方案管理平台私钥
   - 不要在代码中硬编码私钥
   - 使用专用的最小权限钱包地址
   - 定期检查钱包余额和交易记录

### 数据库安全

- 使用强密码
- 限制数据库访问 IP（只允许应用服务器访问）
- 定期备份数据库
- 使用 SSL/TLS 连接（生产环境）
- 定期更新数据库版本和安全补丁

### API 安全

- 使用 HTTPS（生产环境）
- 实施请求限流（防止 DDoS）
- 验证所有输入参数
- 使用 CORS 限制跨域访问
- 定期更新依赖包（安全漏洞修复）

### JWT 安全

- 使用强随机字符串作为 secret（至少 32 位）
- 设置合理的过期时间
- 在生产环境使用 HTTPS（防止 token 被窃取）
- 考虑实现 token 刷新机制

---

## 🔧 配置说明

### 环境变量

可以通过环境变量覆盖配置：

- `CONFIG_PATH` - 配置文件路径（默认: `config.yaml`）

```bash
export CONFIG_PATH=/path/to/config.yaml
go run cmd/server/main.go
```

### 配置文件结构

详细配置说明请参考 `config.yaml.example` 文件。主要配置项：

- **应用配置**: 应用名称、环境、端口、超时时间、日志级别
- **数据库配置**: 连接信息、连接池配置
- **JWT 配置**: Secret、过期时间
- **以太坊配置**: RPC URL、WSS URL、合约地址、私钥、链 ID
- **Etherscan 配置**: API Key、链 ID
- **Redis 配置**: 地址、密码、连接池配置

---

## 🧪 开发指南

### 生成 Swagger 文档

```bash
# 安装 swag
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init -g cmd/server/main.go

# 文档会生成到 docs/ 目录
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/services/...

# 显示测试覆盖率
go test -cover ./...
```

### 运行测试文件

如果需要运行 `test/main.go` 进行手动测试（合约调用测试），需要设置 RPC URL 环境变量：

```bash
# Windows
set ETH_RPC_URL=https://your-rpc-provider.com/YOUR_API_KEY
go run test/main.go

# Linux/Mac
export ETH_RPC_URL=https://your-rpc-provider.com/YOUR_API_KEY
go run test/main.go
```

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 使用 `golint` 或 `golangci-lint` 检查代码
- 编写单元测试和集成测试

### 日志

系统使用 `zerolog` 作为日志库，支持结构化日志：

```go
logger.Info("message", "key", "value")
logger.Error("message", "error", err)
logger.Debug("message")
```

日志级别可通过配置文件设置：`debug`, `info`, `warn`, `error`

---

## 🚀 部署指南

### 编译

```bash
# 编译为可执行文件
go build -o auction-api cmd/server/main.go

# 交叉编译（Linux）
GOOS=linux GOARCH=amd64 go build -o auction-api cmd/server/main.go

# 交叉编译（Windows）
GOOS=windows GOARCH=amd64 go build -o auction-api.exe cmd/server/main.go
```

### 使用 systemd（Linux）

创建服务文件 `/etc/systemd/system/auction-api.service`：

```ini
[Unit]
Description=Auction Market API
After=network.target mysql.service redis.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/auction-api
ExecStart=/opt/auction-api/auction-api
Restart=always
RestartSec=5
Environment="CONFIG_PATH=/opt/auction-api/config.yaml"

# 日志
StandardOutput=journal
StandardError=journal
SyslogIdentifier=auction-api

[Install]
WantedBy=multi-user.target
```

启动服务：
```bash
sudo systemctl enable auction-api
sudo systemctl start auction-api
sudo systemctl status auction-api

# 查看日志
sudo journalctl -u auction-api -f
```

### 使用 Docker

创建 `Dockerfile`：

```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o auction-api cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/auction-api .
COPY --from=builder /app/config.yaml.example .
EXPOSE 8080
CMD ["./auction-api"]
```

构建和运行：
```bash
docker build -t auction-api .
docker run -d -p 8080:8080 \
  -v $(pwd)/config.yaml:/root/config.yaml \
  --name auction-api \
  auction-api
```

### 使用 Docker Compose

创建 `docker-compose.yml`：

```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./config.yaml:/root/config.yaml
    depends_on:
      - mysql
      - redis
    environment:
      - CONFIG_PATH=/root/config.yaml

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: auction_market_db
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  mysql_data:
```

启动：
```bash
docker-compose up -d
```

### 生产环境建议

1. **使用反向代理**（Nginx、Caddy 等）
2. **启用 HTTPS**（使用 Let's Encrypt 免费证书）
3. **配置防火墙**（只开放必要端口）
4. **监控和日志**（使用 Prometheus、Grafana 等）
5. **备份策略**（定期备份数据库）
6. **负载均衡**（如需要，使用多个实例）

---

## 📚 相关文档

- [项目总结文档](../my-auction-market-front/PROJECT_SUMMARY.md) - 详细的项目文档（包含前后端）
- [金额存储设计文档](./docs/AMOUNT_STORAGE_DESIGN.md) - 金额存储设计说明
- [拍卖任务调度器文档](./docs/AUCTION_TASK_SCHEDULER.md) - 任务调度器设计说明
- [Swagger API 文档](http://localhost:8080/api/swagger/index.html) - 交互式 API 文档（需运行服务）

---

## 🤝 贡献指南

在提交代码前，请确保：

1. ✅ 所有敏感信息已从代码中移除
2. ✅ `config.yaml` 未被提交（已在 .gitignore 中）
3. ✅ 使用 `config.yaml.example` 作为配置模板
4. ✅ 代码通过了测试和 lint 检查
5. ✅ 遵循 Go 代码规范
6. ✅ 添加了必要的注释和文档

### 提交流程

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

---

## 📄 许可证

MIT License

---

**最后更新**: 2025年
