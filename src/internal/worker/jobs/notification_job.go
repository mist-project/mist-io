package jobs

import (
	"context"
	"log/slog"
	"mist-io/src/internal/faults"

	"github.com/gorilla/websocket"
)

type NotificationJob struct {
	content []byte
	wsConn  WebSocketConnection
	ctx     context.Context
}

func NewNotificationJob(content []byte, wsConn WebSocketConnection) *NotificationJob {
	return &NotificationJob{
		content: content,
		wsConn:  wsConn,
		ctx:     context.Background(), // You can pass a specific context if needed
	}
}

func (nj *NotificationJob) Execute(workerID int) error {
	err := nj.wsConn.WriteMessage(websocket.BinaryMessage, nj.content)

	if err != nil {
		// Log the error if writing to the WebSocket connection fails.
		return faults.SocketWriteError("error writing notification to WebSocket connection", slog.LevelWarn)
	}

	return nil
}

func (nj *NotificationJob) Ctx() context.Context {
	// Return a context for the job execution.
	// In a real implementation, you might want to use a more specific context.
	return nj.ctx
}
