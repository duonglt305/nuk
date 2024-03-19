package auth

import (
	"duonglt.net/internal/auth/application/services"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/internal/auth/infrastructure/models"
	infrasRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	"duonglt.net/internal/auth/presentation"
	sharedServices "duonglt.net/internal/shared/application/services"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveJwtService,
	ResolveAuthService,
	ResolveUserRepository,
	infrasRepositories.NewTokenRepository,
	services.NewUserService,
	presentation.NewAuthMiddleware,
	presentation.NewHttpHandler,
)

func ResolveUserRepository(db *sqlx.DB) repositories.UserRepository[models.UserModel, entities.User] {
	return infrasRepositories.NewUserRepository[models.UserModel](db)
}

// ResolveJwtService function is used to resolve token service
func ResolveJwtService() sharedServices.JwtService[entities.Token] {
	return sharedServices.NewJwtService[entities.Token]([]byte(viper.GetString("JWT_SECRET")))
}

// ResolveAuthService function is used to resolve auth service
func ResolveAuthService(
	sfService *sharedServices.SfService,
	jwtService sharedServices.JwtService[entities.Token],
	tokenRepository repositories.ITokenRepository,
) services.TokenService {
	return services.NewTokenService(
		sfService,
		jwtService,
		tokenRepository,
		viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"),
		viper.GetDuration("JWT_REFRESH_TOKEN_LIFETIME"),
	)
}
