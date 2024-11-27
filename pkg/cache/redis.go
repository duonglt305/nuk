package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

// Connect connects to a Redis server
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

// Set sets a key-value pair with an expiration time
func (rc *RedisCache) Set(key string, value any, expiration time.Duration) error {
	return rc.client.Set(context.Background(), key, value, expiration).Err()
}

// Get gets a value by key
func (rc *RedisCache) Get(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}

// Delete deletes a key
func (rc *RedisCache) Delete(key string) error {
	return rc.client.Del(context.Background(), key).Err()
}

// Lock locks a key with a value and expiration time
func (rc *RedisCache) Lock(key string, expiration time.Duration) (bool, error) {
	val := "lock"
	return rc.client.SetNX(context.Background(), key, val, expiration).Result()
}

// Unlock unlocks a key
func (rc *RedisCache) Unlock(key string) error {
	return rc.client.Del(context.Background(), key).Err()
}
