package authValueObjects

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    uint64 `json:"expires_at"`
}
