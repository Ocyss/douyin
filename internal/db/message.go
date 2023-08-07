package db

import "github.com/Ocyss/douyin/internal/model"

// MessagePush 发送消息
func MessagePush(fid, tid int64, content string) error {
	data := model.Message{
		ToUserID:   tid,
		FromUserID: fid,
		Content:    content,
	}
	return db.Create(&data).Error
}

// MessageGet 获取消息列表
func MessageGet(fid, tid, preTime int64) ([]*model.Message, error) {
	var data []*model.Message
	err := db.Where(model.Message{ToUserID: tid, FromUserID: fid}).Order("created_at DESC").Find(&data).Error
	return data, err
}
