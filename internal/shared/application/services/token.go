package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenClaims is a struct for token claims
type TokenClaims[T any] struct {
	Data T `json:"id"`
	jwt.MapClaims
}

// JwtService is a service for creating and extracting tokens
type JwtService[T any] struct {
	secretKey []byte
}

// NewJwtService creates a new token service
func NewJwtService[T any](
	secretKey []byte,
) JwtService[T] {
	return JwtService[T]{
		secretKey: secretKey,
	}
}

// Create creates a new token
func (s JwtService[T]) Create(data T, expiresAt time.Time) (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims[T]{
		Data: data,
		MapClaims: jwt.MapClaims{
			"exp": expiresAt.Unix(),
			"iat": now.Unix(),
			"nbf": now.Unix(),
		},
	})
	return token.SignedString(s.secretKey)
}

// ExtractClaims extracts the claims from the token
func (s JwtService[T]) ExtractClaims(token string) (*TokenClaims[T], error) {
	claims := TokenClaims[T]{}
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return s.secretKey, nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	return &claims, nil
}
