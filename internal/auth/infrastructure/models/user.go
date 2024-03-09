package models

import (
	"duonglt.net/internal/auth/domain/entities"
	"time"
)

type User struct {
	Id        uint64
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *User) ToEntity() *entities.User {
	return &entities.User{
		Id:        u.Id,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
