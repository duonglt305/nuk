package infrastructure

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient creates a new redis client
func NewRedisClient(redisUrl string) *redis.Client {
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		fmt.Printf("failed to parse redis url: %+v\n", err)
		return nil
	}
	return redis.NewClient(opts)
}
