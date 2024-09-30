package gormV2

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	//*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primarykey" json:"id"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
	UpdatedAt string `gorm:"updated_at" json:"updated_at"`
}

type BaseDel struct {
	//*gorm.DB  `gorm:"-" json:"-"`
	Id        int64                 `gorm:"primarykey" json:"id"`
	CreatedAt string                `gorm:"created_at" json:"created_at"`
	UpdatedAt string                `gorm:"updated_at" json:"updated_at"`
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
