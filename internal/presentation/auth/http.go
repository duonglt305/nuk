package presentation_auth

import (
	"encoding/json"
	"net/http"

	auth_services "duonglt.net/internal/application/auth/services"
)

type AuthHttp struct {
	tokenCreateHandler tokenCreateHandler
}

func NewAuthHttpHandler(authService auth_services.AuthService) AuthHttp {
	return AuthHttp{
		tokenCreateHandler: newTokenCreateHandler(authService),
	}
}
func (h AuthHttp) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /token", h.tokenCreateHandler)
}

type tokenCreateHandler struct {
	authService auth_services.AuthService
}

func newTokenCreateHandler(authService auth_services.AuthService) tokenCreateHandler {
	return tokenCreateHandler{authService: authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, err := h.authService.CreateToken(46589833908224)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(map[string]any{"token": token})
	w.Write(b)
}
