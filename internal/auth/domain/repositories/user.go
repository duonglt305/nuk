package repositories

import (
	"duonglt.net/internal/auth/domain/entities"
	shared "duonglt.net/internal/shared/domain"
)

type UserRepository[M any, E entities.User] interface {
	shared.IRepository[E, M]
	FindByEmail(email string) (E, error)
}
