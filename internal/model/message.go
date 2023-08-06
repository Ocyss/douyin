package model

import (
	"time"

	"github.com/Ocyss/douyin/utils"
	"gorm.io/gorm"
)

type (
	// Message 消息表
	Message struct {
		ID         int64  `json:"id" gorm:"primarykey;comment:主键"`
		CreatedAt  int64  `json:"created_at" gorm:"autoUpdateTime:nano"`
		ToUserID   int64  `json:"to_user_id" gorm:"primaryKey;comment:该消息接收者的id"`
		FromUserID int64  `json:"from_user_id" gorm:"primaryKey;comment:该消息发送者的id"`
		ToUser     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		FromUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content    string `json:"content" gorm:"comment:消息内容"`
		CreateTime string `json:"create_time" gorm:"comment:消息创建时间"` // hook生成,省的每次查询都生成,格式yyyy-MM-dd HH:MM:ss
	}
)

func (u *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(3, 114514)
	}
	u.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

func init() {
	addMigrate(&Message{})
}
