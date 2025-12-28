package websocket

import (
	"encoding/json"
	"sync"
	"time"

	"my-auction-market-api/internal/logger"

	"github.com/gorilla/websocket"
)

const (
	// 允许客户端写入的最大时间
	writeWait = 10 * time.Second

	// 允许客户端读取下一个 pong 消息的时间
	pongWait = 60 * time.Second

	// 在此时间内发送 ping，必须小于 pongWait
	pingPeriod = (pongWait * 9) / 10

	// 允许的最大消息大小
	maxMessageSize = 512 * 1024
)

// Subscription 表示客户端订阅/取消订阅房间的请求
type Subscription struct {
	client *Client
	roomID string
}

// Hub 维护所有活跃的客户端连接
type Hub struct {
	// 注册的客户端
	clients map[*Client]bool

	// 房间映射：roomID -> 订阅该房间的客户端集合
	rooms map[string]map[*Client]bool

	// 从客户端接收消息的通道
	broadcast chan []byte

	// 注册新客户端的通道
	register chan *Client

	// 取消注册客户端的通道
	unregister chan *Client

	// 客户端订阅房间的通道
	subscribe chan *Subscription

	// 客户端取消订阅房间的通道
	unsubscribe chan *Subscription

	// 保护 clients 和 rooms map 的互斥锁
	mu sync.RWMutex
}

// NewHub 创建新的 Hub
func NewHub() *Hub {
	return &Hub{
		clients:     make(map[*Client]bool),
		rooms:       make(map[string]map[*Client]bool),
		broadcast:   make(chan []byte, 256),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		subscribe:   make(chan *Subscription),
		unsubscribe: make(chan *Subscription),
	}
}

// Run 启动 Hub，处理客户端注册、注销和消息广播
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			logger.Info("websocket client connected, total clients: %d", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				// 从所有房间中移除该客户端
				for roomID, clients := range h.rooms {
					delete(clients, client)
					if len(clients) == 0 {
						delete(h.rooms, roomID)
					}
				}
			}
			h.mu.Unlock()
			logger.Info("websocket client disconnected, total clients: %d", len(h.clients))

		case sub := <-h.subscribe:
			h.mu.Lock()
			if h.rooms[sub.roomID] == nil {
				h.rooms[sub.roomID] = make(map[*Client]bool)
			}
			h.rooms[sub.roomID][sub.client] = true
			h.mu.Unlock()
			logger.Info("client subscribed to room: %s, total clients in room: %d", sub.roomID, len(h.rooms[sub.roomID]))

		case sub := <-h.unsubscribe:
			h.mu.Lock()
			if clients, ok := h.rooms[sub.roomID]; ok {
				delete(clients, sub.client)
				if len(clients) == 0 {
					delete(h.rooms, sub.roomID)
				}
			}
			h.mu.Unlock()
			logger.Info("client unsubscribed from room: %s", sub.roomID)

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// BroadcastMessage 广播消息给所有客户端
func (h *Hub) BroadcastMessage(message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	select {
	case h.broadcast <- data:
	default:
		logger.Warn("websocket broadcast channel is full, dropping message")
	}

	return nil
}

// GetClientCount 获取当前连接的客户端数量
func (h *Hub) GetClientCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

// BroadcastToRoom 向特定房间的所有客户端推送消息
func (h *Hub) BroadcastToRoom(roomID string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	h.mu.RLock()
	clients, exists := h.rooms[roomID]
	if !exists || len(clients) == 0 {
		h.mu.RUnlock()
		logger.Debug("no clients in room: %s, skipping message", roomID)
		return nil
	}

	// 复制客户端列表，避免在发送时持有锁
	clientList := make([]*Client, 0, len(clients))
	for client := range clients {
		clientList = append(clientList, client)
	}
	h.mu.RUnlock()

	// 发送消息给房间内的所有客户端
	sentCount := 0
	for _, client := range clientList {
		select {
		case client.send <- data:
			sentCount++
		default:
			// 客户端发送通道已满，关闭连接
			h.mu.Lock()
			close(client.send)
			delete(h.clients, client)
			// 从房间中移除
			if roomClients, ok := h.rooms[roomID]; ok {
				delete(roomClients, client)
			}
			h.mu.Unlock()
		}
	}

	logger.Debug("broadcasted message to room: %s, sent to %d/%d clients", roomID, sentCount, len(clientList))
	return nil
}

