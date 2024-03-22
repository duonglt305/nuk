package repositories

import (
	"fmt"

	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/pkg/db"
	"github.com/jmoiron/sqlx"
)

type UserRepository[M db.IModel[E], E entities.User] struct {
	db.Repository[M, E]
	dbIns *sqlx.DB
}

// NewUserRepository is a constructor to create a new UserRepository
func NewUserRepository[M db.IModel[E], E entities.User](dbIns *sqlx.DB) UserRepository[M, E] {
	return UserRepository[M, E]{
		Repository: db.NewRepository[M](dbIns),
		dbIns:      dbIns,
	}
}

// FindByEmail is a method to find a user by email
func (rep UserRepository[M, E]) FindByEmail(email string) (E, error) {
	var model M
	err := rep.dbIns.Get(&model, fmt.Sprintf("SELECT * FROM %s WHERE email = $1", model.Table()), email)
	return model.ToEntity(), err
}
