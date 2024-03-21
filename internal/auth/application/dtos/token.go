package dtos

type AuthToken struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	ExpiresAt    int64   `json:"expires_at"`
}

type TokenCreate struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenRefresh struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
