package authRepositories

import (
	authEntities "duonglt.net/internal/auth/domain/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*authEntities.User, error)
	Create(user *authEntities.User) (uint64, error)
}
