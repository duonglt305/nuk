package shared

import (
	"duonglt.net/internal/shared/application/services"
	"duonglt.net/internal/shared/infrastructure/cache"
	"duonglt.net/internal/shared/presentation"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// WireSet variable is used to define wire set
var WireSet = wire.NewSet(
	ResolveRedisClient,
	ResolveSnowflakeService,
	presentation.NewRouter,
)

// ResolveSnowflakeService function is used to resolve snowflake service
func ResolveSnowflakeService() *services.SfService {
	return services.NewSfService(uint16(viper.GetInt("SF_WORKER")))
}

// ResolveRedisClient function is used to resolve redis client
func ResolveRedisClient() *redis.Client {
	return cache.NewRedisClient(viper.GetString("REDIS_URL"))
}
