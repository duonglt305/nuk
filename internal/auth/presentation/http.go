package authPresentation

import (
	"encoding/json"
	"net/http"

	authServices "duonglt.net/internal/auth/application/services"
)

type Http struct {
	tokenCreateHandler tokenCreateHandler
}

func NewHttp(authService authServices.AuthService) Http {
	return Http{
		tokenCreateHandler: newTokenCreateHandler(authService),
	}
}

func (h Http) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /token", h.tokenCreateHandler)
}

type tokenCreateHandler struct {
	authService authServices.AuthService
}

func newTokenCreateHandler(authService authServices.AuthService) tokenCreateHandler {
	return tokenCreateHandler{authService: authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(map[string]string{})
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
