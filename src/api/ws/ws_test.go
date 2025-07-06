package ws_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mist-io/src/api/ws"
	"mist-io/src/auth"
	"mist-io/src/testutil/mocks"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	t.Run("Error: unauthenticated_request_returns_401", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}

		server := httptest.NewServer(http.HandlerFunc(ws.WsHandler(upgrader, deps)))
		defer server.Close()

		wsURL := "ws" + server.URL[len("http"):] + "/ws"

		// ACT
		_, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "bad handshake")
	})

	t.Run("Error:invalid_session_token_errors", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}

		server := httptest.NewServer(http.HandlerFunc(ws.WsHandler(upgrader, deps)))
		defer server.Close()

		wsURL := "ws" + server.URL[len("http"):] + "/ws"
		fullURL := fmt.Sprintf("%s?token=%s", wsURL, "invalid")

		redisGet := redis.NewStringCmd(context.Background())
		redisGet.SetErr(fmt.Errorf("invalid session token"))

		mockRedis.On("Get", mock.Anything, "invalid").Return(redisGet)

		// ACT
		_, _, err := websocket.DefaultDialer.Dial(fullURL, nil)

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "bad handshake")
	})

	t.Run("Error:invalid_claims_errors", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}

		server := httptest.NewServer(http.HandlerFunc(ws.WsHandler(upgrader, deps)))
		defer server.Close()

		wsURL := "ws" + server.URL[len("http"):] + "/ws"
		sessionToken := "some-session-token"
		fullURL := fmt.Sprintf("%s?token=%s", wsURL, sessionToken)
		token := createJwtToken(t, &CreateTokenParams{
			iss:       "booom",
			aud:       []string{"test-audience"},
			secretKey: "test-secret",
			userId:    "123",
		})

		authorization := fmt.Sprintf("Bearer %s", token)
		redisGet := redis.NewStringCmd(context.Background())
		redisGet.SetVal(authorization)

		mockRedis.On("Get", mock.Anything, sessionToken).Return(redisGet)
		mockRedis.On("Del", mock.Anything, []string{sessionToken}).Return(redis.NewIntCmd(context.Background()))

		// ACT
		_, _, err := websocket.DefaultDialer.Dial(fullURL, nil)

		// ASSERT
		require.Error(t, err)
		assert.Contains(t, err.Error(), "bad handshake")
	})

	t.Run("authenticated_request_upgrades_to_websocket", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}
		server := httptest.NewServer(http.HandlerFunc(ws.WsHandler(upgrader, deps)))
		defer server.Close()
		wsURL := "ws" + server.URL[len("http"):] + "/ws"
		token := createJwtToken(t, &CreateTokenParams{
			iss:       os.Getenv("MIST_API_JWT_ISSUER"),
			aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
			secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			userId:    "123",
		})

		sessionToken := "some-session-token"
		authorization := fmt.Sprintf("Bearer %s", token)

		redisGet := redis.NewStringCmd(context.Background())
		redisGet.SetVal(authorization)

		mockRedis.On("Get", mock.Anything, sessionToken).Return(redisGet)
		mockRedis.On("Del", mock.Anything, []string{sessionToken}).Return(redis.NewIntCmd(context.Background()))

		fullURL := fmt.Sprintf("%s?token=%s", wsURL, sessionToken)

		// ACT
		conn, _, err := websocket.DefaultDialer.Dial(fullURL, nil)

		// ASSERT
		require.NoError(t, err)
		defer conn.Close()

		assert.NotNil(t, conn)
	})
}

func TestSessionHandler(t *testing.T) {
	t.Run("Error:returns_401_if_authorization_header_is_missing", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}
		handler := ws.SessionHandler(deps)

		req := httptest.NewRequest("POST", "/ws/session", nil)
		rec := httptest.NewRecorder()

		// ACT
		handler(rec, req)

		// ASSERT
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "Unauthenticated\n", rec.Body.String())
	})

	t.Run("Error:returns_401_if_token_is_invalid", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}
		handler := ws.SessionHandler(deps)

		req := httptest.NewRequest("POST", "/ws/session", nil)
		req.Header.Set("Authorization", "Bearer invalid.token")
		rec := httptest.NewRecorder()

		// ACT
		handler(rec, req)

		// ASSERT
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, "Unauthenticated\n", rec.Body.String())
	})

	t.Run("Error:returns_500_if_redis.Set_fails", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}
		handler := ws.SessionHandler(deps)

		req := httptest.NewRequest("POST", "/ws/session", nil)
		token := createJwtToken(t, &CreateTokenParams{
			iss:       os.Getenv("MIST_API_JWT_ISSUER"),
			aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
			secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			userId:    "123",
		})
		authorization := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", authorization)
		rec := httptest.NewRecorder()

		redisStatus := redis.NewStatusCmd(context.Background())
		redisStatus.SetErr(errors.New("redis failure"))

		mockRedis.On("Set", mock.Anything, mock.Anything, authorization, mock.AnythingOfType("time.Duration")).
			Return(redisStatus)

		// ACT
		handler(rec, req)

		// ASSERT
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Internal Server Error\n", rec.Body.String())
	})

	t.Run("returns 200 and session token if successful", func(t *testing.T) {
		// ARRANGE
		mockRedis := new(mocks.MockRedis)
		deps := ws.WsServerDeps{RedisClient: mockRedis}
		handler := ws.SessionHandler(deps)

		req := httptest.NewRequest("POST", "/ws/session", nil)
		token := createJwtToken(t, &CreateTokenParams{
			iss:       os.Getenv("MIST_API_JWT_ISSUER"),
			aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
			secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			userId:    "123",
		})
		authorization := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", authorization)
		rec := httptest.NewRecorder()

		mockRedis.On("Set", mock.Anything, mock.Anything, authorization, mock.AnythingOfType("time.Duration")).
			Return(redis.NewStatusCmd(context.Background()))

		// ACT
		handler(rec, req)

		// ASSERT
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		var response ws.SessionTokenResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		require.NoError(t, err)
		assert.NotEmpty(t, response.SessionToken)
	})
}

func TestAddHandlers_RegistersIORoute(t *testing.T) {
	// ARRANGE
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	deps := ws.WsServerDeps{}

	// Get the handler returned by AddHandlers
	handler := ws.AddHandlers(upgrader, deps)

	// Start a test server with the returned handler
	server := httptest.NewServer(handler)
	defer server.Close()

	// ACT
	resp, err := http.Get(server.URL + "/io")
	require.NoError(t, err)
	defer resp.Body.Close()

	// ASSERT
	assert.NotEqual(t, http.StatusNotFound, resp.StatusCode, "expected /io route to exist")
}
