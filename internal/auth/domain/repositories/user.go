package repositories

import "duonglt.net/internal/auth/domain/entities"

type IUserRepository interface {
	FindById(id uint64) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Save(user *entities.User) error
}
