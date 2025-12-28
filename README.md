# My Auction Market API

NFT 拍卖商城后端 API，基于 Gin 和 GORM 框架，整合以太坊客户端用于与智能合约交互。

## 项目特性

- RESTful API 设计
- JWT 身份认证
- 用户注册/登录
- NFT 拍卖管理
- 出价功能
- 与以太坊智能合约集成
- MySQL 数据库
- Swagger API 文档

## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **以太坊**: go-ethereum (ethclient)
- **认证**: JWT
- **日志**: zerolog
- **配置**: YAML

## 项目结构

```
my-auction-market-api/
├── cmd/
│   └── server/
│       └── main.go          # 应用入口
├── internal/
│   ├── config/              # 配置管理
│   ├── database/            # 数据库连接和 GORM 配置
│   ├── ethereum/            # 以太坊客户端封装
│   ├── handlers/            # HTTP 处理器
│   ├── jwt/                 # JWT 认证
│   ├── logger/              # 日志工具
│   ├── middleware/          # 中间件
│   ├── models/              # 数据模型
│   ├── page/                # 分页工具
│   ├── response/            # 响应格式
│   ├── router/              # 路由配置
│   ├── server/              # 服务器配置
│   ├── services/            # 业务逻辑
│   └── validator/           # 验证器
├── config.yaml.example      # 配置文件模板（请复制为 config.yaml 并填写实际配置）
├── config.yaml              # 配置文件（已加入 .gitignore，不会被提交）
├── sql/                      # 数据库脚本
│   ├── schema.sql           # 数据库表结构
│   ├── init.sql             # 初始化脚本
│   ├── drop.sql             # 删除表脚本
│   └── README.md            # 数据库说明文档
├── go.mod                    # Go 模块
└── README.md                 # 项目说明
```

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置文件

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
  host: localhost
  port: 3306
  user: root
  password: YOUR_DATABASE_PASSWORD  # 请替换为实际数据库密码
  name: auction_market_db
```

2. **JWT 密钥**：
```yaml
jwt:
  secret: YOUR_JWT_SECRET_KEY_CHANGE_IN_PRODUCTION  # 请使用强随机字符串
  expiration: 24h
```

3. **以太坊配置**：
```yaml
ethereum:
  rpc_url: https://your-rpc-provider.com/YOUR_API_KEY  # 请替换为实际的 RPC URL 和 API Key
  wss_url: wss://your-wss-provider.com/YOUR_API_KEY   # 请替换为实际的 WebSocket URL 和 API Key
  auction_contract_address: 0xYOUR_AUCTION_CONTRACT_ADDRESS  # 请替换为实际的合约地址
  platform_private_key: YOUR_PLATFORM_PRIVATE_KEY  # ⚠️ 敏感信息，请妥善保管
  chain_id: 11155111
```

4. **Etherscan API Key**（可选，用于合约交互）：
```yaml
etherscan:
  api_key: YOUR_ETHERSCAN_API_KEY  # 请替换为实际的 API Key
  chain_id: 11155111
```

5. **Redis 配置**（如果 Redis 设置了密码）：
```yaml
redis:
  addr: localhost:6379
  password: YOUR_REDIS_PASSWORD  # 如果 Redis 无密码可留空
  db: 0
```

**安全提示**：
- ⚠️ **永远不要**将包含真实敏感信息的 `config.yaml` 文件提交到代码仓库
- ⚠️ `platform_private_key` 是平台私钥，泄露会导致资产损失，请务必妥善保管
- 生产环境请使用环境变量或安全的密钥管理服务
- JWT secret 建议使用至少 32 位的随机字符串

### 4. 初始化数据库

有两种方式初始化数据库：

#### 方式一：使用 SQL 脚本（推荐）

```bash
# 执行 SQL 脚本创建数据库和表
mysql -u root -p < sql/schema.sql

