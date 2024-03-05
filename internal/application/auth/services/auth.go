package auth_services

import (
	infra_services "duonglt.net/internal/infrastructure/services"
)

type AuthService struct {
	tokenService *infra_services.TokenService[int64]
}

func NewAuthService(tokenService *infra_services.TokenService[int64]) *AuthService {
	return &AuthService{
		tokenService: tokenService,
	}

}
