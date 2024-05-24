package entities

import (
	"time"
)

type Token struct {
	ID        uint64    `json:"id"`
	Uid       uint64    `json:"uid"`
	Tid       *uint64   `json:"tid"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
