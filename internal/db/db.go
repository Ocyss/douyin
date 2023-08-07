package db

import (
	"context"
	"reflect"
	"time"

	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/internal/model"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

// InitDb 初始化数据库服务
func InitDb(d *gorm.DB) {
	db = d
	for _, m := range model.GetMigrate() {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Fatalf("%s 模型自动迁移失败: %s", reflect.TypeOf(m), err.Error())
		}
	}
	err := db.SetupJoinTable(&model.Video{}, "CoAuthor", &model.UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,Video: %s", err)
	}
	err = db.SetupJoinTable(&model.User{}, "Videos", &model.UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,User: %s", err)
	}
}

// InitRdb 初始化 Redis
func InitRdb(r *redis.Client) {
	rdb = r
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("连接redis出错，错误信息：%v", err)
	}
	// 内存模式下清空 Redis
	if flags.Memory {
		rdb.FlushAll(ctx)
	}
}

// id 快捷用法返回一个Model{id:val}
func id(val int64) model.Model {
	return model.Model{ID: val}
}

func GetDb() *gorm.DB {
	return db
}

func GetRdb() *redis.Client {
	return rdb
}
