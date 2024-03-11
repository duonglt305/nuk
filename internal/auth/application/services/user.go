package services

import (
	"duonglt.net/internal/auth/application/dtos"
	"duonglt.net/internal/auth/domain/entities"
	"duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/internal/auth/infrastructure/models"
	sharedServices "duonglt.net/internal/shared/application/services"
	"time"
)

type UserService struct {
	sfService   *sharedServices.SfService
	uRepository repositories.UserRepository[models.UserModel, entities.User]
}

func NewUserService(
	sfService *sharedServices.SfService,
	uRepository repositories.UserRepository[models.UserModel, entities.User],
) UserService {
	return UserService{
		sfService:   sfService,
		uRepository: uRepository,
	}
}

// Create creates a new user
func (s UserService) Create(data dtos.UserCreateInput) (*entities.User, error) {
	now := time.Now().UTC()
	user := &entities.User{
		Id:        s.sfService.NewSFID(),
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

// Update updates a user
func (s UserService) Update(data dtos.UserUpdateInput) (entities.User, error) {
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
