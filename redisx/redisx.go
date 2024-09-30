package redisx

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func DelPrefix(redisClient *redis.Redis, prefix string) {
	if list, _, err := redisClient.Scan(0, prefix, 0); err == nil {
		for _, item := range list {
			_, _ = redisClient.Del(item)
		}
	}
}
