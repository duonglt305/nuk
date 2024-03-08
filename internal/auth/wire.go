package auth

import (
	authServices "duonglt.net/internal/auth/application/services"
	authEntities "duonglt.net/internal/auth/domain/entities"
	authRepositories "duonglt.net/internal/auth/domain/repositories"
	authInfrastructureRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	authPresentation "duonglt.net/internal/auth/presentation"
	sharedServices "duonglt.net/internal/shared/application/services"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveTokenService,
	ResolveAuthService,
	authPresentation.NewHttp,
	authInfrastructureRepositories.NewTokenRepository,
)

// ResolveTokenService function is used to resolve token service
func ResolveTokenService() *sharedServices.TokenService[authEntities.Token] {
	return sharedServices.NewTokenService[authEntities.Token]([]byte(viper.GetString("JWT_SECRET")))
}

// ResolveAuthService function is used to resolve auth service
func ResolveAuthService(
	sfService *sharedServices.SfService,
	tokenService *sharedServices.TokenService[authEntities.Token],
	tokenRepository authRepositories.ITokenRepository,
) authServices.AuthService {
	return authServices.NewAuthService(
		sfService,
		tokenService,
		tokenRepository,
		viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"),
		viper.GetDuration("JWT_REFRESH_TOKEN_LIFETIME"),
	)
}
