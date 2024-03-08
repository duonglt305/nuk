package repositories

import (
	authEntities "duonglt.net/internal/auth/domain/entities"
)

type IUserRepository interface {
	FindByEmail(email string) (*authEntities.User, error)
	Create(user *authEntities.User) (uint64, error)
}
