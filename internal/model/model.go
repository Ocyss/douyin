package model

import (
	"gorm.io/gorm"
	"time"
)

// 总结：如果需要查出所有关联的数据就用Preload，查一条关联数据用Related

var migrate = make([]any, 0, 10)

type Model struct {
	ID        int64          `json:"id" gorm:"primarykey;comment:主键"`
	CreatedAt time.Time      `json:"-" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"-" gorm:"comment:修改时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"comment:删除时间"`
}

func GetMigrate() []any {
	return migrate
}
func addMigrate(model ...any) {
	migrate = append(migrate, model...)
}
