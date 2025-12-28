package handlers

import (
	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/response"
	"my-auction-market-api/internal/services"
)

type AuctionTaskHandler struct {
	scheduler *services.AuctionTaskScheduler
}

func NewAuctionTaskHandler(scheduler *services.AuctionTaskScheduler) *AuctionTaskHandler {
	return &AuctionTaskHandler{
		scheduler: scheduler,
	}
}

// CancelAuctionTask 取消拍卖任务
// @Summary      Cancel auction end task
// @Description  Cancel a scheduled auction end task by setting a cancellation marker
// @Tags         auction-tasks
// @Accept       json
// @Produce      json
// @Param        auctionId  path      string  true  "Auction ID"
// @Success      200        {object}  response.Response
// @Failure      400        {object}  response.Response
// @Failure      500        {object}  response.Response
// @Router       /auction-tasks/{auctionId} [delete]
func (h *AuctionTaskHandler) CancelAuctionTask(c *gin.Context) {
	auctionID := c.Param("auctionId")
	if auctionID == "" {
		response.BadRequest(c, "auctionId is required")
		return
	}

	// 取消任务（设置 Redis 取消标记）
	if err := h.scheduler.CancelAuctionEndTask(auctionID); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, gin.H{
		"message":   "Task cancelled successfully",
		"auctionId": auctionID,
		"note":      "Task cancellation marker has been set. The task will be skipped when it executes.",
	})
}

// GetTaskStatus 获取任务状态（用于调试）
// @Summary      Get auction task status
// @Description  Get the status of a scheduled auction end task
// @Tags         auction-tasks
// @Accept       json
// @Produce      json
// @Param        auctionId  path      string  true  "Auction ID"
// @Success      200        {object}  response.Response
// @Failure      400        {object}  response.Response
// @Router       /auction-tasks/{auctionId} [get]
func (h *AuctionTaskHandler) GetTaskStatus(c *gin.Context) {
	auctionID := c.Param("auctionId")
	if auctionID == "" {
		response.BadRequest(c, "auctionId is required")
		return
	}

	status := h.scheduler.GetTaskStatus(auctionID)
	response.Success(c, status)
}

