package presentation

import (
	"encoding/json"
	"net/http"

	authServices "duonglt.net/internal/application/auth/services"
)

type AuthHttp struct {
	tokenCreateHandler tokenCreateHandler
}

func NewAuthHttp(authService authServices.AuthService) AuthHttp {
	return AuthHttp{
		tokenCreateHandler: newTokenCreateHandler(authService),
	}
}

func (h AuthHttp) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /token", h.tokenCreateHandler)
}

type tokenCreateHandler struct {
	authService authServices.AuthService
}

func newTokenCreateHandler(authService authServices.AuthService) tokenCreateHandler {
	return tokenCreateHandler{authService: authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, err := h.authService.CreateToken(46589833908224)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(map[string]string{"access_token": token})
	if _, err = w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
