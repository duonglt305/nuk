package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceInterface[T int64 | string] interface {
	Create(id T) (string, error)
	Parse(token string) (*T, error)
}

type JWTService[T int64 | string] struct {
	secretKey []byte
	lifetime  time.Duration
}

type jwtClaims[T int64 | string] struct {
	ID T `json:"id"`
	jwt.MapClaims
}

func NewJWTService[T int64 | string](
	secretKey []byte,
	lifetime time.Duration,
) *JWTService[T] {
	return &JWTService[T]{
		secretKey: secretKey,
		lifetime:  lifetime,
	}
}

func (s *JWTService[T]) Create(id T) (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims[T]{
		ID: id,
		MapClaims: jwt.MapClaims{
			"exp": now.Add(s.lifetime).Unix(),
			"iat": now.Unix(),
			"nbf": now.Unix(),
		},
	})
	return token.SignedString(s.secretKey)
}

func (s *JWTService[T]) Parse(token string) (*T, error) {
	claims := jwtClaims[T]{}
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("token invalid or expired")
	}
	return &claims.ID, nil
}
