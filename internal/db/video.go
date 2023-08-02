package db

import (
	"errors"
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

func VideoLike(uid, vid int64, _type int) error {
	var err error
	association := db.Model(&model.User{Model: id(uid)}).Association("Favorite")
	switch _type {
	case 1:
		err = association.Append(&model.Video{Model: id(vid)})
	case 2:
		err = association.Delete(&model.Video{Model: id(vid)})
	default:
		err = errors.New("你看看你传的什么东西吧")
	}
	if err != nil {
		return err
	}
	return nil
}

func VideoLikeList(uid int64) ([]*model.Video, error) {
	var data []*model.Video
	err := db.Model(&model.User{Model: id(uid)}).Association("Favorite").Find(&data)
	if err != nil {
		return nil, err
	}
	for i := range data {
		err = db.Preload("Author").Find(data[i]).Error
		if err != nil {
			return nil, err
		}
		data[i].IsFavorite = true
	}
	return data, nil
}