// SubscribeRoom 订阅房间
func (h *Hub) SubscribeRoom(client *Client, roomID string) {
	select {
	case h.subscribe <- &Subscription{client: client, roomID: roomID}:
	default:
		logger.Warn("subscribe channel is full, dropping subscription request")
	}
}

// UnsubscribeRoom 取消订阅房间
func (h *Hub) UnsubscribeRoom(client *Client, roomID string) {
	select {
	case h.unsubscribe <- &Subscription{client: client, roomID: roomID}:
	default:
		logger.Warn("unsubscribe channel is full, dropping unsubscription request")
	}
}

// Client 表示单个 WebSocket 连接
type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	userID uint // 可选的用户ID，用于定向推送
}

// readPump 从 WebSocket 连接读取消息
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		logger.Debug("websocket: received pong from client")
		return nil
	})

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("websocket read error: %v", err)
			}
			break
		}

		// 处理文本消息（可能是 PING 或其他消息）
		if messageType == websocket.TextMessage {
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				logger.Debug("websocket: failed to unmarshal message: %v, raw: %s", err, string(message))
				continue
			}

			// 处理 PING 消息
			if msg.Type == MessageTypePing {
				logger.Info("websocket: received ping from client (userID=%d), sending pong", c.userID)
				pongMsg := NewMessage(MessageTypePong, nil)
				pongData, err := json.Marshal(pongMsg)
				if err != nil {
					logger.Error("websocket: failed to marshal pong message: %v", err)
					continue
				}

				select {
				case c.send <- pongData:
					logger.Info("websocket: pong message queued for client (userID=%d)", c.userID)
				default:
					logger.Warn("websocket: client send channel is full, dropping pong message (userID=%d)", c.userID)
				}
				continue
			}

			// 处理订阅消息
			if msg.Type == MessageTypeSubscribe {
				var subMsg SubscribeMessage
				// 尝试从 Data 字段解析 room_id
				if dataMap, ok := msg.Data.(map[string]interface{}); ok {
					if roomID, ok := dataMap["room_id"].(string); ok {
						subMsg.RoomID = roomID
					}
				}

				if subMsg.RoomID != "" {
					c.hub.SubscribeRoom(c, subMsg.RoomID)
					// 发送订阅成功响应
					successMsg := NewMessage(MessageTypeSubscribeSuccess, map[string]interface{}{
						"room_id": subMsg.RoomID,
					})
					successData, _ := json.Marshal(successMsg)
					select {
					case c.send <- successData:
					default:
						logger.Warn("websocket: client send channel is full, dropping subscribe success message")
					}
					logger.Info("websocket: client (userID=%d) subscribed to room: %s", c.userID, subMsg.RoomID)
				} else {
					logger.Warn("websocket: invalid subscribe message, missing room_id")
				}
				continue
			}

			// 处理取消订阅消息
			if msg.Type == MessageTypeUnsubscribe {
				var subMsg SubscribeMessage
				// 尝试从 Data 字段解析 room_id
				if dataMap, ok := msg.Data.(map[string]interface{}); ok {
					if roomID, ok := dataMap["room_id"].(string); ok {
						subMsg.RoomID = roomID
					}
				}

				if subMsg.RoomID != "" {
					c.hub.UnsubscribeRoom(c, subMsg.RoomID)
					// 发送取消订阅成功响应
					successMsg := NewMessage(MessageTypeUnsubscribeSuccess, map[string]interface{}{
						"room_id": subMsg.RoomID,
					})
					successData, _ := json.Marshal(successMsg)
					select {
					case c.send <- successData:
					default:
						logger.Warn("websocket: client send channel is full, dropping unsubscribe success message")
					}
					logger.Info("websocket: client (userID=%d) unsubscribed from room: %s", c.userID, subMsg.RoomID)
				} else {
					logger.Warn("websocket: invalid unsubscribe message, missing room_id")
				}
				continue
			}

			// 其他消息类型可以在这里处理
			logger.Debug("websocket: received message type: %s from client (userID=%d)", msg.Type, c.userID)
		}
	}
}

// writePump 向 WebSocket 连接写入消息
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// 将队列中的其他消息也一起发送
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
