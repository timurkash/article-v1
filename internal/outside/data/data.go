package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/conf"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/migrate"
	"gitlab.com/mcsolutions/find-psy/back/common/cent"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/config/conf_ent"
	"gitlab.com/mcsolutions/find-psy/back/common/config/conf_redis"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo, NewCache)

// Data .
type Data struct {
	relational  *ent.Client
	redisClient *redis.Client
}

// NewData .
func NewData(confData *conf.Data, logger log.Logger) (*Data, func(), error) {
	if confData == nil {
		return nil, nil, cerrors.GetBadConfigError("data")
	}
	driver, err := conf_ent.GetRelationalDriver(confData.Relational)
	if err != nil {
		return nil, nil, err
	}
	logHelper := log.NewHelper(logger)
	client := ent.NewClient(ent.Driver(cent.DebugWithContext(driver, logHelper)))
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		logHelper.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	redisClient, err := conf_redis.GetRedisClient(confData.Redis)
	if err != nil {
		return nil, nil, err
	}
	data := &Data{
		relational:  client,
		redisClient: redisClient,
	}
	cleanup := func() {
		logHelper.Info("closing the data resources")
		if err := data.relational.Close(); err != nil {
			logHelper.Error(err)
		}
	}
	return data, cleanup, nil
}
