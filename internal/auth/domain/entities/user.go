package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	Id        uint64     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Bio       *string    `json:"bio"`
	LoggedAt  *time.Time `json:"logged_at"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// HashPassword hashes the user's password using bcrypt
func (u *UserEntity) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

// ComparePassword compares the user's password with the provided password
func (u UserEntity) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
