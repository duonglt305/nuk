package repositories

import (
	"fmt"
	"time"

	authRepositories "duonglt.net/internal/auth/domain/repositories"
	"duonglt.net/pkg/cache"
)

const (
	blacklistKey = "blacklist"
)

// TokenRepository struct is used to define token repository
type TokenRepository struct {
	cache cache.ICache
}

// NewTokenRepository function is used to create a new token repository
func NewTokenRepository(
	cache cache.ICache,
) authRepositories.ITokenRepository {
	return TokenRepository{cache}
}

// AddToBlacklist function is used to add token to blacklist
func (rep TokenRepository) AddToBlacklist(uid uint64, expiresIn time.Duration) error {
	key := fmt.Sprintf("%s:%d", blacklistKey, uid)
	rep.cache.Set(key, true, expiresIn)
	return nil
}

// IsBlacklisted function is used to check if token is blacklisted
func (rep TokenRepository) IsBlacklisted(uid uint64) bool {
	key := fmt.Sprintf("%s:%d", blacklistKey, uid)
	_, err := rep.cache.Get(key)
	return err == nil
}
