package authServices

import (
	authRepositories "duonglt.net/internal/domain/auth/repositories"
	infraServices "duonglt.net/internal/infrastructure/services"
	"github.com/spf13/viper"
)

// AuthService struct is used to define auth service
type AuthService struct {
	tokenService    *infraServices.TokenService[uint64]
	tokenRepository authRepositories.TokenRepository
}

// NewAuthService function is used to create a new auth service
func NewAuthService(
	tokenService *infraServices.TokenService[uint64],
	tokenRepository authRepositories.TokenRepository,
) AuthService {
	return AuthService{
		tokenService:    tokenService,
		tokenRepository: tokenRepository,
	}
}

// CreateToken function is used to create a new token
func (s AuthService) CreateToken(uid uint64) (string, error) {
	token, err := s.tokenRepository.Create(uid)
	if err != nil {
		return "", err
	}
	expiresAt := token.CreatedAt.Add(viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"))
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
