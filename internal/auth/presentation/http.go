package presentation

import (
	"net/http"

	vHttp "duonglt.net/pkg/http"

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
	body := dtos.TokenCreate{}
	if err := vHttp.NewValidator(r, &body); err != nil {
		vHttp.Error(w, err)
		return
	}
	user, err := h.uService.FindByEmail(body.Email)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	if ok := user.ComparePassword(body.Password); !ok {
		vHttp.Error(w, err)
		return
	}
	tk, err := h.authService.CreateToken(user.Id)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	if err := h.uService.MarkAsLogged(user); err != nil {
		vHttp.Error(w, err)
		return
	}
	vHttp.Ok(w, tk)
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
	body := dtos.TokenRefresh{}
	if err := vHttp.NewValidator(r, &body); err != nil {
		vHttp.Error(w, err)
		return
	}
	tk, err := h.tokenService.RefreshToken(body.RefreshToken)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	vHttp.Ok(w, tk)
}

type registrationHandler struct {
	uService services.UserService
}

func newRegistrationHandler(userService services.UserService) registrationHandler {
	return registrationHandler{uService: userService}
}

func (h registrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := dtos.UserCreateInput{}
	if err := vHttp.NewValidator(r, &body); err != nil {
		vHttp.Error(w, err)
		return
	}
	user, err := h.uService.Create(body)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	vHttp.Ok(w, user)
}

// profileHandler is used to handle profile
type profileHandler struct {
	uService services.UserService
}

func newProfileHandler(uService services.UserService) profileHandler {
	return profileHandler{uService}
}

func (h profileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("UID").(uint64)
	u, err := h.uService.FindByID(uid)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	vHttp.Ok(w, u)
}

// updateProfileHandler is used to handle profile update
type updateProfileHandler struct {
	uService services.UserService
}

func newUpdateProfileHandler(uService services.UserService) updateProfileHandler {
	return updateProfileHandler{uService}
}

func (h updateProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("UID").(uint64)
	body := dtos.UserUpdateInput{Id: uid}
	if err := vHttp.NewValidator(r, &body); err != nil {
		vHttp.Error(w, err)
		return
	}
	u, err := h.uService.Update(body)
	if err != nil {
		vHttp.Error(w, err)
		return
	}
	vHttp.Ok(w, u)
}
