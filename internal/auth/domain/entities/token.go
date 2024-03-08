package authEntities

import "time"

type Token struct {
	ID            uint64
	Uid           uint64
	AccessTokenID *uint64
	ExpiresAt     time.Time
	CreatedAt     time.Time
}
