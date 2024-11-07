package captchaTool

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type CacheStore struct {
	Cache cache.Cache
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("CaptchaCache")
	cacheErr      = errors.New("not found")
)

func NewCacheStore(conf cache.CacheConf) *CacheStore {
	cc := cache.New(conf, singleFlights, stats, cacheErr)
	return &CacheStore{
		cc,
	}
}

func captchaKey(id string) string {
	return fmt.Sprintf("cache:captcha:%s", id)
}

func (cs *CacheStore) Set(id string, value string) error {
	return cs.Cache.SetWithExpire(captchaKey(id), value, 600)
}

func (cs *CacheStore) Get(id string, clear bool) (res string) {
	key := captchaKey(id)
	if err := cs.Cache.Get(key, &res); err != nil {
		return ""
	}
	if clear == true {
		_ = cs.Cache.Del(key)
	}
	return res
}

func (cs *CacheStore) Verify(id, answer string, clear bool) bool {
	key := captchaKey(id)
	res := ""
	if err := cs.Cache.Get(key, &res); err != nil {
		return false
	}
	if res == answer {
		if clear == true {
			_ = cs.Cache.Del(key)
		}
		return true
	}
	return false
}
