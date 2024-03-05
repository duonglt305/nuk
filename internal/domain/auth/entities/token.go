package auth_entities

import "time"

type Token struct {
	ID        uint64
	Uid       uint64
	RevokedAt *time.Time
	CreatedAt *time.Time
}
