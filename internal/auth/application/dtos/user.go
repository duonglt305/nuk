package dtos

type UserCreate struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=6"`
	Bio      *string `json:"bio" validate:"omitempty,max=255"`
}

type UserUpdate struct {
	Id  uint64  `json:"id"`
	Bio *string `json:"bio" validate:"omitempty,max=255"`
}
