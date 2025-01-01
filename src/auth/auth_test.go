package auth

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type DummyRequest struct{}

var mockHandler = func(ctx context.Context, req interface{}) (interface{}, error) { return req, nil }

func TestAuthenticateRequest(t *testing.T) {
	t.Run("valid_token", func(t *testing.T) {
		// ARRANGE
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			})

		// Set the Authorization header in http.Header
		headers := http.Header{}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

		// ACT
		tokenAndClaims, err := AuthenticateRequest(headers)

		// ASSERT
		assert.Nil(t, err)
		assert.Equal(t, tokenAndClaims.Token, tokenStr)
	})

	t.Run("invalid_audience", func(t *testing.T) {
		// ARRANGE
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{"invalid-audience"},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			})

		// Set the Authorization header in http.Header
		headers := http.Header{}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid audience claim")
		assert.Nil(t, token)
	})

	t.Run("invalid_issuer", func(t *testing.T) {
		// ARRANGE
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: os.Getenv("MIST_API_JWT_SECRET_KEY"),
			})

		// Set the Authorization header in http.Header
		headers := http.Header{}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid issuer claim")
		assert.Nil(t, token)
	})

	t.Run("invalid_secret_key", func(t *testing.T) {
		// ARRANGE
		tokenStr := createJwtToken(t,
			&CreateTokenParams{
				iss:       os.Getenv("MIST_API_JWT_ISSUER"),
				aud:       []string{os.Getenv("MIST_API_JWT_AUDIENCE")},
				secretKey: "wrong-secret-key",
			})

		// Set the Authorization header in http.Header
		headers := http.Header{}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error parsing token")
		assert.Nil(t, token)
	})

	t.Run("invalid_token_format", func(t *testing.T) {
		// ARRANGE
		headers := http.Header{}
		headers.Set("Authorization", "Bearer bad_token")

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "token is malformed")
		assert.Nil(t, token)
	})

	t.Run("missing_authorization_header", func(t *testing.T) {
		// ARRANGE
		headers := http.Header{} // No Authorization header

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unauthenticated")
		assert.Nil(t, token)
	})

	t.Run("invalid_authorization_bearer_header", func(t *testing.T) {
		// ARRANGE
		headers := http.Header{}
		headers.Set("Authorization", "token invalid")

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid token")
		assert.Nil(t, token)
	})

	t.Run("empty_authorization_bearer_header", func(t *testing.T) {
		// ARRANGE
		headers := http.Header{}
		headers.Set("Authorization", "")

		// ACT
		token, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid token")
		assert.Nil(t, token)
	})

	t.Run("invalid_claims_format_for_audience", func(t *testing.T) {
		// ARRANGE
		// Create a token with invalid format for the "aud" claim (e.g., not an array)
		claims := &jwt.RegisteredClaims{
			Issuer:    os.Getenv("MIST_API_JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(os.Getenv("MIST_API_JWT_SECRET_KEY")))

		headers := http.Header{}
		headers.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

		// ACT
		returnToken, err := AuthenticateRequest(headers)

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid audience claim")
		assert.Nil(t, returnToken)
	})

	t.Run("missing_header_errors", func(t *testing.T) {
		// ARRANGE
		// ACT
		token, err := AuthenticateRequest(http.Header{})

		// ASSERT
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unauthenticated")
		assert.Nil(t, token)
	})
}

func TestGetJWTClaims(t *testing.T) {
	t.Run("can_successfully_get_claims_from_context", func(t *testing.T) {
		// ARRANGE
		claims := &CustomJWTClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:   "dummy issuer",
				Audience: jwt.ClaimStrings{"oo aud"},

				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			UserID: uuid.NewString(),
		}
		ctx := context.Background()
		ctx = context.WithValue(ctx, JwtClaimsContextKey, claims)

		// ACT
		ctxClaims, err := GetJWTClaims(ctx)

		// ASSERT
		assert.NotNil(t, ctxClaims)
		assert.Nil(t, err)
	})

	t.Run("invalid_claims_return_error", func(t *testing.T) {
		// ARRANGE
		ctx := context.Background()
		ctx = context.WithValue(ctx, JwtClaimsContextKey, "boom")

		// ACT
		ctxClaims, err := GetJWTClaims(ctx)

		// ASSERT
		assert.Nil(t, ctxClaims)
		assert.NotNil(t, err)
	})
}