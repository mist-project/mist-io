package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"

	"mist-io/src/auth"
	"mist-io/src/message"
)

func WsHandler(upgrader *websocket.Upgrader, clientConn *grpc.ClientConn) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// First authenticate the request
		// TODO: add more information about the user session
		// TODO: move this to a separate function that allows more "middlewares"
		// For example: what device. this will be stored to save sessions
		tokenAndClaims, err := auth.AuthenticateRequest(r.URL.Query())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthenticated. error: %s", err), http.StatusUnauthorized)
			return
		}

		fmt.Println("Establishing new connection...")

		// Upgrade HTTP connection to WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			http.Error(w, "Unable to upgrade connection.", http.StatusBadRequest)
			return
		}
		defer conn.Close()

		// TODO: ADD client connection to a dictionary? TBD
		wsConnection := message.WsConnection{
			Conn: conn,
			// Mutex:    &sync.Mutex{}, // TBD if needed
			JwtToken: tokenAndClaims.Token,
			Claims:   tokenAndClaims.Claims,
			Client:   message.Client{Conn: clientConn},
		}
		wsConnection.Manage()
	}
	return handler
}

func AddHandlers(upgrader *websocket.Upgrader, clientConn *grpc.ClientConn) {
	http.HandleFunc("/io", WsHandler(upgrader, clientConn))
}

func Initialize(address string) {

	log.Printf("Starting WebSocket server on %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Panicf("Error starting server: %v", err)
	}
}
