//go:build wireinject
// +build wireinject

package internal

import (
	"duonglt.net/pkg/http"
	"github.com/google/wire"
)

func InitializeRouter() (*http.Router, error) {
	wire.Build(
		http.NewRouter,
	)
	return &http.Router{}, nil
}
