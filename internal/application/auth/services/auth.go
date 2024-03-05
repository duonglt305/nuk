package auth_services

import (
	"fmt"

	auth_repositories "duonglt.net/internal/infrastructure/auth/repositories/pg"
	infra_services "duonglt.net/internal/infrastructure/services"
)

// AuthService struct is used to define auth service
type AuthService struct {
	tokenService    *infra_services.TokenService[uint64]
	tokenRepository auth_repositories.TokenRepository
}

// NewAuthService function is used to create a new auth service
func NewAuthService(
	tokenService *infra_services.TokenService[uint64],
	tokenRepository auth_repositories.TokenRepository,
) AuthService {
	return AuthService{
		tokenService:    tokenService,
		tokenRepository: tokenRepository,
	}
}

// CreateToken function is used to create a new token
func (s AuthService) CreateToken(uid uint64) (string, error) {
	token, err := s.tokenRepository.CreateToken(uid)
	if err != nil {
		return "", err
	}
	fmt.Printf("token: %+v\n", token)
	return s.tokenService.Create(token.ID)
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
