package main

import (
	"fmt"
	"log"
	"mist-io/src/ws"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func checkOrigin(r *http.Request) bool {
	return true
}
func main() {
	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	// First establish connection with the backend service
	// TODO: Create multiple channels to be used by clients to avoid channel bottlenecks
	clientConn, err := grpc.NewClient(
		os.Getenv("MIST_BACKEND_APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Panicf("Error communicating with backend service: %v", err)
	}

	// TODO: change the origin to something safe
	upgrader := websocket.Upgrader{CheckOrigin: checkOrigin}

	// Add routes
	ws.AddHandlers(&upgrader, clientConn)
	ws.Initialize(address)
}
