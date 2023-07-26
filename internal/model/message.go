package model

type Message struct {
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	ToUser     User   `json:"-" gorm:"comment:该消息接收者的id"`
	FromUser   User   `json:"-" gorm:"comment:该消息发送者的id"`
	Content    string `json:"content" gorm:"comment:消息内容"`
	//CreateTime string `json:"create_time" gorm:"comment:消息创建时间"`
}
