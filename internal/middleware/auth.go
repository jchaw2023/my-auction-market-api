package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	appJWT "my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := appJWT.ParseToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		// Store user info in context
		c.Set("userId", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		c.Next()
	}
}

