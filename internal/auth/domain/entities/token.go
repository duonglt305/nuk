package entities

import "time"

type Token struct {
	Id             uint64
	Uid            uint64
	RefreshTokenId *uint64
	ExpiresAt      time.Time
	CreatedAt      time.Time
}
