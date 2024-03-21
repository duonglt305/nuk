package repositories

import "time"

type ITokenRepository interface {
	AddToBlacklist(uid uint64, expiresIn time.Duration) error
	IsBlacklisted(uid uint64) bool
}
