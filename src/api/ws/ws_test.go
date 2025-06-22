package ws_test

import (
	"fmt"
	"mist-io/src/api/ws"
	"mist-io/src/auth"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CreateTokenParams struct {
	iss       string
	aud       []string
	secretKey string
	userId    string
}

func createJwtToken(t *testing.T, params *CreateTokenParams) string {
	// Define secret key for signing the token

	// Define JWT claims
	claims := &auth.CustomJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   params.iss,
			Audience: params.aud,

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: params.userId,
	}
	// Create a new token with specified claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	tokenString, err := token.SignedString([]byte(params.secretKey))
	if err != nil {
		t.Fatalf("error signing the token %v", err)
	}
	return tokenString
}

// --- Actual Test ---

func TestWsHandler(t *testing.T) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	deps := ws.WsServerDeps{} // Fill this if needed

	// Set up HTTP server with the WebSocket handler
	server := httptest.NewServer(http.HandlerFunc(ws.WsHandler(upgrader, deps)))
	defer server.Close()

	// Convert HTTP -> WS URL
	wsURL := "ws" + server.URL[len("http"):] + "/ws"

	// ---- Unauthenticated request ----
	t.Run("Error: unauthenticated_request_returns 401", func(t *testing.T) {
		// ACT
		_, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "bad handshake") // Or 401 depending on server
	})

	t.Run("authenticated_request_upgrades_to_websocket", func(t *testing.T) {
		// ARRANGE
		token := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
				userId:    "123", // Provide a value!
			})

		// Pass token in query param
		tokenParam := url.QueryEscape(fmt.Sprintf("Bearer %s", token))
		fullURL := fmt.Sprintf("%s?authorization=%s", wsURL, tokenParam)

		// ACT
		conn, _, err := websocket.DefaultDialer.Dial(fullURL, nil)
		require.NoError(t, err)
		defer conn.Close()

		// ASSERT
		assert.NotNil(t, conn)
	})
}

func TestAddHandlers_RegistersIORoute(t *testing.T) {
	// ARRANGE
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	deps := ws.WsServerDeps{}

	// Reset the default mux to avoid cross-test interference
	http.DefaultServeMux = http.NewServeMux()
	ws.AddHandlers(upgrader, deps)

	// Start test server
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// ACT
	resp, err := http.Get(server.URL + "/io")
	require.NoError(t, err)
	defer resp.Body.Close()

	// ASSERT
	assert.NotEqual(t, http.StatusNotFound, resp.StatusCode, "expected /io route to exist")
}
