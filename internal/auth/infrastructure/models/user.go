package models

import (
	"duonglt.net/internal/auth/domain/entities"
	"time"
)

type UserModel struct {
	Id        uint64     `db:"id"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	Bio       *string    `db:"bio"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func (m UserModel) Table() string {
	return "users"
}

func (m UserModel) ToEntity() entities.User {
	return entities.User{
		Id:        m.Id,
		Email:     m.Email,
		Password:  m.Password,
		Bio:       m.Bio,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m UserModel) FromEntity(ent entities.User) any {
	return UserModel{
		Id:        ent.Id,
		Email:     ent.Email,
		Password:  ent.Password,
		Bio:       ent.Bio,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}
