package services

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"duonglt.net/internal/auth/application/dtos"

	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	sharedServices "duonglt.net/internal/shared/application/services"
)

// TokenService struct is used to define auth service
type TokenService struct {
	sfService            *sharedServices.SfService
	jwtService           sharedServices.JwtService[entities.Token]
	tokenRepository      repositories.ITokenRepository
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

// NewTokenService function is used to create a new auth service
func NewTokenService(
	sfService *sharedServices.SfService,
	jwtService sharedServices.JwtService[entities.Token],
	tokenRepository repositories.ITokenRepository,
	accessTokenLifetime time.Duration,
	refreshTokenLifetime time.Duration,
) TokenService {
	return TokenService{
		sfService:            sfService,
		jwtService:           jwtService,
		tokenRepository:      tokenRepository,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

// CreateToken function is used to create token
func (s TokenService) CreateToken(uid uint64) (*dtos.AuthToken, error) {
	createdAt := time.Now().UTC()
	rfid := s.sfService.NewSFID()

	accessToken, err := s.createToken(uid, nil, createdAt, s.refreshTokenLifetime)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.createToken(uid, &rfid, createdAt, s.accessTokenLifetime)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: &refreshToken,
		ExpiresAt:    createdAt.Add(s.accessTokenLifetime * time.Second).Unix(),
	}, nil
}

// createToken function is used to create a new access token
func (s TokenService) createToken(
	uid uint64,
	refreshTokenId *uint64,
	createdAt time.Time,
	lifetime time.Duration,
) (string, error) {
	id := s.sfService.NewSFID()
	if refreshTokenId != nil {
		id = *refreshTokenId
	}
	token := entities.Token{
		Id:             id,
		Uid:            uid,
		RefreshTokenId: refreshTokenId,
		CreatedAt:      createdAt,
		ExpiresAt:      createdAt.Add(lifetime * time.Second),
	}

	return s.jwtService.Create(token, token.ExpiresAt)
}

// RefreshToken function is used to refresh token
func (s TokenService) RefreshToken(refreshToken string) (*dtos.AuthToken, error) {
	now := time.Now().UTC()
	claims, err := s.jwtService.ExtractClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	accessTokenExpiresAt := claims.Data.CreatedAt.Add(s.accessTokenLifetime * time.Second)
	if accessTokenExpiresAt.After(now) {
		fmt.Println("Access token is still valid")
		return nil, nil
	}
	// TODO: Generate new access token

	return nil, nil
}

// VerifyToken function is used to verify token
func (s TokenService) VerifyToken(token string) (entities.Token, error) {
	claims, err := s.jwtService.ExtractClaims(token)
	if err != nil {
		return *new(entities.Token), err
	}
	return claims.Data, nil
}

// ExtractRawToken function is used to extract token
func (s TokenService) ExtractRawToken(r *http.Request) (string, error) {
	reg, err := regexp.Compile("Bearer (.*)")
	if err != nil {
		return "", err
	}
	results := reg.FindStringSubmatch(r.Header.Get("Authorization"))

	if len(results) < 2 {
		return "", errors.New("invalid token")
	}
	return results[1], nil
}
