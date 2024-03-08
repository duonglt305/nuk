package services

import (
	"duonglt.net/internal/auth/application/dtos"
	"fmt"
	"time"

	authEntities "duonglt.net/internal/auth/domain/entities"
	authRepositories "duonglt.net/internal/auth/domain/repositories"
	sharedServices "duonglt.net/internal/shared/application/services"
)

// AuthService struct is used to define auth service
type AuthService struct {
	sfService            *sharedServices.SfService
	tokenService         sharedServices.TokenService[authEntities.Token]
	tokenRepository      authRepositories.ITokenRepository
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
}

// NewAuthService function is used to create a new auth service
func NewAuthService(
	sfService *sharedServices.SfService,
	tokenService sharedServices.TokenService[authEntities.Token],
	tokenRepository authRepositories.ITokenRepository,
	accessTokenLifetime time.Duration,
	refreshTokenLifetime time.Duration,
) AuthService {
	return AuthService{
		sfService:            sfService,
		tokenService:         tokenService,
		tokenRepository:      tokenRepository,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

// CreateToken function is used to create token
func (s AuthService) CreateToken(uid uint64) (*dtos.AuthToken, error) {
	createdAt := time.Now().UTC()
	var accessToken, refreshToken string
	var err error
	accessTokenID := s.sfService.New()
	if accessToken, err = s.createAccessToken(uid, accessTokenID, createdAt); err != nil {
		return nil, err
	}
	if refreshToken, err = s.createAccessToken(uid, accessTokenID, createdAt); err != nil {
		return nil, err
	}
	return &dtos.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: &refreshToken,
		ExpiresAt:    uint64(createdAt.Add(s.accessTokenLifetime * time.Second).Unix()),
	}, nil
}

// createAccessToken function is used to create a new access token
func (s AuthService) createAccessToken(uid, accessTokenID uint64, createdAt time.Time) (string, error) {
	token := authEntities.Token{
		ID:        accessTokenID,
		Uid:       uid,
		CreatedAt: createdAt,
		ExpiresAt: createdAt.Add(s.accessTokenLifetime * time.Second),
	}

	return s.tokenService.Create(token, token.ExpiresAt)
}

// createRefreshToken function is used to create a new refresh token
func (s AuthService) createRefreshToken(uid, accessTokenID uint64, createdAt time.Time) (string, error) {
	token := authEntities.Token{
		ID:            s.sfService.New(),
		Uid:           uid,
		AccessTokenID: &accessTokenID,
		CreatedAt:     createdAt,
		ExpiresAt:     createdAt.Add(s.refreshTokenLifetime * time.Second),
	}

	return s.tokenService.Create(token, token.ExpiresAt)
}

// RefreshToken function is used to refresh token
func (s AuthService) RefreshToken(refreshToken string) (*dtos.AuthToken, error) {
	now := time.Now().UTC()
	claims, err := s.tokenService.ExtractClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	accessTokenExpiresAt := claims.Data.CreatedAt.Add(s.accessTokenLifetime * time.Second)
	if accessTokenExpiresAt.After(now) {
		if err := s.tokenRepository.AddToBlacklist(*claims.Data.AccessTokenID, accessTokenExpiresAt.Sub(now)); err != nil {
			fmt.Printf("failed to add to blacklist: %+v\n", err)
		}
	}
	// TODO: Generate new access token

	return nil, nil
}

// VerifyToken function is used to verify token
func (s AuthService) VerifyToken(token string) (uint64, error) {
	return 0, nil
}
