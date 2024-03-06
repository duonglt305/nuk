package authRepositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	authEntities "duonglt.net/internal/auth/domain/entities"
	authRepositories "duonglt.net/internal/auth/domain/repositories"
	sharedServices "duonglt.net/internal/shared/application/services"
)

// TokenRepository struct is used to define token repository
type TokenRepository struct {
	rdb       *redis.Client
	sfService *sharedServices.SfService
}

// NewTokenRepository function is used to create a new token repository
func NewTokenRepository(
	rdb *redis.Client,
	sfService *sharedServices.SfService,
) authRepositories.ITokenRepository {
	return TokenRepository{rdb, sfService}
}

// Create function is used to create a new token
func (r TokenRepository) Create(uid uint64) (*authEntities.Token, error) {
	createdAt := time.Now().UTC()
	tk := &authEntities.Token{
		ID:        r.sfService.New(),
		Uid:       uid,
		CreatedAt: &createdAt,
	}
	r.rdb.Set(context.Background(), fmt.Sprintf("token:%d", tk.Uid), tk, -1)
	return tk, nil
}

func (r TokenRepository) Get(uid uint64) (*authEntities.Token, error) {
	return &authEntities.Token{}, nil
}
