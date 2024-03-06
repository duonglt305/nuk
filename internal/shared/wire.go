package shared

import (
	sharedServices "duonglt.net/internal/shared/application/services"
	sharedInfrastructure "duonglt.net/internal/shared/infrastructure"
	sharedPresentation "duonglt.net/internal/shared/presentation"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// SharedSet variable is used to define wire set
var SharedSet = wire.NewSet(
	ResolveRedisClient,
	ResolveTokenService,
	ResolveSnowflakeService,
	sharedPresentation.NewRouter,
)

// resolveSnowflakeService function is used to resolve snowflake service
func ResolveSnowflakeService() *sharedServices.SfService {
	return sharedServices.NewSfService(uint16(viper.GetInt("SF_WORKER")))
}

// resolveTokenService function is used to resolve token service
func ResolveTokenService() *sharedServices.TokenService[uint64] {
	return sharedServices.NewTokenService[uint64]([]byte(viper.GetString("JWT_SECRET")))
}

// resolveRedisClient function is used to resolve redis client
func ResolveRedisClient() *redis.Client {
	return sharedInfrastructure.NewRedisClient(viper.GetString("REDIS_URL"))
}
