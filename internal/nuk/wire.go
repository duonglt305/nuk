//go:build wireinject
// +build wireinject

package nuk

import (
	"github.com/google/wire"
)

func InitializeRouter() (Router, error) {
	wire.Build(
		NewRouter,
	)
	return Router{}, nil
}
