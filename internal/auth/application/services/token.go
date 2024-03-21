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
	accessLifetime       time.Duration
	refreshTokenLifetime time.Duration
}

// NewTokenService function is used to create a new auth service
func NewTokenService(
	sfService *sharedServices.SfService,
	jwtService sharedServices.JwtService[entities.Token],
	tokenRepository repositories.ITokenRepository,
	accessLifetime time.Duration,
	refreshLifetime time.Duration,
) TokenService {
	return TokenService{
		sfService:            sfService,
		jwtService:           jwtService,
		tokenRepository:      tokenRepository,
		accessLifetime:       accessLifetime,
		refreshTokenLifetime: refreshLifetime,
	}
}

// CreateToken function is used to create token
func (serv TokenService) CreateToken(uid uint64) (*dtos.AuthToken, error) {
	createdAt := time.Now().UTC()
	tid := new(uint64)

	access, err := serv.createToken(uid, tid, createdAt, serv.refreshTokenLifetime)
	if err != nil {
		return nil, err
	}
	refresh, err := serv.createToken(uid, tid, createdAt, serv.accessLifetime)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthToken{
		AccessToken:  access,
		RefreshToken: &refresh,
		ExpiresAt:    createdAt.Add(serv.accessLifetime * time.Second).Unix(),
	}, nil
}

// createToken function is used to create a new access token
func (serv TokenService) createToken(
	uid uint64,
	tkid *uint64,
	createdAt time.Time,
	lifetime time.Duration,
) (string, error) {
	id := serv.sfService.New()
	if *tkid == 0 {
		*tkid = id
	}
	tk := entities.Token{
		ID:        id,
		Uid:       uid,
		CreatedAt: createdAt,
		ExpiresAt: createdAt.Add(lifetime * time.Second),
	}
	if id != *tkid {
		tk.Tid = tkid
	}
	return serv.jwtService.Create(tk, tk.ExpiresAt)
}

// RefreshToken function is used to refresh token
func (serv TokenService) RefreshToken(refreshToken string) (*dtos.AuthToken, error) {
	now := time.Now().UTC()
	claims, err := serv.jwtService.ExtractClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	if serv.tokenRepository.IsBlacklisted(*claims.Data.Tid) {
		return nil, errors.New("refresh token is invalid")
	}
	accessExpiresAt := claims.Data.CreatedAt.Add(serv.accessLifetime * time.Second)
	if accessExpiresAt.After(now) {
		if err := serv.tokenRepository.AddToBlacklist(*claims.Data.Tid, claims.Data.ExpiresAt.Sub(now)); err != nil {
			fmt.Printf("failed to add token to blacklist: %+v\n", err)
		}
	}
	return serv.CreateToken(claims.Data.Uid)
}

// VerifyToken function is used to verify token
func (serv TokenService) VerifyToken(token string) (*entities.Token, error) {
	tk := new(entities.Token)
	claims, err := serv.jwtService.ExtractClaims(token)
	if err != nil {
		return tk, err
	}
	*tk = claims.Data
	if tk.Tid != nil || serv.tokenRepository.IsBlacklisted(tk.ID) {
		return tk, errors.New("token is invalid")
	}
	return tk, nil
}

// ExtractRawToken function is used to extract token
func (serv TokenService) ExtractRawToken(r *http.Request) (string, error) {
	reg, err := regexp.Compile("Bearer (.*)")
	if err != nil {
		return "", err
	}
	results := reg.FindStringSubmatch(r.Header.Get("Authorization"))

	if len(results) < 2 {
		return "", errors.New("token is invalid")
	}
	return results[1], nil
}
