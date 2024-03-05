package authEntities

import "time"

type Token struct {
	ID        uint64
	Uid       uint64
	CreatedAt *time.Time
}
