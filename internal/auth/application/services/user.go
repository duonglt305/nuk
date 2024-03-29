package services

import (
	"time"

	"duonglt.net/internal/auth/application/dtos"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/internal/auth/infrastructure/models"
	"duonglt.net/pkg/utils"
)

type UserService struct {
	sfManager   *utils.SnowflakeManager
	uRepository repositories.UserRepository[models.UserModel, entities.User]
}

func NewUserService(
	sfManager *utils.SnowflakeManager,
	uRepository repositories.UserRepository[models.UserModel, entities.User],
) UserService {
	return UserService{
		sfManager:   sfManager,
		uRepository: uRepository,
	}
}

// Create creates a new user
func (s UserService) Create(data dtos.UserCreate) (*entities.User, error) {
	now := time.Now().UTC()
	user := &entities.User{
		Id:        s.sfManager.New(),
		Email:     data.Email,
		Password:  data.Password,
		Bio:       data.Bio,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if err := user.HashPassword(); err != nil {
		return nil, err
	}
	if err := s.uRepository.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail finds a user by email
func (s UserService) FindByEmail(email string) (entities.User, error) {
	return s.uRepository.FindByEmail(email)
}

// FindByID finds a user by ID
func (s UserService) FindByID(id uint64) (entities.User, error) {
	return s.uRepository.FindById(id)
}

// MarkAsLogged marks a user as logged
func (s UserService) MarkAsLogged(user entities.User) error {
	now := time.Now().UTC()
	user.LoggedAt = &now
	return s.uRepository.Update(&user)
}

// Update updates a user
func (s UserService) Update(data dtos.UserUpdate) (entities.User, error) {
	now := time.Now().UTC()
	user, err := s.uRepository.FindById(data.Id)
	if err != nil {
		return *new(entities.User), err
	}
	user.Bio = data.Bio
	user.UpdatedAt = &now
	if err := s.uRepository.Update(&user); err != nil {
		return *new(entities.User), err
	}
	return user, nil
}
