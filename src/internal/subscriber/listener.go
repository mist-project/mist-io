package subscriber

import (
	"context"
	"fmt"
	"log/slog"
	"mist-io/src/api/ws"
	"mist-io/src/internal/faults"
	"mist-io/src/internal/subscriber/mist_redis"
	"mist-io/src/internal/worker"
	"mist-io/src/internal/worker/jobs"
	"mist-io/src/protos/v1/event"
	"os"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type Job interface {
	Execute(int) error
	Ctx() context.Context
}

type RedisListenerOptions struct {
	RedisClient mist_redis.RedisInterface
	WorkerPool  *worker.WorkerPool
	WsManager   ws.WebSocketManager
}

type RedisListener struct {
	redisClient mist_redis.RedisInterface
	wp          *worker.WorkerPool
	wsManager   ws.WebSocketManager
}

func NewRedisListener(opts *RedisListenerOptions) *RedisListener {

	return &RedisListener{
		redisClient: opts.RedisClient,
		wp:          opts.WorkerPool,
		wsManager:   opts.WsManager,
	}
}

func (rs *RedisListener) StartListening() {

	ctx := context.Background()
	pubsub := rs.redisClient.Subscribe(ctx, os.Getenv("REDIS_NOTIFICATION_CHANNEL"))

	go func() {
		for {
			ctx := context.Background()

			msg, err := pubsub.ReceiveMessage(ctx)

			if err != nil {
				faults.UnknownError(fmt.Sprintf("Error receiving message from Redis: %v", err), slog.LevelError).LogError(ctx)
				continue
			}

			rs.ProcessMessage(ctx, msg)
		}
	}()
}

func (rs *RedisListener) ProcessMessage(ctx context.Context, msg *redis.Message) {

	var (
		e           *event.Event = &event.Event{}
		err         error
		content     []byte
		ids         []string
		connections map[string]*websocket.Conn
	)
	err = proto.Unmarshal([]byte(msg.Payload), e)

	if err != nil {
		faults.UnknownError(fmt.Sprintf("Error unmarshalling message: %v", err), slog.LevelError).LogError(ctx)
		return
	}

	if len(e.GetMeta().GetAppusers()) == 0 {
		return
	}

	ids = make([]string, len(e.GetMeta().GetAppusers()))

	for i, u := range e.GetMeta().GetAppusers() {
		ids[i] = u.GetId()
	}

	e.Meta.Appusers = nil // Clear appusers from the payload to avoid sending it back to the users
	content, err = proto.Marshal(e)

	if err != nil {
		faults.MarshallError(fmt.Sprintf("Error marshalling notification payload: %v", err), slog.LevelError).LogError(ctx)
		return
	}

	connections = rs.wsManager.BulkFindSocketConnection(ids)

	for _, conn := range connections {
		rs.wp.AddJob(jobs.NewNotificationJob(content, conn))
	}
}

func (rs *RedisListener) RedisClient() mist_redis.RedisInterface {
	return rs.redisClient
}

func (rs *RedisListener) WorkerPool() *worker.WorkerPool {
	return rs.wp
}

func (rs *RedisListener) WsManager() ws.WebSocketManager {
	return rs.wsManager
}
