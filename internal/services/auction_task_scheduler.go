package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"my-auction-market-api/internal/config"
	"my-auction-market-api/internal/contracts/my_auction"
	"my-auction-market-api/internal/database"
	"my-auction-market-api/internal/ethereum"
	"my-auction-market-api/internal/logger"
	"my-auction-market-api/internal/models"
	"my-auction-market-api/internal/redisdb"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

// AuctionTaskScheduler 拍卖任务调度器
// 用于处理基于 end_time 的一次性任务
type AuctionTaskScheduler struct {
	client          *asynq.Client
	server          *asynq.Server
	mux             *asynq.ServeMux
	inspector       *asynq.Inspector // Inspector 用于删除任务
	redisConn       asynq.RedisClientOpt
	redisClient     *redisdb.Client          // Redis 客户端，用于存储取消标记
	ethClient       *ethereum.Client         // 以太坊客户端
	auctionContract *my_auction.MyXAuctionV2 // 拍卖合约实例
	cfg             *config.Config
	mu              sync.RWMutex
}

var (
	auctionTaskSchedulerInstance *AuctionTaskScheduler
	auctionTaskSchedulerOnce     sync.Once
)

// GetAuctionTaskScheduler 获取拍卖任务调度器单例
func GetAuctionTaskScheduler(cfg *config.Config) *AuctionTaskScheduler {
	auctionTaskSchedulerOnce.Do(func() {
		redisConn := asynq.RedisClientOpt{
			Addr:     cfg.Redis.Addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
		}

		client := asynq.NewClient(redisConn)
		inspector := asynq.NewInspector(redisConn)

		server := asynq.NewServer(redisConn, asynq.Config{
			Concurrency: 10, // 并发处理任务数
			Queues: map[string]int{
				"auctions": 10, // 拍卖任务队列
			},
		})

		mux := asynq.NewServeMux()

		// 初始化 Redis 客户端（用于存储取消标记）
		redisClient, err := redisdb.NewClient(cfg.Redis.ToRedisConfig())
		if err != nil {
			logger.Error("Failed to initialize Redis client for task scheduler: %v", err)
			// 继续执行，取消功能可能不可用
		}

		// 初始化以太坊客户端和合约（用于调用合约方法）
		var ethClient *ethereum.Client
		var auctionContract *my_auction.MyXAuctionV2
		if cfg.Ethereum.AuctionContractAddress != "" && cfg.Ethereum.RPCURL != "" {
			ethClient, err = ethereum.NewClient(cfg.Ethereum)
			if err != nil {
				logger.Error("Failed to initialize Ethereum client for task scheduler: %v", err)
				logger.Warn("Contract calls will not be available")
			} else {
				// 创建合约实例
				contractAddress := common.HexToAddress(cfg.Ethereum.AuctionContractAddress)
				auctionContract, err = my_auction.NewMyXAuctionV2(contractAddress, ethClient.GetClient())
				if err != nil {
					logger.Error("Failed to create auction contract instance: %v", err)
					logger.Warn("Contract calls will not be available")
					auctionContract = nil
				} else {
					logger.Info("Auction contract initialized for task scheduler: %s", cfg.Ethereum.AuctionContractAddress)
				}
			}
		} else {
			logger.Warn("Ethereum configuration not available, contract calls will not be available")
		}

		auctionTaskSchedulerInstance = &AuctionTaskScheduler{
			client:          client,
			server:          server,
			mux:             mux,
			inspector:       inspector,
			redisConn:       redisConn,
			redisClient:     redisClient,
			ethClient:       ethClient,
			auctionContract: auctionContract,
			cfg:             cfg,
		}

		// 注册任务处理器
		auctionTaskSchedulerInstance.registerHandlers()
	})

	return auctionTaskSchedulerInstance
}

// registerHandlers 注册任务处理器
func (s *AuctionTaskScheduler) registerHandlers() {
	// 注册拍卖结束任务处理器
	// 注意：任务类型必须与 NewTask 的第一个参数一致
	s.mux.HandleFunc("auction-end", s.handleAuctionEndTask)
}

// AuctionTaskPayload 拍卖任务负载
type AuctionTaskPayload struct {
	TaskID   string `json:"task_id"`   // auction_id (不重复)
	TaskName string `json:"task_name"` // "auctions" (可能重复)
	UserID   uint64 `json:"user_id"`   // 用户ID
	NFTID    string `json:"nft_id"`    // NFT ID
}

// getCustomTaskID 生成自定义 TaskID（统一格式）
func getCustomTaskID(auctionID string) string {
	return fmt.Sprintf("auction-end:%s", auctionID)
}

