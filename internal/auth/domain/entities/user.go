package entities

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        uint64
	Email     string
	Password  string
	Bio       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
