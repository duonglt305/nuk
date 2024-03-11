package dtos

type UserCreateInput struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"bio,omitempty"`
}

type UserUpdateInput struct {
	Id  uint64  `json:"id"`
	Bio *string `json:"bio,omitempty"`
}
