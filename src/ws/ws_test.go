package ws_test

import (
	"fmt"
	"mist-io/src/auth"
	"mist-io/src/ws"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var upgrader = websocket.Upgrader{}

type CreateTokenParams struct {
	iss       string
	aud       []string
	secretKey string
	userId    string
}

func TestMain(m *testing.M) {
	// ----- INITIALIZE -----
	ws.AddHandlers(&upgrader, nil)
	// ----- EXECUTE TESTS -----
	exitValue := m.Run()

	// ----- CLEANUP -----
	os.Exit(exitValue)
}

func getUrl(s *httptest.Server) string {
	return "ws" + strings.TrimPrefix(s.URL, "http")
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

func TestInitializer(t *testing.T) {
	t.Run("can_create_server_and_succesfully_connect", func(t *testing.T) {
		// ARRANGE
		address := fmt.Sprintf("%s:%s", os.Getenv("TEST_APP_HOST"), os.Getenv("TEST_APP_PORT"))
		go func() {
			ws.Initialize(address)
		}()
		// Wwait for server to be up
		time.Sleep(20 * time.Millisecond)

		// wait for server to initialize
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			})

		wsUrl := fmt.Sprintf("ws://%s/io?authorization=%s", address, url.QueryEscape(fmt.Sprintf("Bearer %s", tokenStr)))

		// ACT
		ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		defer ws.Close()

		// ASSERT
		assert.Nil(t, err)
		assert.NotNil(t, ws)
	})

	t.Run("panics_when_server_cannot_run", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic but got no panic.")
			}
		}()

		// ARRANGE
		ws.Initialize("invalid")
	})
}
func TestWsHandler(t *testing.T) {
	// SETUP
	upgrader := websocket.Upgrader{}
	s := httptest.NewServer(http.HandlerFunc(ws.WsHandler(&upgrader, nil)))
	wsUrl := getUrl(s)
	defer s.Close()

	t.Run("successful_connection", func(t *testing.T) {
		// ARRANGE
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			})

		wsUrl := fmt.Sprintf("%s/io?authorization=%s", getUrl(s), url.QueryEscape(fmt.Sprintf("Bearer %s", tokenStr)))

		// ACT
		ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		defer ws.Close()

		// ASSERT
		assert.Nil(t, err)
		assert.NotNil(t, ws)
	})

	t.Run("missing_authorization_header", func(t *testing.T) {
		// ARRANGE
		wsUrl = getUrl(s)

		// ACT
		_, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)

		// ASSERT
		assert.NotNil(t, err)
	})
}
