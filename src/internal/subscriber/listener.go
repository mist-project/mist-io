package subscriber

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
}

type Job interface {
	Execute(int) error
	Ctx() context.Context
}
