package presentation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"duonglt.net/internal/auth/application/dtos"
	"duonglt.net/internal/auth/application/services"
)

type HttpHandler struct {
	profileHandler       profileHandler
	tokenCreateHandler   tokenCreateHandler
	tokenRefreshHandler  tokenRefreshHandler
	registrationHandler  registrationHandler
	updateProfileHandler updateProfileHandler
}

func NewHttpHandler(
	uService services.UserService,
	authService services.TokenService,
) HttpHandler {
	return HttpHandler{
		profileHandler:       newProfileHandler(uService),
		tokenCreateHandler:   newTokenCreateHandler(uService, authService),
		tokenRefreshHandler:  newTokenRefreshHandler(authService),
		registrationHandler:  newRegistrationHandler(uService),
		updateProfileHandler: newUpdateProfileHandler(uService),
	}
}

func (h HttpHandler) RegisterHandlers(mux *http.ServeMux, authenticated func(http.Handler) http.Handler) {
	mux.Handle("GET /auth/me", authenticated(h.profileHandler))
	mux.Handle("PUT /auth/me", authenticated(h.updateProfileHandler))
	mux.Handle("POST /auth/token", h.tokenCreateHandler)
	mux.Handle("POST /auth/token/refresh", h.tokenRefreshHandler)
	mux.Handle("POST /auth/registration", h.registrationHandler)
}

// TokenCreateHandler is used to handle token creation
type tokenCreateHandler struct {
	uService    services.UserService
	authService services.TokenService
}

func newTokenCreateHandler(
	userService services.UserService,
	authService services.TokenService,
) tokenCreateHandler {
	return tokenCreateHandler{userService, authService}
}

func (h tokenCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := dtos.TokenCreateInput{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user, err := h.uService.FindByEmail(body.Email)
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
	if err := h.uService.MarkAsLogged(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(tk); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// TokenRefreshHandler is used to handle token refresh
type tokenRefreshHandler struct {
	tokenService services.TokenService
}

func newTokenRefreshHandler(authService services.TokenService) tokenRefreshHandler {
	return tokenRefreshHandler{tokenService: authService}
}

func (h tokenRefreshHandler) extractToken(r *http.Request) (string, error) {
	return r.Header.Get("Authorization"), nil
}

func (h tokenRefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, err := h.tokenService.ExtractRawToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tk, err := h.tokenService.RefreshToken(token)
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
	uService services.UserService
}

func newRegistrationHandler(userService services.UserService) registrationHandler {
	return registrationHandler{uService: userService}
}

func (h registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := dtos.UserCreateInput{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user, err := h.uService.Create(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(user)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// profileHandler is used to handle profile
type profileHandler struct {
	uService services.UserService
}

func newProfileHandler(uService services.UserService) profileHandler {
	return profileHandler{uService}
}

func (h profileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid := r.Context().Value("UID").(uint64)
	u, err := h.uService.FindByID(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(u)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// updateProfileHandler is used to handle profile update
type updateProfileHandler struct {
	uService services.UserService
}

func newUpdateProfileHandler(uService services.UserService) updateProfileHandler {
	return updateProfileHandler{uService}
}

func (h updateProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid := r.Context().Value("UID").(uint64)
	body := dtos.UserUpdateInput{Id: uid}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	u, err := h.uService.Update(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, _ := json.Marshal(u)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
