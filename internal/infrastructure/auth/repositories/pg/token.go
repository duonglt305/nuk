package infra_auth_repositories

import (
	"time"

	auth_entities "duonglt.net/internal/domain/auth/entities"
	infra_services "duonglt.net/internal/infrastructure/services"
)

// TokenRepository struct is used to define token repository
type TokenRepository struct {
	sfService *infra_services.SfService
}

// NewTokenRepository function is used to create a new token repository
func NewTokenRepository(sfService *infra_services.SfService) TokenRepository {
	return TokenRepository{sfService: sfService}
}

// Create function is used to create a new token
func (r TokenRepository) CreateToken(uid uint64) (*auth_entities.Token, error) {
	createdAt := time.Now().UTC()
	return &auth_entities.Token{
		ID:        r.sfService.New(),
		Uid:       uid,
		CreatedAt: &createdAt,
	}, nil
}

func (r TokenRepository) Get(uid uint64) (*auth_entities.Token, error) {
	return &auth_entities.Token{}, nil
}
