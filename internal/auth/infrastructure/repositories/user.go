package repositories

import (
	"duonglt.net/internal/auth/domain/entities"
	shared "duonglt.net/internal/shared/infrastructure"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository[M shared.IModel[E], E entities.User] struct {
	shared.Repository[M, E]
	db *sqlx.DB
}

// NewUserRepository is a constructor to create a new UserRepository
func NewUserRepository[M shared.IModel[E], E entities.User](db *sqlx.DB) UserRepository[M, E] {
	return UserRepository[M, E]{
		Repository: shared.NewRepository[M, E](db),
		db:         db,
	}
}

// FindByEmail is a method to find a user by email
func (rep UserRepository[M, E]) FindByEmail(email string) (E, error) {
	var model M
	err := rep.db.Get(&model, fmt.Sprintf("SELECT * FROM %s WHERE email = $1", model.Table()), email)
	return model.ToEntity(), err
}
