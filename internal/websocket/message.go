package websocket

import "time"

// MessageType 消息类型
type MessageType string

const (
	// ========== 业务事件消息（服务端 -> 客户端）==========
	// MessageTypeAuctionCreated 拍卖创建事件
	// 当新的拍卖被创建时发送，广播给所有客户端
	MessageTypeAuctionCreated MessageType = "auction_created"

	// MessageTypeAuctionBidPlaced 出价事件
	// 当有新的出价时发送，仅推送给订阅了该拍卖房间的客户端
	MessageTypeAuctionBidPlaced MessageType = "auction_bid_placed"

	// MessageTypeAuctionEnded 拍卖结束事件
	// 当拍卖结束时发送，广播给所有客户端
	MessageTypeAuctionEnded MessageType = "auction_ended"

	// MessageTypeAuctionCancelled 拍卖取消事件
	// 当拍卖被取消时发送，广播给所有客户端
	MessageTypeAuctionCancelled MessageType = "auction_cancelled"

	// MessageTypeAuctionForceEnded 强制结束拍卖事件
	// 当拍卖被强制结束时发送，广播给所有客户端
	MessageTypeAuctionForceEnded MessageType = "auction_force_ended"

	// MessageTypeNFTApproved NFT授权事件
	// 当NFT被授权给拍卖合约时发送，广播给所有客户端
	MessageTypeNFTApproved MessageType = "nft_approved"

	// ========== 系统消息 ==========
	// MessageTypeError 错误消息
	// 当发生错误时发送给客户端
	MessageTypeError MessageType = "error"

	// MessageTypePing 心跳请求
	// 客户端或服务端发送的心跳请求
	MessageTypePing MessageType = "ping"

	// MessageTypePong 心跳响应
	// 对Ping消息的响应
	MessageTypePong MessageType = "pong"

	// ========== 房间订阅消息（客户端 <-> 服务端）==========
	// MessageTypeSubscribe 订阅房间请求
	// 客户端发送，请求订阅特定房间（如：auction:{auctionID}）
	MessageTypeSubscribe MessageType = "subscribe"

	// MessageTypeUnsubscribe 取消订阅房间请求
	// 客户端发送，请求取消订阅特定房间
	MessageTypeUnsubscribe MessageType = "unsubscribe"

	// MessageTypeSubscribeSuccess 订阅成功响应
	// 服务端发送，确认客户端已成功订阅房间
	MessageTypeSubscribeSuccess MessageType = "subscribe_success"

	// MessageTypeUnsubscribeSuccess 取消订阅成功响应
	// 服务端发送，确认客户端已成功取消订阅房间
	MessageTypeUnsubscribeSuccess MessageType = "unsubscribe_success"
)

// Message WebSocket 消息结构
type Message struct {
	Type      MessageType `json:"type"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
}

// NewMessage 创建新消息
func NewMessage(msgType MessageType, data interface{}) *Message {
	return &Message{
		Type:      msgType,
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
}

// NewErrorMessage 创建错误消息
func NewErrorMessage(err error) *Message {
	return &Message{
		Type:      MessageTypeError,
		Timestamp: time.Now().Unix(),
		Error:     err.Error(),
	}
}

// SubscribeMessage 订阅消息结构
type SubscribeMessage struct {
	RoomID string `json:"room_id"` // 房间ID，例如：auction:{auctionID}
}
