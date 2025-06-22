package ws

import (
	"fmt"
	"log/slog"
	"mist-io/src/internal/faults"
	"mist-io/src/internal/logging/logger"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager interface {
	AddSocketConnection(id string, conn *websocket.Conn) error
	FindSocketConnection(id string) (*websocket.Conn, error)
	RemoveSocketConnection(id string)
	BulkFindSocketConnection(id []string) map[string]*websocket.Conn
}

type WSManager struct {
	connections map[string]*websocket.Conn
	mu          sync.Mutex
}

func NewWSManager() *WSManager {
	return &WSManager{
		connections: make(map[string]*websocket.Conn),
		mu:          sync.Mutex{},
	}
}

func (wsm *WSManager) SetConnections(connections map[string]*websocket.Conn) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	wsm.connections = connections
}

func (wsm *WSManager) AddSocketConnection(id string, conn *websocket.Conn) error {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	val, ok := wsm.connections[id]

	if ok {
		// TODO: for now we are only allowing one connection per user. if the connection isn't the same, first previous
		// one will be disconnected
		if val == conn {
			return nil // if the conn is the same, don't do anything
		}

		logger.Debug("Multi-socket session detected. First socket wasc losed.", "socketId", id)
		val.Close()
	}

	wsm.connections[id] = conn

	return nil
}

func (wsm *WSManager) FindSocketConnection(id string) (*websocket.Conn, error) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	val, ok := wsm.connections[id]

	if !ok {
		return nil, faults.SocketNotFoundError(fmt.Sprintf("socket with id (%s) not found ", id), slog.LevelDebug)
	}

	return val, nil
}

func (wsm *WSManager) BulkFindSocketConnection(id []string) map[string]*websocket.Conn {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	conns := make(map[string]*websocket.Conn)
	for _, v := range id {
		if conn, ok := wsm.connections[v]; ok {
			conns[v] = conn
		}
	}

	return conns
}

func (wsm *WSManager) RemoveSocketConnection(id string) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	delete(wsm.connections, id)
}