# 或使用初始化脚本
mysql -u root -p < sql/init.sql
```

#### 方式二：使用 GORM 自动迁移

```bash
# 创建数据库
mysql -u root -p -e "CREATE DATABASE auction_market_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 运行应用（会自动迁移）
go run cmd/server/main.go
```

**注意**：GORM 自动迁移会在 `internal/server/server.go` 中执行，首次启动时会自动创建表结构。

### 5. 运行应用

```bash
go run cmd/server/main.go
```

应用将在 `http://localhost:8080` 启动。

## API 文档

启动应用后，访问 Swagger 文档：

```
http://localhost:8080/api/swagger/index.html
```

## API 功能说明

### 基础接口

#### 健康检查
- `GET /api/health` - 健康检查接口，用于检查 API 服务是否正常运行

#### 配置信息
- `GET /api/config/ethereum` - 获取以太坊网络配置信息（RPC URL、合约地址、链 ID 等）

### 认证相关

#### 钱包登录（推荐）
- `POST /api/auth/wallet/request-nonce` - 请求 nonce，用于钱包签名登录
  - 请求参数：`{ "walletAddress": "0x..." }`
  - 返回：nonce 字符串，用户需要签名此 nonce 用于验证

- `POST /api/auth/wallet/verify` - 验证钱包签名并登录
  - 请求参数：`{ "walletAddress": "0x...", "signature": "0x...", "nonce": "..." }`
  - 返回：JWT token 和用户信息

**注意**：所有需要认证的接口都需要在请求头中携带 `Authorization: Bearer <token>`。

### 用户相关

#### 用户信息（需要认证）
- `GET /api/users/profile` - 获取当前用户资料信息
- `PUT /api/users/profile` - 更新用户资料（用户名、邮箱）

#### 平台统计（公开接口）
- `GET /api/users/stats` - 获取平台统计数据（总用户数、总拍卖数、总出价数等）

### 拍卖相关

#### 拍卖列表（公开接口）
- `GET /api/auctions` - 获取拍卖列表（支持分页）
  - 查询参数：`page`, `pageSize`
  
- `GET /api/auctions/public` - 获取公开拍卖列表（首页专用，按状态和时间排序）
  - 查询参数：`page`, `pageSize`, `status` (active/ended/all)

#### 拍卖详情（公开接口）
- `GET /api/auctions/:id` - 获取拍卖基本信息（通过拍卖 ID 字符串）
- `GET /api/auctions/:id/detail` - 获取拍卖详细信息（包含卖家钱包地址，通过数字 ID）

#### 拍卖管理（需要认证）
- `POST /api/auctions` - 创建新拍卖
  - 请求体：NFT 地址、Token ID、起拍价、支付代币、开始/结束时间等
  
- `PUT /api/auctions/:id` - 更新拍卖信息（仅限 pending 状态的拍卖）
  
- `POST /api/auctions/:id/cancel` - 取消拍卖（仅限 pending 或 active 状态的拍卖）

- `GET /api/auctions/my` - 获取我创建的拍卖列表
  - 查询参数：`page`, `pageSize`, `status` (支持多个状态筛选)

- `GET /api/auctions/my/history` - 获取我的拍卖历史记录（简化字段）

#### 拍卖工具接口（公开接口）
- `GET /api/auctions/stats` - 获取拍卖简单统计数据（从合约获取：总拍卖数、总出价数、平台费用、锁定总价值）

- `GET /api/auctions/nfts` - 获取所有拍卖中的 NFT 列表（去重，支持分页）

- `GET /api/auctions/supported-tokens` - 获取平台支持的支付代币列表

- `GET /api/auctions/token-price/:token` - 获取代币价格（USD，通过 Chainlink 价格预言机）
  - 路径参数：代币合约地址

- `POST /api/auctions/convert-to-usd` - 将代币数量转换为 USD 价值
  - 请求体：`{ "token": "0x...", "amount": "..." }`

