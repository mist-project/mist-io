package main

import (
	"fmt"
	"mist-io/src/server"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	upgrader := websocket.Upgrader{}

	// Add routes
	server.AddHandlers(&upgrader)
	server.Initialize(address)
}
