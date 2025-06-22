package main

import (
	"context"
	"fmt"
	"log"
	"mist-io/src/api/ws"
	"mist-io/src/internal/logging/logger"
	"mist-io/src/internal/subscriber/mist_redis"
	"mist-io/src/internal/worker"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

func InitializeServer() {
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	upgrader := websocket.Upgrader{CheckOrigin: checkOrigin}

	wp := worker.NewWorkerPool(4, 2048)

	wp.StartWorkers() // Start the worker pool
	defer wp.Stop()

	ws.AddHandlers(&upgrader, ws.WsServerDeps{
		Redis:      connectToRedis(),
		WorkerPool: wp,
	})

	logger.Info(fmt.Sprintf("Starting WebSocket server on %s", address))
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}

func main() {
	InitializeServer()
}

func checkOrigin(r *http.Request) bool {
	// TODO: define a better origin check
	return true
}

func connectToRedis() *redis.Client {
	var client *redis.Client
	ctx := context.Background()

	for client == nil {
		logger.Debug("Initializing Redis client.", "SERVICE", "REDIS")
		client = mist_redis.ConnectToRedis(os.Getenv("REDIS_DB"))

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
