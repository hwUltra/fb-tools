package gormx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
)

type TodoCallback func(model any) (int64, error)
type SelectCallback func(key string) (any, error)

type CacheTool struct {
	Cache cache.Cache
	Db    *gorm.DB
	Conf  cache.CacheConf
}

var (
	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("gormCache")
)

func NewCacheTool(conf cache.CacheConf, db *gorm.DB) *CacheTool {
	cc := cache.New(conf, singleFlights, stats, gorm.ErrRecordNotFound)
	return &CacheTool{
		cc,
		db,
		conf,
	}
}

func (m *CacheTool) FormatPrimary(keyPrefix string, primary any) string {
	return fmt.Sprintf("%s%v", keyPrefix, primary)
}

func (m *CacheTool) CreateUpdate(model any, callBack TodoCallback) (int64, error) {
	return callBack(model)
}

func (m *CacheTool) Delete(model any, callBack TodoCallback) (int64, error) {
	return callBack(model)
}

func (m *CacheTool) Select(key string, callBack SelectCallback) (any, error) {
	return callBack(key)
}

func (m *CacheTool) ClearRedisPrefix(keyPrefix string) {
	redisClient := redis.MustNewRedis(m.Conf[0].RedisConf)
	if list, _, err := redisClient.Scan(0, keyPrefix, 0); err == nil {
		for _, item := range list {
			_, _ = redisClient.Del(item)
		}
	}
}
