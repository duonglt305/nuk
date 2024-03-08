package authRepositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	authRepositories "duonglt.net/internal/auth/domain/repositories"
)

const (
	blacklistKey = "blacklist"
)

// TokenRepository struct is used to define token repository
type TokenRepository struct {
	rdb *redis.Client
}

// NewTokenRepository function is used to create a new token repository
func NewTokenRepository(
	rdb *redis.Client,
) authRepositories.ITokenRepository {
	return TokenRepository{rdb}
}

// Create function is used to create a new token
func (rep TokenRepository) AddToBlacklist(uid uint64, expiresIn time.Duration) error {
	key := fmt.Sprintf("%s:%d", blacklistKey, uid)
	rep.rdb.Set(context.Background(), key, true, expiresIn)
	return nil
}
