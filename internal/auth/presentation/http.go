package authPresentation

import (
	"encoding/json"
	"net/http"
	"regexp"

	authServices "duonglt.net/internal/auth/application/services"
)

type Http struct {
	tokenCreateHandler  tokenCreateHandler
	tokenRefreshHandler tokenRefreshHandler
}

func NewHttp(authService authServices.AuthService) Http {
	return Http{
		tokenCreateHandler:  newTokenCreateHandler(authService),
		tokenRefreshHandler: newTokenRefreshHandler(authService),
	}
}

func (h Http) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("GET /token", h.tokenCreateHandler)
	mux.Handle("GET /token/refresh", h.tokenRefreshHandler)
}

type tokenCreateHandler struct {
	authService authServices.AuthService
}

func newTokenCreateHandler(authService authServices.AuthService) tokenCreateHandler {
	return tokenCreateHandler{authService: authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tk, err := h.authService.CreateToken(1900)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(tk)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

type tokenRefreshHandler struct {
	authService authServices.AuthService
}

func newTokenRefreshHandler(authService authServices.AuthService) tokenRefreshHandler {
	return tokenRefreshHandler{authService: authService}
}

func (h tokenRefreshHandler) extractToken(r *http.Request) (string, error) {
	return r.Header.Get("Authorization"), nil
}

func (h tokenRefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reg, err := regexp.Compile("Bearer (.*)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	results := reg.FindStringSubmatch(r.Header.Get("Authorization"))
	if len(results) != 2 {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	tk, err := h.authService.RefreshToken(results[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(tk)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
