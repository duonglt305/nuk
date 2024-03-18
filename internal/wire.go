//go:build wireinject
// +build wireinject

package internal

import (
	"duonglt.net/internal/auth"
	authPresentation "duonglt.net/internal/auth/presentation"
	"duonglt.net/internal/shared"
	sharedInfrastructure "duonglt.net/internal/shared/infrastructure/db"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// InitializeRouter function is used to initialize router
func Initialize() (*Router, error) {
	wire.Build(
		ResolveRouter,
		ResolvePgClient,
		shared.WireSet,
		auth.WireSet,
	)
	return &Router{}, nil
}

func ResolveRouter(handler authPresentation.HttpHandler, authenticated authPresentation.AuthMiddleware) *Router {
	return NewRouter(viper.GetString("PORT"), handler, authenticated)
}

// ResolvePgClient function is used to resolve pg client
func ResolvePgClient() (*sqlx.DB, error) {
	return sharedInfrastructure.NewPgClient(viper.GetString("DATABASE_URL"))
}
