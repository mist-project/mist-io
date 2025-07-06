package ws

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"

	"mist-io/src/auth"
	"mist-io/src/internal/faults"
	"mist-io/src/internal/logging/logger"
	"mist-io/src/internal/subscriber/mist_redis"
	"mist-io/src/internal/worker"
	"mist-io/src/message"
)

type WsServerDeps struct {
	WorkerPool  *worker.WorkerPool
	WSManager   WebSocketManager
	RedisClient mist_redis.RedisInterface
}

type WsServer struct {
	Upgrader *websocket.Upgrader
	Deps     WsServerDeps
}

type SessionTokenResponse struct {
	SessionToken string `json:"session_token"`
}

var TTL_SESSION = 60 * time.Second // Session token expiration time. 60 seconds should be more than enough

func WsHandler(upgrader *websocket.Upgrader, deps WsServerDeps) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// First authenticate the request
		// TODO: move this to a separate function that allows more "middlewares"
		// TODO: add more information about the user session, this will eventually be used for multi sessions for a user

		// Authenticate the request
		token := r.URL.Query().Get("token")

		if token == "" {
			http.Error(w, "Missing token in query parameters", http.StatusBadRequest)
			return
		}

		val, err := deps.RedisClient.Get(context.Background(), token).Result()

		if err != nil {
			if err == redis.Nil {
				faults.RedisError("Invalid or expired token", slog.LevelDebug).LogError(context.Background())
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			faults.RedisError("Error getting token from Redis", slog.LevelWarn).LogError(context.Background())
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		deps.RedisClient.Del(context.Background(), token)
		tokenAndClaims, err := auth.AuthenticateRequest(val)

		if err != nil {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}

		// Upgrade HTTP connection to WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			http.Error(w, "Unable to upgrade connection.", http.StatusBadRequest)
			return
		}

		logger.Info("WebSocket connection established for user: %s\n", tokenAndClaims.Claims.UserID)

		deps.WSManager.AddSocketConnection(tokenAndClaims.Claims.UserID, conn)
		defer conn.Close()
		defer deps.WSManager.RemoveSocketConnection(tokenAndClaims.Claims.UserID)

		// TODO: ADD client connection to a dictionary? TBD
		wsConnection := message.WsConnection{
			Conn: conn,
			// Mutex:    &sync.Mutex{}, // TBD if needed
			JwtToken: tokenAndClaims.Token,
			Claims:   tokenAndClaims.Claims,
		}

		wsConnection.Manage()

	}

	return handler
}

func SessionHandler(deps WsServerDeps) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		_, err := auth.AuthenticateRequest(authorization)

		if err != nil {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}

		// TODO: add error checker here, but what would cause it though?
		sessionToken, _ := generateSessionToken()

		_, err = deps.RedisClient.Set(context.Background(), sessionToken, authorization, TTL_SESSION).Result()

		if err != nil {
			faults.RedisError(
				fmt.Sprintf("Error setting session token in Redis: %v", err), slog.LevelWarn,
			).LogError(context.Background())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SessionTokenResponse{
			SessionToken: sessionToken,
		})
	}
}

func AddHandlers(upgrader *websocket.Upgrader, deps WsServerDeps) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // TODO: fix the origin for the app
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true, // if sending cookies/auth headers
	})

	mux := http.NewServeMux()
	mux.Handle("/io", http.HandlerFunc(WsHandler(upgrader, deps)))
	mux.Handle("/ws-session", http.HandlerFunc(SessionHandler(deps)))

	return c.Handler(mux)
}

func generateSessionToken() (string, error) {
	b := make([]byte, 32) // 256 bits
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
