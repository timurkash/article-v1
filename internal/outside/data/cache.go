package data

import (
	"context"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/biz"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/test"
)

type cache struct {
	data *Data
	log  *log.Helper
}

func NewCache(data *Data, logger log.Logger) biz.ICache {
	return &cache{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cache) SetEx(ctx context.Context, key string, ttl time.Duration) error {
	if test.IsUnitTest {
		return nil
	}
	count, err := c.data.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return err
	}
	if count == 1 {
		return cerrors.AlreadyPublished
	}
	return c.data.redisClient.Set(ctx, key, nil, ttl).Err()
}
