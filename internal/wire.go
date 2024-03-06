//go:build wireinject
// +build wireinject

package internal

import (
	auth "duonglt.net/internal/auth"
	shared "duonglt.net/internal/shared"
	sharedPresentation "duonglt.net/internal/shared/presentation"
	"github.com/google/wire"
)

// InitializeRouter function is used to initialize router
func InitializeRouter() (*sharedPresentation.Router, error) {
	wire.Build(
		shared.WireSet,
		auth.WireSet,
	)
	return &sharedPresentation.Router{}, nil
}
