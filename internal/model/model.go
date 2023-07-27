package model

import "gorm.io/gorm"

// 总结：如果需要查出所有关联的数据就用Preload，查一条关联数据用Related

var migrate = make([]any, 0, 10)

type Model struct {
	ID        int64          `json:"id" gorm:"primarykey" `
	CreatedAt int64          `json:"-"`
	UpdatedAt int64          `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetMigrate() []any {
	return migrate
}
func addMigrate(model ...any) {
	migrate = append(migrate, model...)
}