// ScheduleAuctionEndTask 调度拍卖结束任务
// auctionID: 拍卖ID (作为 task_id，不重复)
// endTime: 结束时间
// userID: 用户ID
// nftID: NFT ID
func (s *AuctionTaskScheduler) ScheduleAuctionEndTask(auction *models.Auction) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	auctionID := auction.AuctionID
	endTime := auction.EndTime
	userID := auction.UserID
	nftID := auction.NFTID
	if endTime == nil {
		return fmt.Errorf("endTime is nil, cannot schedule auction end task: auctionID=%s", auctionID)
	}
	// 如果任务已存在，先删除旧任务（避免 TaskID 冲突）
	// 使用相同的 customTaskID 会导致 Asynq 返回 ErrTaskIDConflict 错误
	customTaskID := getCustomTaskID(auctionID)
	if s.inspector != nil {
		// 尝试删除可能存在的旧任务（静默失败，因为任务可能不存在）
		if err := s.inspector.DeleteTask("auctions", customTaskID); err == nil {
			logger.Info("Deleted existing task before rescheduling: auctionID=%s, taskID=%s", auctionID, customTaskID)
		}
	}

	// 清除取消标记（如果存在）
	if s.redisClient != nil {
		cancelKey := s.getCancelKey(auctionID)
		s.redisClient.Del(context.Background(), cancelKey)
	}

	// 构建任务负载
	payload := AuctionTaskPayload{
		TaskID:   auctionID,
		TaskName: "auction-end", // 任务名称
		UserID:   userID,
		NFTID:    nftID,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// 创建任务，使用自定义 TaskID（基于 auctionID）
	// 使用自定义 TaskID，而不是依赖自动生成的 UUID
	// 注意：customTaskID 已在上面生成
	task := asynq.NewTask("auction-end", payloadBytes, asynq.TaskID(customTaskID))

	// 计算延迟时间
	now := time.Now()
	var info *asynq.TaskInfo

	// 指定队列名称（与服务器配置中的队列名称一致）
	queueName := "auctions"

	if endTime.Before(now) {
		// 如果结束时间已过，立即执行
		logger.Info("Auction %s end time has passed, executing immediately", auctionID)
		info, err = s.client.Enqueue(task, asynq.Queue(queueName))
	} else {
		// 计算延迟时间
		delay := endTime.Sub(now)
		logger.Info("Scheduling auction end task: auctionID=%s, delay=%v, endTime=%v, queue=%s", auctionID, delay, endTime, queueName)

		// 使用 ProcessIn 选项延迟执行，并指定队列
		info, err = s.client.Enqueue(task, asynq.Queue(queueName), asynq.ProcessIn(delay))
	}

	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	// 任务已调度，使用自定义 TaskID（不需要保存映射，可以直接生成）
	logger.Info("Auction end task scheduled: auctionID=%s, customTaskID=%s, asynqTaskID=%s, endTime=%v",
		auctionID, customTaskID, func() string {
			if info != nil {
				return info.ID
			}
			return "N/A"
		}(), endTime)

	return nil
}

// CancelAuctionEndTask 取消拍卖结束任务
// 通过设置 Redis 取消标记来实现任务删除
func (s *AuctionTaskScheduler) CancelAuctionEndTask(auctionID string) error {
	return s.DeleteTask(auctionID)
}

