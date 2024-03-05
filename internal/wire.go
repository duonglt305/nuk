//go:build wireinject
// +build wireinject

package internal

import (
	auth_services "duonglt.net/internal/application/auth/services"
	infra_auth_repositories "duonglt.net/internal/infrastructure/auth/repositories/pg"
	infra_services "duonglt.net/internal/infrastructure/services"
	presentation_auth "duonglt.net/internal/presentation/auth"
	"duonglt.net/pkg/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func InitializeRouter() (*http.Router, error) {
	wire.Build(
		resolveSfService,
		resolveTokenService,
		infra_auth_repositories.NewTokenRepository,
		auth_services.NewAuthService,
		presentation_auth.NewAuthHttpHandler,
		http.NewRouter,
	)
	return &http.Router{}, nil
}

func resolveTokenService() *infra_services.TokenService[uint64] {
	return infra_services.NewTokenService[uint64](
		[]byte(viper.GetString("JWT_SECRET")),
		viper.GetDuration("JWT_LIFETIME"),
	)
}

func resolveSfService() *infra_services.SfService {
	return infra_services.NewSfService(
		uint16(viper.GetUint("SF_WORKER")),
	)
}
