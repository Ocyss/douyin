package model

type (
	// Message 消息表
	Message struct {
		Model
		ToUserID   int64  `json:"to_user_id" gorm:"primaryKey;comment:该消息接收者的id"`
		FromUserID int64  `json:"from_user_id" gorm:"primaryKey;comment:该消息发送者的id"`
		ToUser     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		FromUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content    string `json:"content" gorm:"comment:消息内容"`
		//CreateTime string `json:"create_time" gorm:"comment:消息创建时间"`
	}
)

func init() {
	addMigrate(&Message{})
}
