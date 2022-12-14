package conf_redis

import (
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"

	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

func GetRedisClient(redisConf *common.Redis) (*redis.Client, error) {
	if redisConf == nil {
		return nil, cerrors.GetBadConfigError("redis")
	}
	if redisConf.HostPort == "" {
		return nil, cerrors.GetBadConfigError("redis.host_port")
	}
	redisOptions := &redis.Options{
		Addr:     redisConf.HostPort,
		DB:       int(redisConf.Db),
		Username: redisConf.Username,
		Password: redisConf.Password,
	}
	if redisConf.DialTimeout != nil {
		redisOptions.DialTimeout = redisConf.DialTimeout.AsDuration()
	}
	if redisConf.WriteTimeout != nil {
		redisOptions.WriteTimeout = redisConf.WriteTimeout.AsDuration()
	}
	if redisConf.ReadTimeout != nil {
		redisOptions.ReadTimeout = redisConf.ReadTimeout.AsDuration()
	}
	clientRedis := redis.NewClient(redisOptions)
	clientRedis.AddHook(redisotel.TracingHook{})
	return clientRedis, nil
}
