package repositories

import (
	"context"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/infrastructure/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return UserRepository{db: db}
}

func (rep UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := new(models.User)
	err := rep.db.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user.ToEntity(), nil
}

func (rep UserRepository) Create(user *entities.User) error {
	uModel := models.User{
		Id:        user.Id,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	if _, err := rep.db.Exec(
		context.Background(),
		"INSERT INTO users (id, email, password) VALUES ($1, $2, $3, $4, $5)",
		uModel.Id,
		uModel.Email,
		uModel.Password,
		uModel.CreatedAt,
		uModel.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}
