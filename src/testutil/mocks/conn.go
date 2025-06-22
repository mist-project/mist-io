package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockWebSocketConnection struct {
	mock.Mock
}

func (wsc *MockWebSocketConnection) WriteMessage(messageType int, data []byte) error {
	args := wsc.Called(messageType, data)
	return args.Error(0)
}
