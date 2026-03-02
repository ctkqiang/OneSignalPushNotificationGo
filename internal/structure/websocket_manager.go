package structure

import (
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketManager 管理所有 WebSocket 连接
type WebSocketManager struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan []byte
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	Mutex      sync.Mutex
}
