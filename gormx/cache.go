package gormx

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
)

type GormCache struct {
	Cache cache.Cache
	Db    *gorm.DB
	Conf  cache.CacheConf
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("gormCache")
)

func NewCacheTool(conf cache.CacheConf, db *gorm.DB) *GormCache {
	return &GormCache{
		cache.New(conf, singleFlights, stats, gorm.ErrRecordNotFound),
		db,
		conf,
	}
}
