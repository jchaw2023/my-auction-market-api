# 以太坊金额存储设计方案对比

## 当前设计问题

当前设计中存在不一致：
- `StartPrice`: `float64` (ETH) - 存储为 ETH
- `HighestBid`: `BigInt` (wei) - 存储为 wei  
- `Amount`: `BigInt` (wei) - 存储为 wei

这种混合设计会导致：
1. 代码复杂度增加（需要频繁转换）
2. 容易出错（单位混淆）
3. 统计查询困难（需要统一单位）

## 方案对比

### 方案 1: 统一存储 ETH（推荐 ⭐⭐⭐⭐⭐）

**设计思路：**
- 所有金额统一存储为 ETH（使用 `decimal(20,8)`）
- 与合约交互时转换为 wei
- 前端直接使用，无需转换

**优点：**
- ✅ 统一单位，代码简洁
- ✅ 易于理解和维护
- ✅ 统计查询简单
- ✅ 前端友好（直接显示 ETH）
- ✅ 存储空间小（decimal(20,8) vs decimal(78,0)）

**缺点：**
- ⚠️ 需要与合约交互时转换（但这是必要的）
- ⚠️ 精度限制（但 8 位小数对 ETH 足够）

**实现：**
```go
type Auction struct {
    StartPrice    decimal.Decimal `json:"startPrice" gorm:"type:decimal(20,8);not null;comment:起拍价(ETH)"`
    HighestBid    decimal.Decimal `json:"highestBid" gorm:"type:decimal(20,8);default:0;comment:最高出价(ETH)"`
}

type Bid struct {
    Amount decimal.Decimal `json:"amount" gorm:"type:decimal(20,8);not null;comment:出价金额(ETH)"`
}

// 与合约交互时转换
func (a *Auction) StartPriceWei() *big.Int {
    return ethers.ParseEther(a.StartPrice.String())
}
```

**适用场景：**
- ✅ 应用层主要使用 ETH 单位
- ✅ 需要频繁统计和查询
- ✅ 前端直接显示 ETH

---

### 方案 2: 统一存储 Wei（当前部分实现）

**设计思路：**
- 所有金额统一存储为 wei（使用 `decimal(78,0)` 或 `BigInt`）
- 显示时转换为 ETH

**优点：**
- ✅ 与链上数据一致，无需转换
- ✅ 精度完整（无精度损失）
- ✅ 适合需要精确计算的场景

**缺点：**
- ❌ 存储空间大
- ❌ 统计查询复杂（大数运算）
- ❌ 前端需要转换
- ❌ 代码复杂度高

**适用场景：**
- ✅ 需要与链上数据完全一致
- ✅ 需要精确到 wei 的计算
- ⚠️ 统计需求较少

---

### 方案 3: 双字段存储（推荐 ⭐⭐⭐⭐）

**设计思路：**
- 同时存储 wei 和 ETH
- wei 用于与合约交互
- ETH 用于显示和统计

**优点：**
- ✅ 兼顾精确性和易用性
- ✅ 统计查询使用 ETH（快速）
- ✅ 合约交互使用 wei（精确）
- ✅ 数据冗余提供校验

**缺点：**
- ⚠️ 存储空间增加（但可接受）
- ⚠️ 需要保持两个字段同步

**实现：**
```go
type Auction struct {
    // 用于合约交互（精确）
    StartPriceWei    BigInt    `json:"startPriceWei" gorm:"type:decimal(78,0);not null;comment:起拍价(wei)"`
    HighestBidWei    BigInt    `json:"highestBidWei" gorm:"type:decimal(78,0);default:0;comment:最高出价(wei)"`
    
    // 用于显示和统计（易用）
    StartPrice       decimal.Decimal `json:"startPrice" gorm:"type:decimal(20,8);not null;comment:起拍价(ETH)"`
    HighestBid       decimal.Decimal `json:"highestBid" gorm:"type:decimal(20,8);default:0;comment:最高出价(ETH)"`
}

// 自动同步方法
func (a *Auction) SetStartPriceWei(wei *big.Int) {
    a.StartPriceWei = NewBigInt(wei)
    a.StartPrice = decimal.NewFromBigInt(wei, -18) // wei to ETH
}
```

**适用场景：**
- ✅ 需要精确的链上交互
- ✅ 需要频繁的统计查询
- ✅ 对存储空间不敏感

---

### 方案 4: 使用 Gwei 作为中间单位（不推荐）

**设计思路：**
- 使用 Gwei（10^9 wei）作为存储单位
- 介于 wei 和 ETH 之间

**优点：**
- ✅ 比 wei 小，比 ETH 精确

**缺点：**
- ❌ 增加转换复杂度
- ❌ 不是标准做法
- ❌ 仍然需要转换

**适用场景：**
- ❌ 不推荐使用

---

## 推荐方案

### 对于你的项目，推荐：**方案 1（统一存储 ETH）**

**理由：**
1. **当前 StartPrice 已经是 ETH**，统一更简单
2. **统计需求**：使用 ETH 统计更直观和高效
3. **前端友好**：直接显示，无需转换
4. **代码简洁**：减少单位转换代码
5. **精度足够**：decimal(20,8) 对 ETH 完全够用

**迁移建议：**
1. 将 `HighestBid` 和 `Amount` 改为 `decimal.Decimal` 类型
2. 与合约交互时统一转换：
   ```go
   // 从合约读取
   wei := contract.GetAmount()
   amount := decimal.NewFromBigInt(wei, -18) // wei to ETH
   
   // 发送到合约
   wei := amount.Mul(decimal.NewFromInt(1e18)).BigInt()
   contract.SetAmount(wei)
   ```

---

## 使用 decimal.Decimal 库

推荐使用 `shopspring/decimal` 库：

```go
import "github.com/shopspring/decimal"

// 安装
go get github.com/shopspring/decimal
```

**优点：**
- ✅ 精确的十进制计算
- ✅ 支持 GORM（实现 Scanner/Valuer）
- ✅ 性能好
- ✅ API 友好

**示例：**
```go
type Auction struct {
    StartPrice decimal.Decimal `gorm:"type:decimal(20,8)"`
}

// 创建
price := decimal.NewFromFloat(1.5) // 1.5 ETH

// 计算
total := price1.Add(price2)
if price1.GreaterThan(price2) { ... }

// 与 big.Int 转换
wei := price.Mul(decimal.NewFromInt(1e18)).BigInt()
price := decimal.NewFromBigInt(wei, -18)
```

---

## 总结

| 方案 | 存储类型 | 优点 | 缺点 | 推荐度 |
|------|---------|------|------|--------|
| **统一 ETH** | decimal(20,8) | 简单、高效、易用 | 需要转换 | ⭐⭐⭐⭐⭐ |
| **统一 Wei** | decimal(78,0) | 精确、一致 | 复杂、低效 | ⭐⭐⭐ |
| **双字段** | 两者都有 | 兼顾 | 冗余 | ⭐⭐⭐⭐ |
| **Gwei** | decimal | - | 不标准 | ⭐ |

**最终建议：使用方案 1（统一存储 ETH）+ `shopspring/decimal` 库**

