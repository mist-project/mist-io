package ws

import (
	"fmt"
	"log/slog"
	"mist-io/src/internal/faults"
	"mist-io/src/internal/logging/logger"
	"sync"
)

type WSConnInterface interface {
	Close() error
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
}

type WebSocketManager interface {
	AddSocketConnection(id string, conn WSConnInterface) error
	FindSocketConnection(id string) (WSConnInterface, error)
	RemoveSocketConnection(id string)
}

type WSManager struct {
	connections map[string]WSConnInterface
	mu          sync.Mutex
}

func NewWSManager() *WSManager {
	return &WSManager{
		connections: make(map[string]WSConnInterface),
		mu:          sync.Mutex{},
	}
}

func (wsm *WSManager) SetConnections(connections map[string]WSConnInterface) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	wsm.connections = connections
}

func (wsm *WSManager) AddSocketConnection(id string, conn WSConnInterface) error {
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

func (wsm *WSManager) FindSocketConnection(id string) (WSConnInterface, error) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()

	val, ok := wsm.connections[id]

	if !ok {
		return nil, faults.SocketNotFoundError(fmt.Sprintf("socket with id (%s) not found ", id), slog.LevelDebug)
	}

	return val, nil
}

func (wsm *WSManager) RemoveSocketConnection(id string) {
	wsm.mu.Lock()
	defer wsm.mu.Unlock()
	delete(wsm.connections, id)
}
