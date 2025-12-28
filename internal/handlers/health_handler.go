package handlers

import (
	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/response"
)

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if the API is running
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	response.Success(c, gin.H{"status": "ok"})
}

