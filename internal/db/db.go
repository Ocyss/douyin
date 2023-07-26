package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
	err := db.AutoMigrate()
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %s", err.Error())
	}
}

func GetDb() *gorm.DB {
	return db
}
