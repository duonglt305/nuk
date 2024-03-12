package dtos

type AuthToken struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	ExpiresAt    int64   `json:"expires_at"`
}

type TokenCreateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenRefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}
