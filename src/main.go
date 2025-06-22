package main

import (
	"fmt"
	"log"
	"mist-io/src/api/ws"
	"mist-io/src/internal/logging/logger"
	"mist-io/src/internal/subscriber"
	"mist-io/src/internal/subscriber/mist_redis"
	"mist-io/src/internal/worker"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func InitializeServer() {
	logger.Debug("Initializing WebSocket server")
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	logger.Debug("Starting Worker Pool")
	wp := worker.NewWorkerPool(4, 2048)
	wp.StartWorkers() // Start the worker pool
	logger.Debug("Worker Pool started")
	defer wp.Stop()

	logger.Debug("Creating WebSocket manager")
	wsManager := ws.NewWSManager()
	logger.Debug("WebSocket manager created")

	upgrader := websocket.Upgrader{CheckOrigin: checkOrigin}
	ws.AddHandlers(&upgrader, ws.WsServerDeps{
		WorkerPool: wp,
		WSManager:  wsManager,
	})

	// Set up REDIS connection and listener
	logger.Debug("Connecting to Redis")
	redisClient := mist_redis.ConnectToRedis(os.Getenv("REDIS_DB"))
	defer redisClient.Close()
	subscriber.NewRedisListener(&subscriber.RedisListenerOptions{
		RedisClient: redisClient,
		WorkerPool:  wp,
		WsManager:   wsManager,
	}).StartListening()
	logger.Debug("Redis listener started")

	logger.Info(fmt.Sprintf("Starting WebSocket server on %s", address))
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}

func main() {
	logger.InitializeLogger()
	InitializeServer()
}

func checkOrigin(r *http.Request) bool {
	// TODO: define a better origin check
	return true
}
