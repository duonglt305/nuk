package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

// NewRedisClient creates a new redis client
func (rc *RedisCache) Connect(redisUrl string) error {
	opts, err := redis.ParseURL(redisUrl)

	if err != nil {
		return err
	}

	rc.client = redis.NewClient(opts)

	if _, err := rc.client.Ping(context.Background()).Result(); err != nil {
		return err
	}

	return nil
}

func (rc *RedisCache) Set(key string, value any, expiration time.Duration) error {
	return rc.client.Set(context.Background(), key, value, expiration).Err()
}

func (rc *RedisCache) Get(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}
