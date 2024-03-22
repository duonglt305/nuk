package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient creates a new redis client
func NewRedisClient(redisUrl string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisUrl)

	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opts)

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}
