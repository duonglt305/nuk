package authServices

import (
	"time"

	authRepositories "duonglt.net/internal/auth/domain/repositories"
	sharedServices "duonglt.net/internal/shared/application/services"
)

// AuthService struct is used to define auth service
type AuthService struct {
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
	tokenService         *sharedServices.TokenService[uint64]
	tokenRepository      authRepositories.ITokenRepository
}

// NewAuthService function is used to create a new auth service
func NewAuthService(
	accessTokenLifetime time.Duration,
	refreshTokenLifetime time.Duration,
	tokenService *sharedServices.TokenService[uint64],
	tokenRepository authRepositories.ITokenRepository,
) AuthService {
	return AuthService{
		tokenService:         tokenService,
		tokenRepository:      tokenRepository,
		accessTokenLifetime:  accessTokenLifetime,
		refreshTokenLifetime: refreshTokenLifetime,
	}
}

// CreateToken function is used to create a new token
func (s AuthService) CreateToken(uid uint64) (string, error) {
	token, err := s.tokenRepository.Create(uid)
	if err != nil {
		return "", err
	}
	expiresAt := token.CreatedAt.Add(s.accessTokenLifetime)
	return s.tokenService.Create(token.Uid, expiresAt)
}

// CreateRefreshToken function is used to create a new refresh token
func (s AuthService) CreateRefreshToken(uid uint64) (string, error) {
	token, err := s.tokenRepository.Create(uid)
	if err != nil {
		return "", err
	}
	expiresAt := token.CreatedAt.Add(s.refreshTokenLifetime)
	return s.tokenService.Create(token.Uid, expiresAt)
}

// VerifyToken function is used to verify token
func (s AuthService) VerifyToken(token string) (uint64, error) {
	id, err := s.tokenService.GetID(token)
	if err != nil {
		return 0, err
	}
	tk, err := s.tokenRepository.Get(*id)
	if err != nil {
		return 0, err
	}
	return tk.Uid, nil
}
