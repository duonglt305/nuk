package auth

import (
	"duonglt.net/internal/auth/application/services"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/internal/auth/infrastructure/models"
	infrRepositories "duonglt.net/internal/auth/infrastructure/repositories"
	"duonglt.net/internal/auth/presentation"
	"duonglt.net/pkg/jwt"
	"duonglt.net/pkg/utils"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// WireSet is used to wire the dependencies of auth module
var WireSet = wire.NewSet(
	ResolveTokenManager,
	ResolveAuthService,
	ResolveUserRepository,
	infrRepositories.NewTokenRepository,
	services.NewUserService,
	presentation.NewAuthMiddleware,
	presentation.NewHttpHandler,
)

func ResolveUserRepository(db *sqlx.DB) repositories.UserRepository[models.UserModel, entities.User] {
	return infrRepositories.NewUserRepository[models.UserModel](db)
}

// ResolveTokenManager function is used to resolve token service
func ResolveTokenManager() jwt.TokenManager[entities.Token] {
	return jwt.NewTokenManager[entities.Token]([]byte(viper.GetString("JWT_SECRET")))
}

// ResolveAuthService function is used to resolve auth service
func ResolveAuthService(
	sfManager *utils.SnowflakeManager,
	jwtService jwt.TokenManager[entities.Token],
	tokenRepository repositories.ITokenRepository,
) services.TokenService {
	return services.NewTokenService(
		sfManager,
		jwtService,
		tokenRepository,
		viper.GetDuration("JWT_ACCESS_TOKEN_LIFETIME"),
		viper.GetDuration("JWT_REFRESH_TOKEN_LIFETIME"),
	)
}
