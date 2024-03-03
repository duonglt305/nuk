package auth_services

import (
	infra_services "duonglt.net/internal/infrastructure/services"
)

type AuthService struct {
	jwtService *infra_services.JWTService[int64]
}

func NewAuthService(jwtService *infra_services.JWTService[int64]) *AuthService {
	return &AuthService{
		jwtService: jwtService,
	}

}
