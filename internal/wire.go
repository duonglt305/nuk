//go:build wireinject
// +build wireinject

package internal

import (
	authServices "duonglt.net/internal/application/auth/services"
	infrastructure "duonglt.net/internal/infrastructure"
	authRepositories "duonglt.net/internal/infrastructure/auth/repositories"
	infraServices "duonglt.net/internal/infrastructure/services"
	presentation "duonglt.net/internal/presentation/auth"
	"duonglt.net/pkg/http"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitializeRouter() (*http.Router, error) {
	wire.Build(
		resolveSfService,
		resolveTokenService,
		resolveRedisClient,
		authRepositories.NewTokenRepository,
		authServices.NewAuthService,
		presentation.NewAuthHttp,
		http.NewRouter,
	)
	return &http.Router{}, nil
}

func resolveTokenService() *infraServices.TokenService[uint64] {
	return infraServices.NewTokenService[uint64]([]byte(viper.GetString("JWT_SECRET")))
}

func resolveSfService() *infraServices.SfService {
	return infraServices.NewSfService(
		uint16(viper.GetUint("SF_WORKER")),
	)
}

func resolveRedisClient() *redis.Client {
	return infrastructure.NewRedisClient(
		viper.GetString("REDIS_URL"),
	)
}
