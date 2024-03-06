package auth

import (
	authServices "duonglt.net/internal/auth/application/services"
	authRepositories "duonglt.net/internal/auth/domain/repositories"
	authInfrastructureRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	authPresentation "duonglt.net/internal/auth/presentation"
	sharedServices "duonglt.net/internal/shared/application/services"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveAuthService,
	authPresentation.NewHttp,
	authInfrastructureRepositories.NewTokenRepository,
)

// ResolveAuthService function is used to resolve auth service
func ResolveAuthService(
	tokenService *sharedServices.TokenService[uint64],
	tokenRepository authRepositories.ITokenRepository,
) authServices.AuthService {
	return authServices.NewAuthService(
		viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"),
		viper.GetDuration("JWT_REFRESH_TOKEN_LIFETIME"),
		tokenService,
		tokenRepository,
	)
}
