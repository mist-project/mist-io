package subscriber_test

import (
	"context"
	"testing"

	"mist-io/src/internal/subscriber"
	"mist-io/src/internal/worker"
	"mist-io/src/internal/worker/jobs"
	"mist-io/src/protos/v1/appuser"
	"mist-io/src/protos/v1/event"
	"mist-io/src/testutil/mocks"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

type mockConn struct {
	websocket.Conn
	jobs.WebSocketConnection
}

func TestNewRedisListener(t *testing.T) {
	t.Run("constructs_listener_with_dependencies", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		mockWP := worker.NewWorkerPool(1, 10)
		mockWSM := new(mocks.MockWebSocketManager)

		opts := &subscriber.RedisListenerOptions{
			RedisClient: mockRedis,
			WorkerPool:  mockWP,
			WsManager:   mockWSM,
		}

		// ACT
		listener := subscriber.NewRedisListener(opts)

		// ASSERT
		assert.NotNil(t, listener)
		assert.Equal(t, mockRedis, listener.RedisClient())
		assert.Equal(t, mockWP, listener.WorkerPool())
		assert.Equal(t, mockWSM, listener.WsManager())
	})
}

func TestRedisListener_ProcessMessage(t *testing.T) {
	t.Run("logs_error_on_unmarshal_failure", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		mockWP := worker.NewWorkerPool(1, 10)
		mockWSM := new(mocks.MockWebSocketManager)

		opts := &subscriber.RedisListenerOptions{
			RedisClient: mockRedis,
			WorkerPool:  mockWP,
			WsManager:   mockWSM,
		}
		listener := subscriber.NewRedisListener(opts)
		msg := &redis.Message{Payload: "not a valid proto"}

		// ACT
		listener.ProcessMessage(context.Background(), msg)

		// ASSERT
		// no panic, error should be logged internally
	})

	t.Run("skips_when_no_appusers", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		mockWP := worker.NewWorkerPool(1, 10)
		mockWSM := new(mocks.MockWebSocketManager)

		opts := &subscriber.RedisListenerOptions{
			RedisClient: mockRedis,
			WorkerPool:  mockWP,
			WsManager:   mockWSM,
		}
		listener := subscriber.NewRedisListener(opts)

		ev := &event.Event{
			Meta: &event.Meta{Appusers: []*appuser.Appuser{}},
		}
		payload, _ := proto.Marshal(ev)

		msg := &redis.Message{Payload: string(payload)}

		// ACT
		listener.ProcessMessage(context.Background(), msg)

		// ASSERT
		// no panic, nothing sent
	})

	t.Run("marshals_event_and_adds_jobs", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		mockWP := worker.NewWorkerPool(1, 10)
		mockWSM := new(mocks.MockWebSocketManager)

		opts := &subscriber.RedisListenerOptions{
			RedisClient: mockRedis,
			WorkerPool:  mockWP,
			WsManager:   mockWSM,
		}
		listener := subscriber.NewRedisListener(opts)

		appuserID := "user-123"

		mockWSM.On("BulkFindSocketConnection", []string{appuserID}).Return(map[string]*websocket.Conn{
			appuserID: nil,
		})

		ev := &event.Event{
			Meta: &event.Meta{
				Appusers: []*appuser.Appuser{
					{Id: appuserID},
				},
			},
		}
		payload, err := proto.Marshal(ev)
		assert.NoError(t, err)

		msg := &redis.Message{Payload: string(payload)}

		// ACT
		listener.ProcessMessage(context.Background(), msg)

		// ASSERT
		mockWSM.AssertExpectations(t)
	})
}
