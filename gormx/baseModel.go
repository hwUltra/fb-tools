package gormx

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Base struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`
}

type BaseDel struct {
	Id        int64                 `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time             `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time             `gorm:"updated_at" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `gorm:"delete_at" json:"-"`
}

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
