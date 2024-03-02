package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const ()

type JWTService[T int64 | string] struct {
	secretKey []byte
	issuer    string
	lifetime  time.Duration
}

func NewJWTService[T int64 | string](
	secretKey []byte,
	issuer string,
	lifetime time.Duration,
) *JWTService[T] {
	return &JWTService[T]{
		secretKey: secretKey,
		issuer:    issuer,
		lifetime:  lifetime,
	}
}

func (s *JWTService[T]) Create(identifier T) (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  identifier,
		"iss": s.issuer,
		"exp": now.Add(s.lifetime).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
	})
	return token.SignedString(s.secretKey)
}

func (s *JWTService[T]) Parse(token string) (*T, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}
	id, ok := claims["id"]
	if !ok {
		return nil, fmt.Errorf("missing id claim")
	}
	identifier, ok := id.(T)
	if !ok {
		return nil, fmt.Errorf("invalid id claim")
	}
	return &identifier, nil
}
