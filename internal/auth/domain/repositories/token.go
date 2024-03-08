package authRepositories

import "time"

type ITokenRepository interface {
	AddToBlacklist(uid uint64, expiresIn time.Duration) error
}
