package entities

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        uint64     `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Bio       *string    `json:"bio"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// HashPassword hashes the user's password using bcrypt
func (u User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

// ComparePassword compares the user's password with the provided password
func (u User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
