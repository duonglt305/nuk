package repositories

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

// AddToBlacklist function is used to add token to blacklist
func (rep TokenRepository) AddToBlacklist(uid uint64, expiresIn time.Duration) error {
	key := fmt.Sprintf("%s:%d", blacklistKey, uid)
	rep.rdb.Set(context.Background(), key, true, expiresIn)
	return nil
}
