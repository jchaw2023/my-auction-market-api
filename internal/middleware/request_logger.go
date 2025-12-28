package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/logger"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		
		c.Next()
		
		latency := time.Since(start)
		status := c.Writer.Status()
		
		logger.Info("%s %s -> %d (%s)", method, path, status, latency)
	}
}

