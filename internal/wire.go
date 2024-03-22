//go:build wireinject
// +build wireinject

package internal

import (
	"duonglt.net/internal/auth"
	authPresentation "duonglt.net/internal/auth/presentation"
	"duonglt.net/pkg/cache"
	"duonglt.net/pkg/db"
	"duonglt.net/pkg/utils"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// InitializeRouter function is used to initialize router
func Initialize() (*Router, error) {
	wire.Build(
		ResolveRouter,
		ResolveDatabase,
		ResolveRedisClient,
		ResolveSnowflakeManager,
		auth.WireSet,
	)
	return &Router{}, nil
}

func ResolveRouter(handler authPresentation.HttpHandler, authenticated authPresentation.AuthMiddleware) *Router {
	return NewRouter(viper.GetString("PORT"), handler, authenticated)
}

// ResolveSnowflakeService function is used to resolve snowflake service
func ResolveSnowflakeManager() *utils.SnowflakeManager {
	return utils.NewSfService(uint16(viper.GetInt("SF_WORKER")))
}

// ResolveDatabase function is used to resolve pg client
func ResolveDatabase() (*sqlx.DB, error) {
	dbIns, err := db.New(viper.GetString("DATABASE_DRIVER"), viper.GetString("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return dbIns.Get(), nil
}

func ResolveRedisClient() (*redis.Client, error) {
	return cache.NewRedisClient(viper.GetString("REDIS_URL"))
}