// DeleteTask 删除任务（推荐使用此方法）
// 直接使用 auctionID 生成自定义 TaskID 来删除任务，无需依赖保存的映射
// 1. 尝试从 Asynq 队列中删除任务（使用自定义 TaskID）
// 2. 设置取消标记作为备用方案
func (s *AuctionTaskScheduler) DeleteTask(auctionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ctx := context.Background()
	deleted := false

	// 直接使用 auctionID 生成自定义 TaskID（与 ScheduleAuctionEndTask 中使用的格式一致）
	customTaskID := getCustomTaskID(auctionID)

	// 方法1: 尝试从 Asynq 队列中删除任务（使用自定义 TaskID）
	// 由于任务明确指定了 "auctions" 队列，只需在此队列中查找
	if s.inspector != nil {
		// 优先使用自定义 TaskID 直接删除（task.ID 就是 customTaskID，无需遍历）
		if err := s.inspector.DeleteTask("auctions", customTaskID); err == nil {
			logger.Info("Task deleted from auctions queue using custom TaskID: auctionID=%s, taskID=%s", auctionID, customTaskID)
			deleted = true
		} else {
			// 如果直接删除失败，记录警告（正常情况下不应该发生）
			logger.Warn("Failed to delete task directly by customTaskID: auctionID=%s, taskID=%s, error=%v", auctionID, customTaskID, err)

			// 备用方案：通过 payload 匹配查找并删除（仅在异常情况下使用）
			// 限制检查数量，避免在任务很多时性能问题（最多检查前 100 个任务）
			scheduledTasks, err := s.inspector.ListScheduledTasks("auctions")
			if err == nil {
				maxCheck := 100 // 最多检查前 100 个任务
				checked := 0
				for _, task := range scheduledTasks {
					if checked >= maxCheck {
						logger.Warn("Reached max check limit (%d) while searching for task by payload: auctionID=%s", maxCheck, auctionID)
						break
					}
					checked++

					// 检查任务的 payload 是否匹配
					var payload AuctionTaskPayload
					if json.Unmarshal(task.Payload, &payload) == nil && payload.TaskID == auctionID {
						if err := s.inspector.DeleteTask("auctions", task.ID); err == nil {
							logger.Info("Task deleted by payload match: auctionID=%s, taskID=%s, checked=%d tasks", auctionID, task.ID, checked)
							deleted = true
							break
						}
					}
				}
				if !deleted && checked >= maxCheck {
					logger.Warn("Task not found in first %d scheduled tasks: auctionID=%s", maxCheck, auctionID)
				}
			}
		}
	}

	// 方法2: 设置取消标记作为备用方案（即使任务已从队列删除，也设置标记以防万一）
	if s.redisClient != nil {
		cancelKey := s.getCancelKey(auctionID)
		if err := s.redisClient.Set(ctx, cancelKey, "1", 24*time.Hour).Err(); err != nil {
			logger.Error("Failed to set cancel marker: %v", err)
		} else {
			logger.Info("Cancel marker set: auctionID=%s", auctionID)
		}
	}

	if deleted {
		logger.Info("Auction end task deleted successfully: auctionID=%s", auctionID)
	} else {
		logger.Info("Auction end task cancellation marker set (task may not be in queue): auctionID=%s", auctionID)
	}

	return nil
}

// getCancelKey 获取取消标记的 Redis key
func (s *AuctionTaskScheduler) getCancelKey(auctionID string) string {
	return fmt.Sprintf("auction:task:cancel:%s", auctionID)
}

// isTaskCancelled 检查任务是否已被取消
func (s *AuctionTaskScheduler) isTaskCancelled(auctionID string) bool {
	if s.redisClient == nil {
		return false
	}

	cancelKey := s.getCancelKey(auctionID)
	ctx := context.Background()

	exists, err := s.redisClient.Exists(ctx, cancelKey).Result()
	if err != nil {
		logger.Error("Failed to check cancel marker: %v", err)
		return false
	}

	return exists > 0
}

// handleAuctionEndTask 处理拍卖结束任务
func (s *AuctionTaskScheduler) handleAuctionEndTask(ctx context.Context, t *asynq.Task) error {
	var payload AuctionTaskPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	logger.Info("Processing auction end task: auctionID=%s, userID=%d, nftID=%s",
		payload.TaskID, payload.UserID, payload.NFTID)

	// 首先检查任务是否已被取消
	if s.isTaskCancelled(payload.TaskID) {
		logger.Info("Task has been cancelled, skipping: auctionID=%s", payload.TaskID)
		// 清除取消标记
		if s.redisClient != nil {
			cancelKey := s.getCancelKey(payload.TaskID)
			s.redisClient.Del(ctx, cancelKey)
		}
		return nil // 不返回错误，避免重试
	}

	// 检查拍卖是否仍然有效（可能已被取消）
	var auction models.Auction
	if err := database.DB.Where("auction_id = ?", payload.TaskID).First(&auction).Error; err != nil {
		logger.Warn("Auction not found or already processed: auctionID=%s, error=%v", payload.TaskID, err)
		return nil // 不返回错误，避免重试
	}

	// 检查拍卖状态
	if auction.Status != "active" {
		logger.Info("Auction is not active, skipping: auctionID=%s, status=%s", payload.TaskID, auction.Status)
		return nil
	}

	// 检查结束时间是否已到（防止提前执行）
	now := time.Now()
	if auction.EndTime != nil && auction.EndTime.After(now) {
		logger.Warn("Auction end time not reached yet, rescheduling: auctionID=%s, endTime=%v, now=%v", payload.TaskID, auction.EndTime, now)

		// 重新调度任务
		delay := auction.EndTime.Sub(now)
		task := asynq.NewTask("auction-end", t.Payload())
		_, err := s.client.Enqueue(task, asynq.Queue("auctions"), asynq.ProcessIn(delay))
		if err != nil {
			logger.Error("Failed to reschedule task: %v", err)
		}
		return nil
	}

	// 执行拍卖结束逻辑
	if err := s.processAuctionEnd(&auction); err != nil {
		logger.Error("Failed to process auction end: auctionID=%s, error=%v", payload.TaskID, err)
		return err // 返回错误会触发重试
	}

	logger.Info("Auction end task completed: auctionID=%s", payload.TaskID)
	return nil
}

