package dtos

type ForgotPassword struct {
	Email string `json:"email" validate:"required,email"`
}
