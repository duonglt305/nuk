package auth

import (
	"duonglt.net/internal/auth/application/services"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/internal/auth/infrastructure/models"
	infrasRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	"duonglt.net/internal/auth/presentation"
	sharedServices "duonglt.net/internal/shared/application/services"
	sharedInfrastructure "duonglt.net/internal/shared/infrastructure/db"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveJwtService,
	ResolveAuthService,
	ResolvePgClient,
	ResolveUserRepository,
	infrasRepositories.NewTokenRepository,
	services.NewUserService,
	presentation.NewAuthMiddleware,
	presentation.NewHttp,
)

func ResolveUserRepository(db *sqlx.DB) repositories.UserRepository[models.UserModel, entities.User] {
	return infrasRepositories.NewUserRepository[models.UserModel, entities.User](db)
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

// ResolvePgClient function is used to resolve pg client
func ResolvePgClient() *sqlx.DB {
	return sharedInfrastructure.NewPgClient(viper.GetString("DATABASE_URL"))
}
