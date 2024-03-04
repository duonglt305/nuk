package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService[T any] interface {
	Create(id T) (string, error)
	GetID(token string) (*T, error)
	ExtractClaims(token string) (*TokenClaims[T], error)
}

type TokenService[T any] struct {
	secretKey []byte
	lifetime  time.Duration
}

type TokenClaims[T any] struct {
	ID T `json:"id"`
	jwt.MapClaims
}

func NewTokenService[T any](
	secretKey []byte,
	lifetime time.Duration,
) *TokenService[T] {
	return &TokenService[T]{
		secretKey: secretKey,
		lifetime:  lifetime,
	}
}

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

func (s *TokenService[T]) GetID(token string) (*T, error) {
	claims, err := s.ExtractClaims(token)
	if err != nil {
		return nil, err
	}
	return &claims.ID, nil
}

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
