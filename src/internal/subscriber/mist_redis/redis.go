package mist_redis

import (
	"context"
	"fmt"
	"log"
	"mist-io/src/internal/logging/logger"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}

var (
	username     = os.Getenv("REDIS_USERNAME")
	password     = os.Getenv("REDIS_PASSWORD")
	hostname     = os.Getenv("REDIS_HOSTNAME")
	port         = os.Getenv("REDIS_PORT")
	eventChannel = os.Getenv("REDIS_NOTIFICATION_CHANNEL")
)

func ConnectToRedis(dbStr string) *redis.Client {
	var client *redis.Client
	ctx := context.Background()

	for client == nil {
		logger.Debug("Initializing Redis client.", "SERVICE", "REDIS")
		client = createClient(dbStr)

		// Perform a health check by setting a key
		result, err := client.Set(ctx, "health-io", "check", 0).Result()

		if err != nil {
			logger.Error(fmt.Sprintf("Failed to connect to Redis: %v", err), "SERVICE", "REDIS")
			logger.Debug("Retrying in 5 seconds...", "SERVICE", "REDIS")
			client.Close()
			client = nil // Reset client to retry connection
			// Wait for 5 seconds before retrying
			<-time.After(5 * time.Second)
		}

		if result == "OK" {
			logger.Debug("Redis connection established", "SERVICE", "REDIS")
		}

		// Clean up the health check key after setting it
		client.Del(ctx, "health-io").Result()

	}

	return client
}

func createClient(dbStr string) *redis.Client {
	var client *redis.Client

	if eventChannel == "" {
		log.Fatal("REDIS_NOTIFICATION_CHANNEL environment variable is not set")
	}

	if username == "" || password == "" || hostname == "" || port == "" || dbStr == "" {
		log.Fatal("Redis connection details are not set in environment variables")
	}
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	for client == nil {
		logger.Debug("Initializing Redis client.", "SERVICE", "REDIS")
		// TODO: the client needs to be fault tolerant, if it fails to connect it should retry after a few seconds
		// If the connection stops working, it should retry to connect
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", hostname, port),
			Username: username,
			Password: password,
			DB:       db, // use default DB
		})
	}

	return client
}
