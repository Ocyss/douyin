package db

import (
	"github.com/Ocyss/douyin/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
	err := db.AutoMigrate(&model.Comment{}, &model.User{}, &model.FriendUser{}, &model.Video{}, &model.Message{})
	if err != nil {
		log.Fatalf("数据库自动迁移失败: %s", err.Error())
	}
}

func GetDb() *gorm.DB {
	return db
}
