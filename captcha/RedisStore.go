package captchaTool

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisStore struct {
	RedisClient *redis.Redis
}

func NewRedisStore(conf redis.RedisConf) *RedisStore {
	return &RedisStore{
		RedisClient: redis.MustNewRedis(conf),
	}
}

func captchaKey(id string) string {
	return fmt.Sprintf("captcha:%s", id)
}

func (r *RedisStore) Set(id string, value string) error {
	key := captchaKey(id)
	return r.RedisClient.Setex(key, value, 600)
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	key := captchaKey(id)
	exists, err := r.RedisClient.Exists(key)
	if err != nil {
		return false
	}
	if exists {
		if info, err := r.RedisClient.Get(key); err == nil {
			if info == answer {
				if clear {
					_, _ = r.RedisClient.Del(key)
				}
				return true
			}
		}
	}
	return false
}

func (r *RedisStore) Get(id string, clear bool) string {
	key := captchaKey(id)
	exists, err := r.RedisClient.Exists(key)
	if err != nil {
		return ""
	}
	if exists {
		if info, err := r.RedisClient.Get(key); err == nil {
			if clear {
				_, _ = r.RedisClient.Del(key)
			}
			return info
		}
	}
	return ""
}
