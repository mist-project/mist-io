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

	// ----- WORKER POOL SETUP -----
	logger.Debug("Starting Worker Pool")
	wp := worker.NewWorkerPool(4, 2048)
	wp.StartWorkers() // Start the worker pool
	logger.Debug("Worker Pool started")
	defer wp.Stop()

	// ----- WEBSOCKET MANAGER SETUP -----
	logger.Debug("Creating WebSocket manager")
	wsManager := ws.NewWSManager()
	logger.Debug("WebSocket manager created")

	// ----- REDIS SETUP -----
	logger.Debug("Connecting to Redis")
	redisClient := mist_redis.ConnectToRedis(os.Getenv("REDIS_DB"))
	defer redisClient.Close()
	subscriber.NewRedisListener(&subscriber.RedisListenerOptions{
		RedisClient: redisClient,
		WorkerPool:  wp,
		WsManager:   wsManager,
	}).StartListening()
	logger.Debug("Redis listener started")

	// ----- WEBSOCKET HANDLER SETUP -----
	upgrader := websocket.Upgrader{CheckOrigin: checkOrigin}
	h := ws.AddHandlers(&upgrader, ws.WsServerDeps{
		WorkerPool:  wp,
		WSManager:   wsManager,
		RedisClient: redisClient,
	})

	// ----- HTTP SERVER SETUP -----
	logger.Info(fmt.Sprintf("Starting WebSocket server on %s", address))
	if err := http.ListenAndServe(address, h); err != nil {
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
