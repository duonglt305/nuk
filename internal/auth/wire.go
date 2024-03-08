package auth

import (
	"duonglt.net/internal/auth/application/services"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	infrasRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	"duonglt.net/internal/auth/presentation"
	sharedServices "duonglt.net/internal/shared/application/services"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveTokenService,
	ResolveAuthService,
	presentation.NewHttp,
	infrasRepositories.NewTokenRepository,
)

// ResolveTokenService function is used to resolve token service
func ResolveTokenService() sharedServices.TokenService[entities.Token] {
	return sharedServices.NewTokenService[entities.Token]([]byte(viper.GetString("JWT_SECRET")))
}

// ResolveAuthService function is used to resolve auth service
func ResolveAuthService(
	sfService *sharedServices.SfService,
	tokenService sharedServices.TokenService[entities.Token],
	tokenRepository repositories.ITokenRepository,
) services.AuthService {
	return services.NewAuthService(
		sfService,
		tokenService,
		tokenRepository,
		viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"),
		viper.GetDuration("JWT_REFRESH_TOKEN_LIFETIME"),
	)
}
