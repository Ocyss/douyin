package db

import (
	"github.com/Ocyss/douyin/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
	for _, m := range model.GetMigrate() {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Fatalf("%s 模型自动迁移失败: %s", reflect.TypeOf(m), err.Error())
		}
	}
}

func GetDb() *gorm.DB {
	return db
}
