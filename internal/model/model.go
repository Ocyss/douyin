package model

import "gorm.io/gorm"

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey" `
	CreatedAt int            `json:"created_at"`
	UpdatedAt int            `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
