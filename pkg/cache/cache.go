package cache

import (
	"sync"
	"time"
)

type ICache interface {
	Connect(cacheUrl string) error
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

const (
	RedisDriver = "redis"
	MemDriver   = "mem"
)

var (
	cacheIns ICache
	once     sync.Once
)

func New(driver string, cacheUrl string) (ICache, error) {
	var err error
	once.Do(func() {
		switch driver {
		case RedisDriver:
			cacheIns = &RedisCache{}
		}
		err = cacheIns.Connect(cacheUrl)
	})
	return cacheIns, err
}
