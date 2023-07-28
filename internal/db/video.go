package db

import (
	"github.com/Ocyss/douyin/internal/model"
)

func Feed(latestTime string) ([]model.Video, error) {
	var data []model.Video
	if len(latestTime) != 19 {
		latestTime = "9223372036854775806"
	}
	//t := time.Unix(0, latestTime*int64(time.Millisecond))
	err := db.Where("id < ?", latestTime).Order("id DESC").Limit(5).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Action 视频投稿
func Action(id int64, file []byte, url, title string) (int64, string, error) {
	var data = model.Video{
		AuthorID: id,
		PlayUrl:  url,
		Title:    title,
	}
	if file != nil {
		// TODO: file数据上传
	}
	err := db.Create(&data).Error
	if err != nil {
		return 0, "", err
	}
	return data.ID, "", nil
}
