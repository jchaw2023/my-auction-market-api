package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	appJWT "my-auction-market-api/internal/jwt"
	"my-auction-market-api/internal/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 在生产环境中，应该检查 Origin 头
		// 这里允许所有来源（仅用于开发）
		return true
	},
}

// ServeWS 处理 WebSocket 连接请求
func ServeWS(hub *Hub, c *gin.Context) {
	// 可选：验证 JWT token（如果需要认证）
	var userID uint = 0
	if token := c.Query("token"); token != "" {
		claims, err := appJWT.ParseToken(token)
		if err == nil {
			userID = uint(claims.UserID)
		}
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("websocket upgrade failed: %v", err)
		return
	}

	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		userID: userID,
	}

	client.hub.register <- client

	// 启动客户端的读写 goroutines
	go client.writePump()
	go client.readPump()
}

