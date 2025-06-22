package mocks

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/mock"
)

type MockWebSocketManager struct {
	mock.Mock
}

func (m *MockWebSocketManager) AddSocketConnection(id string, conn *websocket.Conn) error {
	args := m.Called(id, conn)
	return args.Error(0)
}

func (m *MockWebSocketManager) FindSocketConnection(id string) (*websocket.Conn, error) {
	args := m.Called(id)
	return args.Get(0).(*websocket.Conn), args.Error(1)
}

func (m *MockWebSocketManager) RemoveSocketConnection(id string) {
	m.Called(id)
}

func (m *MockWebSocketManager) BulkFindSocketConnection(ids []string) map[string]*websocket.Conn {
	args := m.Called(ids)
	return args.Get(0).(map[string]*websocket.Conn)
}
