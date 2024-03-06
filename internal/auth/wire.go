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

var AuthSet = wire.NewSet(
	ResolveAuthService,
	authPresentation.NewHttp,
	authInfrastructureRepositories.NewTokenRepository,
)

// resolveAuthService function is used to resolve auth service
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
