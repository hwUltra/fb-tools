package cachex

import (
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
)

type TodoCallback func(model any) (int64, error)
type SelectCallback func(key string) (any, error)

type CacheTodo interface {
	FormatPrimary(keyPrefix string, primary any) string
	CreateUpdate(model any, callBack TodoCallback) (int64, error)
	Delete(model any, callBack TodoCallback) (int64, error)
	Select(key string, callBack SelectCallback) (any, error)
	ClearRedisPrefix(keyPrefix string)
}

type Store struct {
	Cache     cache.Cache
	CacheConf cache.CacheConf
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("fb-cache")
	cacheErr      = errors.New("404 not found")
)

func NewStore(conf cache.CacheConf) *Store {
	cc := cache.New(conf, singleFlights, stats, cacheErr)
	return &Store{
		cc,
		conf,
	}
}

func (s *Store) FormatPrimary(keyPrefix string, primary any) string {
	return fmt.Sprintf("%s%v", keyPrefix, primary)
}

func (s *Store) ClearRedisPrefix(keyPrefix string) {
	redisClient := redis.MustNewRedis(s.CacheConf[0].RedisConf)
	if list, _, err := redisClient.Scan(0, keyPrefix, 0); err == nil {
		for _, item := range list {
			_, _ = redisClient.Del(item)
		}
	}
}
