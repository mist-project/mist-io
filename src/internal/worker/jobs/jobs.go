package jobs

import (
	"context"
)

// ------ STOP WORKER JOB -----
type StopWorkerJob struct {
	ctx context.Context
}

type WebSocketConnection interface {
	WriteMessage(messageType int, data []byte) error
}