- `POST /api/auctions/check-nft-approval` - 检查 NFT 是否已授权给平台合约
  - 请求体：`{ "nftAddress": "0x...", "tokenId": "..." }`

### 出价相关

#### 出价查询（公开接口）
- `GET /api/auctions/:auctionId/bids` - 获取指定拍卖的所有出价记录（支持分页）
  - 查询参数：`page`, `pageSize`
  - 返回：出价者钱包地址和出价信息（不包含完整用户信息）

- `GET /api/bids/:id` - 获取单个出价详情（通过出价 ID）

**注意**：出价功能通常在链上直接进行，通过监听合约事件同步到数据库。

### NFT 相关（需要认证）

#### NFT 查询
- `GET /api/nfts/my` - 获取我拥有的 NFT 列表（从数据库查询，支持分页）
- `GET /api/nfts/my/list` - 获取我拥有的 NFT 完整列表（不分页，所有合约）

- `GET /api/nfts/:id` - 根据 NFT ID 获取 NFT 详情
- `GET /api/nfts/my/ownership/:nftId` - 获取我的 NFT 所有权记录（通过 nftId）

#### NFT 同步
- `POST /api/nfts/sync` - 同步用户 NFT（从区块链同步到数据库）
  - 功能：扫描用户钱包地址，获取所有 ERC721 NFT 并保存到数据库
  
- `GET /api/nfts/sync/status` - 获取 NFT 同步状态

#### NFT 验证
- `POST /api/nfts/verify` - 验证用户是否拥有指定 NFT
  - 请求体：`{ "nftAddress": "0x...", "tokenId": "..." }`

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

## 数据模型

### User (用户)

- ID
- Username (用户名)
- Email (邮箱)
- Password (密码哈希)
- WalletAddress (钱包地址)
- CreatedAt / UpdatedAt

### Auction (拍卖)

- ID
- UserID (创建者ID)
- ContractAuctionID (合约中的拍卖ID)
- NFTAddress (NFT合约地址)
- TokenID (Token ID)
- StartPrice (起拍价)
- StartPriceUSD (起拍价USD)
- PaymentToken (支付代币地址)
- StartTime / EndTime (开始/结束时间)
- Status (状态: active, ended, cancelled)
- HighestBid (最高出价)
- HighestBidder (最高出价者)
- BidCount (出价次数)

### Bid (出价)

- ID
- AuctionID (拍卖ID)
- UserID (出价者ID)
- Amount (出价金额)
- AmountUSD (出价金额USD)
- PaymentToken (支付代币)
- TransactionHash (交易哈希)
- BlockNumber (区块号)
- IsHighest (是否为最高出价)

## 安全注意事项

### 配置文件安全

本项目已配置 `.gitignore` 排除敏感配置文件，但仍需注意：

1. **如果 `config.yaml` 已被提交到 Git 历史中**，需要从历史记录中删除：
   ```bash
   # 使用 git filter-branch 或 BFG Repo-Cleaner 从历史中删除敏感文件
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

## 环境变量

可以通过环境变量覆盖配置：

- `CONFIG_PATH` - 配置文件路径（默认: `config.yaml`）

## 开发

### 生成 Swagger 文档

```bash
swag init -g cmd/server/main.go
```

### 运行测试

```bash
go test ./...
```

### 运行测试文件

如果需要运行 `test/main.go` 进行手动测试，需要设置 RPC URL 环境变量：

```bash
# Windows
set ETH_RPC_URL=https://your-rpc-provider.com/YOUR_API_KEY
go run test/main.go

# Linux/Mac
export ETH_RPC_URL=https://your-rpc-provider.com/YOUR_API_KEY
go run test/main.go
```

## 贡献指南

在提交代码前，请确保：

1. ✅ 所有敏感信息已从代码中移除
2. ✅ `config.yaml` 未被提交（已在 .gitignore 中）
3. ✅ 使用 `config.yaml.example` 作为配置模板
4. ✅ 代码通过了测试和 lint 检查

## 许可证

MIT License

