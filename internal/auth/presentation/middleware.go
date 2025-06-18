package presentation

import (
	"context"
	"net/http"

	"duonglt.net/pkg/cache"
	vHttp "duonglt.net/pkg/http"

	"duonglt.net/internal/auth/application/services"
)

const (
	ContextUidKey = "uid"
)

type AuthMiddleware struct {
	tokenService services.TokenService
	cache        cache.ICache
}

func NewAuthMiddleware(tokenService services.TokenService, cache cache.ICache) AuthMiddleware {
	return AuthMiddleware{tokenService, cache}
}

func (m AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.tokenService.ExtractRawToken(r)
		if err != nil {
			vHttp.Unauthorized(w, err)
			return
		}
		tk, err := m.tokenService.VerifyToken(token)
		if err != nil {
			vHttp.Unauthorized(w, err)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), ContextUidKey, tk.Uid))
		next.ServeHTTP(w, r)
	})
}