// processAuctionEnd 处理拍卖结束逻辑
func (s *AuctionTaskScheduler) processAuctionEnd(auction *models.Auction) error {
	// 更新拍卖状态为 ended（只更新 Status 字段，避免更新所有字段）
	// 更新 online_lock 为 0，表示拍卖结束，可以被其他用户竞拍
	return database.DB.Transaction(func(tx *gorm.DB) error {
		onlineLock := fmt.Sprintf("%s:%s", auction.NFTID, auction.AuctionID) //相当于给现在的拍卖与NFT直接解锁了,后续其他购买的用户也可以进行竞拍
		if err := tx.Model(auction).
			Updates(map[string]interface{}{
				"status":      "ended",
				"online_lock": onlineLock}).Error; err != nil {
			return fmt.Errorf("failed to update auction status: %w", err)
		}

		// 如果有最高出价者，调用合约方法将 NFT 转移给最高出价者
		// highestBidder := auction.HighestBidder
		if s.auctionContract != nil && auction.ContractAuctionID > 0 {
			// 调用合约的 endAuctionAndClaimNFT 方法
			// 注意：这需要平台私钥来签名交易，需要在配置中添加平台私钥
			// 目前先记录日志，实际调用需要配置私钥

			// TODO: 需要配置平台私钥来签名交易
			// 示例代码（需要配置私钥后启用）：
			ctx := context.Background()
			auth, err := s.ethClient.GetAuth(ctx, s.cfg.Ethereum.PlatformPrivateKey) // 需要实现 getAuth 方法，使用平台私钥
			if err != nil {
				logger.Error("Failed to get transaction auth: %v", err)
				return fmt.Errorf("failed to get transaction auth: %w", err)
			}

			contractAuctionID := big.NewInt(int64(auction.ContractAuctionID))
			auctionOnChain, err := s.auctionContract.GetAuction(&bind.CallOpts{Context: ctx}, contractAuctionID)
			if err != nil {
				logger.Error("Failed to get auction on chain: %v", err)
				return fmt.Errorf("failed to get auction on chain: %w", err)
			}
			highestBidder := strings.ToLower(auctionOnChain.HighestBidder.Hex())
			logger.Info("Auction ended with highest bidder, should transfer NFT: auctionID=%s, contractAuctionID=%d, highestBidder=%s", auction.AuctionID, auction.ContractAuctionID, highestBidder)
			tx, err := s.auctionContract.EndAuctionAndClaimNFT(auth, contractAuctionID)
			if err != nil {
				logger.Error("Failed to call EndAuctionAndClaimNFT: %v", err)
				return fmt.Errorf("failed to call EndAuctionAndClaimNFT: %w", err)
			}
			logger.Info("NFT transfer transaction sent: txHash=%s, auctionID=%s", tx.Hash().Hex(), auction.AuctionID)

			if auctionOnChain.HighestBidder != common.BigToAddress(big.NewInt(0)) {
				auctionOnChainHighestBidder := strings.ToLower(auctionOnChain.HighestBidder.Hex())
				if auctionOnChainHighestBidder != "" && auctionOnChainHighestBidder != strings.ToLower(highestBidder) {
					logger.Error("Highest bidder is not the winner: auctionID=%s, highestBidder=%s, winner-on-chain=%s", auction.AuctionID, highestBidder, auctionOnChainHighestBidder)
					return fmt.Errorf("highest bidder is not the winner: auctionID=%s, highestBidder=%s, winner-on-chain=%s", auction.AuctionID, highestBidder, auctionOnChainHighestBidder)
				}
				//将nft_ownerships里面状态改为已出售
				if err := database.DB.Model(&models.NFTOwnership{}).
					Where("nft_id = ? and user_id = ?", auction.NFTID, auction.UserID).
					Updates(
						map[string]interface{}{
							"status":        models.NFTOwnershipStatusSold,
							"owner_address": highestBidder,
							"approved":      0,
							"updated_at":    time.Now(),
						}).Error; err != nil {
					logger.Error("Failed to update NFT ownership status: %v", err)
					return fmt.Errorf("failed to update NFT ownership status: %w", err)
				}
			} else { //没有最高出价者，则将nft_ownerships里面状态改为卖出者
				if err := database.DB.Model(&models.NFTOwnership{}).
					Where("nft_id = ? and user_id = ?", auction.NFTID, auction.UserID).
					Updates(
						map[string]interface{}{
							"status":        models.NFTOwnershipStatusHolding,
							"owner_address": strings.ToLower(auctionOnChain.Seller.Hex()),
							"approved":      0,
							"updated_at":    time.Now(),
						}).Error; err != nil {
					logger.Error("Failed to update NFT ownership status: %v", err)
					return fmt.Errorf("failed to update NFT ownership status: %w", err)
				}
			}

		} else {
			return fmt.Errorf("auction contract not available, cannot transfer NFT: auctionID=%s", auction.AuctionID)
		}
		logger.Info("Auction ended successfully: auctionID=%s", auction.AuctionID)
		return nil
	})
}

