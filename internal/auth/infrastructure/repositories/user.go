package repositories

import (
	"context"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

// NewUserRepository is a constructor to create a new UserRepository
func NewUserRepository(db *pgx.Conn) repositories.IUserRepository {
	return UserRepository{db: db}
}

// FindById is a method to find user by id
func (rep UserRepository) FindById(id uint64) (*entities.User, error) {
	user := new(entities.User)
	if err := rep.db.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE id = $1",
		id,
	).Scan(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail is a method to find user by email
func (rep UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := new(entities.User)
	if err := rep.db.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(&user.Id, &user.Email, &user.Password, &user.Bio, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

// Save is a method to save user
func (rep UserRepository) Save(user *entities.User) error {
	if _, err := rep.db.Exec(
		context.Background(),
		"INSERT INTO users (id, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Id,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}
