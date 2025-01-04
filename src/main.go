package main

import (
	"fmt"
	"mist-io/src/server"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	return true
}
func main() {
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	// TODO: change the origin to something safe
	upgrader := websocket.Upgrader{CheckOrigin: checkOrigin}

	// Add routes
	server.AddHandlers(&upgrader)
	server.Initialize(address)
}
