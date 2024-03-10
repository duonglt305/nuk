package presentation

import (
	"duonglt.net/internal/auth/application/dtos"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"duonglt.net/internal/auth/application/services"
)

type Http struct {
	tokenCreateHandler  tokenCreateHandler
	tokenRefreshHandler tokenRefreshHandler
	registrationHandler registrationHandler
}

func NewHttp(
	userService services.UserService,
	authService services.AuthService,
) Http {
	return Http{
		tokenCreateHandler:  newTokenCreateHandler(userService, authService),
		tokenRefreshHandler: newTokenRefreshHandler(authService),
		registrationHandler: newRegistrationHandler(userService),
	}
}

func (h Http) RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("POST /auth/token", h.tokenCreateHandler)
	mux.Handle("GET /auth/token/refresh", h.tokenRefreshHandler)
	mux.Handle("POST /auth/registration", h.registrationHandler)
}

// TokenCreateHandler is used to handle token creation
type tokenCreateHandler struct {
	userService services.UserService
	authService services.AuthService
}

func newTokenCreateHandler(
	userService services.UserService,
	authService services.AuthService,
) tokenCreateHandler {
	return tokenCreateHandler{userService, authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := dtos.TokenCreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user, err := h.userService.FindByEmail(body.Email)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok := user.ComparePassword(body.Password); !ok {
		http.Error(w, fmt.Errorf("invalid password").Error(), http.StatusBadRequest)
		return
	}
	tk, err := h.authService.CreateToken(user.Id)
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

// TokenRefreshHandler is used to handle token refresh
type tokenRefreshHandler struct {
	authService services.AuthService
}

func newTokenRefreshHandler(authService services.AuthService) tokenRefreshHandler {
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

type registrationHandler struct {
	userService services.UserService
}

func newRegistrationHandler(userService services.UserService) registrationHandler {
	return registrationHandler{userService: userService}
}

func (h registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := dtos.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user, err := h.userService.Create(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("User: %+v", user)
}
