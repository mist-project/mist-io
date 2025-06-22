package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"mist-io/src/auth"
	"mist-io/src/internal/subscriber"
	"mist-io/src/internal/worker"
	"mist-io/src/message"
)

type WsServerDeps struct {
	Redis      subscriber.RedisInterface
	WorkerPool *worker.WorkerPool
}

type WsServer struct {
	Upgrader *websocket.Upgrader
	Deps     WsServerDeps
}

func WsHandler(upgrader *websocket.Upgrader, deps WsServerDeps) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// First authenticate the request
		// TODO: move this to a separate function that allows more "middlewares"
		// TODO: add more information about the user session, this will eventually be used for multi sessions for a user
		// TODO: TECH DEBT- ws session should first be authenticated via a regular REST and then that will return session token which
		// the user sends to this endpoint, and then that token gets verified (single use token, can probably store in memory or redis

		tokenAndClaims, err := auth.AuthenticateRequest(r.URL.Query())

		if err != nil {
			http.Error(w, fmt.Sprintf("Unauthenticated. error: %s", err), http.StatusUnauthorized)
			return
		}

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
			Client:   message.Client{},
		}
		wsConnection.Manage()
	}

	return handler
}

func AddHandlers(upgrader *websocket.Upgrader, deps WsServerDeps) {
	http.HandleFunc("/io", WsHandler(upgrader, deps))
}
