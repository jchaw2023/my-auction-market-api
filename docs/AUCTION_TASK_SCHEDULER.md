# 拍卖任务调度器框架使用说明

## 概述

这是一个基于 Asynq 的拍卖任务调度框架，用于处理拍卖结束时间（end_time）的一次性定时任务。

## 功能特性

1. ✅ **系统重启自动恢复**：系统启动时自动从数据库恢复未执行的拍卖任务
2. ✅ **动态添加任务**：创建拍卖时自动调度任务
3. ✅ **动态删除任务**：提供 API 接口删除任务（通过检查状态实现）
4. ✅ **精确时间触发**：基于 `end_time` 精确触发，只执行一次
5. ✅ **任务去重**：使用 `auction_id` 作为唯一标识，确保不重复

## 架构设计

### 任务标识

- **task_id**: 使用 `auction_id`（字符串），不重复
- **task_name**: 固定为 `"auctions"`，可能重复（多个拍卖使用相同名称）
- **params**: 包含 `user_id` 和 `nft_id`

### 任务负载结构

```json
{
  "task_id": "1234567890",
  "task_name": "auctions",
  "user_id": 1,
  "nft_id": "nft-123"
}
```

## 使用方式

### 1. 自动调度（创建拍卖时）

创建拍卖时会自动调度任务，无需手动调用：

```go
// 在 AuctionService.Create 中已自动集成
auction, err := auctionService.Create(userID, payload)
// 任务已自动调度，会在 end_time 时触发
```

### 2. 手动调度

```go
scheduler := services.GetAuctionTaskScheduler(cfg)
err := scheduler.ScheduleAuctionEndTask(
    auctionID,    // 拍卖ID
    endTime,      // 结束时间
    userID,       // 用户ID
    nftID,        // NFT ID
)
```

### 3. 取消任务

#### 方式1：通过 API

```bash
DELETE /api/auction-tasks/{auctionId}
Authorization: Bearer {token}
```

#### 方式2：在代码中

```go
scheduler := services.GetAuctionTaskScheduler(cfg)
err := scheduler.CancelAuctionEndTask(auctionID)
```

**注意**：Asynq 不支持直接取消延迟任务，取消逻辑在任务执行时检查拍卖状态。

### 4. 系统启动恢复

系统启动时会自动恢复所有未执行的拍卖任务：

```go
// 在 server.go 中已自动集成
s.serviceManager.StartAuctionTaskScheduler(ctx)
```

## 任务处理逻辑

### 任务执行流程

1. **任务触发**：在 `end_time` 时触发
2. **状态检查**：检查拍卖是否仍然有效
3. **时间验证**：验证结束时间是否已到（防止提前执行）
4. **执行处理**：更新拍卖状态为 `ended`
5. **错误处理**：如果失败会触发重试

### 任务处理器

任务处理器位于 `internal/services/auction_task_scheduler.go` 的 `handleAuctionEndTask` 方法中。

可以在此方法中添加自定义逻辑：

```go
func (s *AuctionTaskScheduler) processAuctionEnd(auction *models.Auction) error {
    // 1. 更新拍卖状态
    auction.Status = "ended"
    database.DB.Save(auction)
    
    // 2. 添加自定义逻辑
    // - 通知用户
    // - 处理退款
    // - 更新 NFT 状态
    // - 发送 WebSocket 消息
    // ...
    
    return nil
}
```

## 配置

### Redis 配置

在 `config.yaml` 中配置 Redis：

```yaml
redis:
  addr: localhost:6379
  password: ""
  db: 0
  pool_size: 10
  min_idle_conns: 5
  dial_timeout: 5s
  read_timeout: 3s
  write_timeout: 3s
```

### 任务队列配置

在 `auction_task_scheduler.go` 中配置：

```go
server := asynq.NewServer(redisConn, asynq.Config{
    Concurrency: 10, // 并发处理任务数
    Queues: map[string]int{
        "auctions": 10, // 拍卖任务队列
    },
})
```

## API 接口

### 取消拍卖任务

```
DELETE /api/auction-tasks/{auctionId}
```

**请求头**：
```
Authorization: Bearer {token}
```

**响应**：
```json
{
  "success": true,
  "data": {
    "message": "Task cancellation requested",
    "auctionId": "1234567890",
    "note": "Task will be checked when it executes. If auction is cancelled, task will be skipped."
  }
}
```

## 注意事项

1. **任务去重**：使用 `auction_id` 作为唯一标识，相同 `auction_id` 的任务会覆盖之前的任务
2. **时间精度**：任务会在 `end_time` 时触发，如果时间已过，会立即执行
3. **任务取消**：Asynq 不支持直接取消延迟任务，取消逻辑在任务执行时检查
4. **系统重启**：系统重启后会自动恢复所有未执行的拍卖任务
5. **错误重试**：任务执行失败会自动重试（Asynq 默认重试机制）

## 扩展开发

### 添加自定义处理逻辑

在 `processAuctionEnd` 方法中添加：

```go
func (s *AuctionTaskScheduler) processAuctionEnd(auction *models.Auction) error {
    // 1. 更新拍卖状态
    auction.Status = "ended"
    if err := database.DB.Save(auction).Error; err != nil {
        return err
    }

    // 2. 添加自定义逻辑
    // 例如：发送通知
    // notificationService.SendAuctionEndNotification(auction)
    
    // 例如：处理最高出价者
    // if auction.HighestBidder != "" {
    //     processWinner(auction)
    // }
    
    return nil
}
```

### 添加任务类型

如果需要添加其他类型的任务，可以扩展 `AuctionTaskScheduler`：

```go
// 添加新的任务类型
func (s *AuctionTaskScheduler) ScheduleCustomTask(taskType string, taskID string, executeTime time.Time, params map[string]interface{}) error {
    // 实现自定义任务调度逻辑
}
```

## 故障排查

### 任务未执行

1. 检查 Redis 连接是否正常
2. 检查任务是否已入队（查看 Redis）
3. 检查任务处理器是否正常启动
4. 查看日志中的错误信息

### 任务重复执行

1. 检查是否有多个实例运行
2. 检查 Redis 配置是否正确
3. 检查任务去重逻辑

### 任务执行时间不准确

1. 检查系统时间是否准确
2. 检查时区配置
3. 检查 `end_time` 是否正确设置

## 监控和日志

任务调度器会记录以下日志：

- 任务调度：`Auction end task scheduled: auctionID=xxx, endTime=xxx`
- 任务执行：`Processing auction end task: auctionID=xxx`
- 任务完成：`Auction end task completed: auctionID=xxx`
- 任务失败：`Failed to process auction end: auctionID=xxx, error=xxx`
- 任务恢复：`Restored X/Y auction end tasks`

## 总结

这个框架提供了完整的拍卖任务调度功能，包括：

- ✅ 自动调度和恢复
- ✅ 动态添加和删除
- ✅ 精确时间触发
- ✅ 任务去重
- ✅ 错误处理和重试
- ✅ API 接口支持

可以直接使用，也可以根据需求扩展。

