package infras_services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ITokenService is an interface for token service
type ITokenService[T any] interface {
	Create(id T) (string, error)
	GetID(token string) (*T, error)
	ExtractClaims(token string) (*TokenClaims[T], error)
}

// TokenService is a service for creating and extracting tokens
type TokenService[T any] struct {
	secretKey []byte
	lifetime  time.Duration
}

// TokenClaims is a struct for token claims
type TokenClaims[T any] struct {
	ID T `json:"id"`
	jwt.MapClaims
}

// NewTokenService creates a new token service
func NewTokenService[T any](
	secretKey []byte,
	lifetime time.Duration,
) *TokenService[T] {
	return &TokenService[T]{
		secretKey: secretKey,
		lifetime:  lifetime,
	}
}

// Create creates a new token
func (s *TokenService[T]) Create(id T) (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims[T]{
		ID: id,
		MapClaims: jwt.MapClaims{
			"exp": now.Add(s.lifetime).Unix(),
			"iat": now.Unix(),
			"nbf": now.Unix(),
		},
	})
	return token.SignedString(s.secretKey)
}

// GetID gets the id from the token
func (s *TokenService[T]) GetID(token string) (*T, error) {
	claims, err := s.ExtractClaims(token)
	if err != nil {
		return nil, err
	}
	return &claims.ID, nil
}

// ExtractClaims extracts the claims from the token
func (s *TokenService[T]) ExtractClaims(token string) (*TokenClaims[T], error) {
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
