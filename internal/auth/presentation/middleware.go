package presentation

import (
	"context"
	"duonglt.net/internal/auth/application/services"
	"net/http"
)

type AuthMiddleware struct {
	tokenService services.TokenService
}

func NewAuthMiddleware(tokenService services.TokenService) AuthMiddleware {
	return AuthMiddleware{tokenService}
}

func (m AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := m.tokenService.ExtractToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		uid, err := m.tokenService.VerifyToken(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "UID", uid))
		next.ServeHTTP(w, r)
	})
}
