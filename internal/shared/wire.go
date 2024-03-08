package shared

import (
	sharedServices "duonglt.net/internal/shared/application/services"
	sharedInfrastructure "duonglt.net/internal/shared/infrastructure"
	sharedPresentation "duonglt.net/internal/shared/presentation"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// WireSet variable is used to define wire set
var WireSet = wire.NewSet(
	ResolveRedisClient,
	ResolveSnowflakeService,
	sharedPresentation.NewRouter,
)

// ResolveSnowflakeService function is used to resolve snowflake service
func ResolveSnowflakeService() *sharedServices.SfService {
	return sharedServices.NewSfService(uint16(viper.GetInt("SF_WORKER")))
}

// ResolveRedisClient function is used to resolve redis client
func ResolveRedisClient() *redis.Client {
	return sharedInfrastructure.NewRedisClient(viper.GetString("REDIS_URL"))
}
