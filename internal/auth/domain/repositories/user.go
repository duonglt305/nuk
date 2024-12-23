package repositories

import (
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/pkg/db"
)

type UserRepository[M any, E entities.UserEntity] interface {
	db.IRepository[E, M]
	FindByEmail(email string) (E, error)
}
