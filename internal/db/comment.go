package db

import "github.com/Ocyss/douyin/internal/model"

func CommentPush(uid, vid int64, content string) (*model.Comment, error) {
	data := model.Comment{UserID: uid, VideoID: vid, Content: content}
	err := db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func CommentDel(cid int64) error {
	return db.Delete(&model.Comment{}, cid).Error
}
