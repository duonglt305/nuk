package services

import (
	"errors"
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
func (s TokenService) CreateToken(uid uint64) (*dtos.AuthToken, error) {
	createdAt := time.Now().UTC()
	tkid := new(uint64)

	access, err := s.createToken(uid, tkid, createdAt, s.refreshTokenLifetime)
	if err != nil {
		return nil, err
	}
	refresh, err := s.createToken(uid, tkid, createdAt, s.accessLifetime)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthToken{
		AccessToken:  access,
		RefreshToken: &refresh,
		ExpiresAt:    createdAt.Add(s.accessLifetime * time.Second).Unix(),
	}, nil
}

// createToken function is used to create a new access token
func (s TokenService) createToken(
	uid uint64,
	tkid *uint64,
	createdAt time.Time,
	lifetime time.Duration,
) (string, error) {
	id := s.sfService.New()
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
		tk.Tkid = tkid
	}
	return s.jwtService.Create(tk, tk.ExpiresAt)
}

// RefreshToken function is used to refresh token
func (s TokenService) RefreshToken(refreshToken string) (*dtos.AuthToken, error) {
	now := time.Now().UTC()
	claims, err := s.jwtService.ExtractClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	if s.tokenRepository.IsBlacklisted(*claims.Data.Tkid) {
		return nil, errors.New("refresh token is invalid")
	}
	accessExpiresAt := claims.Data.CreatedAt.Add(s.accessLifetime * time.Second)
	if accessExpiresAt.After(now) {
		s.tokenRepository.AddToBlacklist(*claims.Data.Tkid, claims.Data.ExpiresAt.Sub(now))
	}
	return s.CreateToken(claims.Data.Uid)
}

// VerifyToken function is used to verify token
func (s TokenService) VerifyToken(token string) (*entities.Token, error) {
	tk := new(entities.Token)
	claims, err := s.jwtService.ExtractClaims(token)
	if err != nil {
		return tk, err
	}
	*tk = claims.Data
	if tk.Tkid != nil || s.tokenRepository.IsBlacklisted(tk.ID) {
		return tk, errors.New("token is invalid")
	}
	return tk, nil
}

// ExtractRawToken function is used to extract token
func (s TokenService) ExtractRawToken(r *http.Request) (string, error) {
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
