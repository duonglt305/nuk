package presentation

import (
	"context"
	"net/http"

	vHttp "duonglt.net/pkg/http"

	"duonglt.net/internal/auth/application/services"
)

const (
	ContextUidKey = "UID"
)

type AuthMiddleware struct {
	tokenService services.TokenService
}

func NewAuthMiddleware(tokenService services.TokenService) AuthMiddleware {
	return AuthMiddleware{tokenService}
}

func (m AuthMiddleware) Handle(next http.Handler) http.Handler {
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
