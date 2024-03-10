package services

import (
	"duonglt.net/internal/auth/application/dtos"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	sharedServices "duonglt.net/internal/shared/application/services"
	"time"
)

type UserService struct {
	sfService      *sharedServices.SfService
	userRepository repositories.IUserRepository
}

func NewUserService(
	sfService *sharedServices.SfService,
	userRepository repositories.IUserRepository,
) UserService {
	return UserService{
		sfService:      sfService,
		userRepository: userRepository,
	}
}

// Create creates a new user
func (s UserService) Create(data dtos.CreateUserRequest) (*entities.User, error) {
	now := time.Now().UTC()
	user := &entities.User{
		Id:        s.sfService.NewSFID(),
		Email:     data.Email,
		Password:  data.Password,
		Bio:       &data.Bio,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	if err := s.userRepository.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) FindByEmail(email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(email)
}
