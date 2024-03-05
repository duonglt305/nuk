package auth_repositories

import (
	auth_entities "duonglt.net/internal/domain/auth/entities"
)

type UserRepository interface {
	FindByEmail(email string) (*auth_entities.User, error)
	Create(user *auth_entities.User) (uint64, error)
}