// RestoreAuctionTasks 恢复未执行的拍卖任务（系统启动时调用）
func (s *AuctionTaskScheduler) RestoreAuctionTasks() error {
	// 查找所有活跃的拍卖，其结束时间在未来
	var auctions []*models.Auction
	now := time.Now()

	if err := database.DB.Where("status IN ? AND end_timestamp > ? and online=1 ",
		[]string{"active"}, now.Unix()).Find(&auctions).Error; err != nil {
		return fmt.Errorf("failed to load auctions: %w", err)
	}

	logger.Info("Restoring %d auction end tasks", len(auctions))

	successCount := 0
	for _, auction := range auctions {
		// 重新调度任务
		if err := s.ScheduleAuctionEndTask(auction); err != nil {
			logger.Error("Failed to restore task for auction %s: %v", auction.AuctionID, err)
			continue
		}

		successCount++
	}

	logger.Info("Restored %d/%d auction end tasks", successCount, len(auctions))
	return nil
}

// Start 启动任务处理器
func (s *AuctionTaskScheduler) Start(ctx context.Context) error {
	// 恢复未执行的任务
	if err := s.RestoreAuctionTasks(); err != nil {
		logger.Error("Failed to restore auction tasks: %v", err)
	}

	// 启动服务器
	logger.Info("Starting auction task scheduler server...")
	if err := s.server.Run(s.mux); err != nil {
		return fmt.Errorf("failed to start task server: %w", err)
	}

	return nil
}

// StartAsync 异步启动任务处理器
func (s *AuctionTaskScheduler) StartAsync(ctx context.Context) {
	go func() {
		if err := s.Start(ctx); err != nil {
			logger.Error("Auction task scheduler error: %v", err)
		}
	}()
}

// Shutdown 关闭任务调度器
func (s *AuctionTaskScheduler) Shutdown() {
	if s.client != nil {
		s.client.Close()
	}
	if s.server != nil {
		s.server.Shutdown()
	}
	if s.inspector != nil {
		s.inspector.Close()
	}
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	logger.Info("Auction task scheduler shut down")
}

// GetTaskStatus 获取任务状态（用于调试）
// 通过查询 Asynq 队列来检查任务是否存在
func (s *AuctionTaskScheduler) GetTaskStatus(auctionID string) map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	status := map[string]interface{}{
		"auction_id": auctionID,
		"scheduled":  false,
		"cancelled":  false,
	}

	// 生成自定义 TaskID
	customTaskID := getCustomTaskID(auctionID)
	status["task_id"] = customTaskID

	// 检查任务是否在队列中（通过查询 Asynq）
	if s.inspector != nil {
		// 在 scheduled 队列中查找
		scheduledTasks, err := s.inspector.ListScheduledTasks("auctions")
		if err == nil {
			for _, task := range scheduledTasks {
				if task.ID == customTaskID {
					status["scheduled"] = true
					status["task_info"] = map[string]interface{}{
						"id":      task.ID,
						"type":    task.Type,
						"payload": string(task.Payload),
						"next":    task.NextProcessAt,
					}
					break
				}
			}
		}

		// 如果不在 scheduled 队列，检查 pending 队列
		if !status["scheduled"].(bool) {
			pendingTasks, err := s.inspector.ListPendingTasks("auctions")
			if err == nil {
				for _, task := range pendingTasks {
					if task.ID == customTaskID {
						status["scheduled"] = true
						status["task_info"] = map[string]interface{}{
							"id":      task.ID,
							"type":    task.Type,
							"payload": string(task.Payload),
						}
						break
					}
				}
			}
		}
	}

	// 检查是否已取消
	if s.isTaskCancelled(auctionID) {
		status["cancelled"] = true
	}

	return status
}
