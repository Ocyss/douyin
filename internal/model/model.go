package model

import "gorm.io/gorm"

type Model struct {
	ID        int64          `json:"id" gorm:"primarykey" `
	CreatedAt int64          `json:"-"`
	UpdatedAt int64          `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
