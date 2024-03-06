package authRepositories

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"

	authEntities "duonglt.net/internal/domain/auth/entities"
	authRepositories "duonglt.net/internal/domain/auth/repositories"
	infraServices "duonglt.net/internal/infrastructure/services"
)

// TokenRepository struct is used to define token repository
type TokenRepository struct {
	rdb       *redis.Client
	sfService *infraServices.SfService
}

// NewTokenRepository function is used to create a new token repository
func NewTokenRepository(rdb *redis.Client, sfService *infraServices.SfService) authRepositories.TokenRepository {
	return TokenRepository{rdb: rdb, sfService: sfService}
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
