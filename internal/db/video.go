package db

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/Ocyss/douyin/utils"

	"gorm.io/gorm"

	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/utils/upload"
)

// Feed 获取视频流
func Feed(uid int64, ip, ty string, repeat bool) ([]model.Video, error) {
	var data []model.Video
	res := make([]model.Video, 0, 10)
	// 循环20次,随机生成20个主键id,通过IP来减少重复推送
	for batch := 0; len(res) < 3 && batch < 6; batch++ {
		var rv []int64
		if v, ok := videoAll[ty]; ok {
			rv = utils.RandVid(v, 20)
		}
		db.Set("user_id", uid).Where(rv).Find(&data)
		for i := 0; i < len(data) && len(res) < 3; i++ {
			if data[i].ViewedFilter(ip) || repeat {
				data[i].PlayCount++
				res = append(res, data[i])
			}
		}
	}
	return res, nil
}

// VideoUpload 视频投稿
func VideoUpload(uid int64, file multipart.File, PlayUrl, CoverUrl, title, typeOf string, UserCreations []*model.UserCreation) (int64, string, error) {
	if typeOf == "" {
		typeOf = "t"
	}
	data := model.Video{
		AuthorID: uid,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
		Title:    title,
		TypeOf:   typeOf,
	}
	// 开启事务,上传失败不添加数据
	tx := db.Begin()
	err := tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}
	if file != nil {
		// reader := bytes.NewReader(file)
		fname := fmt.Sprintf("t/%d.mp4", data.ID)
		url, err := upload.Aliyun(fname, file)
		if err != nil {
			tx.Rollback()
			return 0, "上传出错...", err
		}
		data.PlayUrl = url + fname
		data.CoverUrl = url + fmt.Sprintf("t/%d.jpg", data.ID)
		err = tx.Save(&data).Error
		if err != nil {
			tx.Rollback()
			return 0, "更新出错...", err
		}
	}
	UserCreations = append([]*model.UserCreation{{VideoID: data.ID, UserID: uid, Type: "Up主"}}, UserCreations...)
	for _, uc := range UserCreations {
		uc.VideoID = data.ID
		err := tx.Create(uc).Error
		if err != nil {
			tx.Rollback()
			return 0, "创建出错...", err
		}
		tx.Model(&model.User{Model: id(uc.UserID)}).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1))
	}
	tx.Commit()
	videoAll["all"] = append(videoAll["all"], data.ID)
	videoAll["t"] = append(videoAll["t"], data.ID)
	return data.ID, "", nil
}

// VideoLike 视频点赞操作
func VideoLike(uid, vid int64, _type int) error {
	var err error
	// association := db.Model(&model.User{Model: id(uid)}).Omit("Favorite").Association("Favorite")
	val := &model.Video{Model: id(vid)}
	switch _type {
	case 1:
		row := db.Exec("INSERT INTO `user_favorite` (`user_id`,`video_id`) VALUES (?,?)", uid, vid)
		if row.Error == nil && row.RowsAffected == 1 {
			val.HIncrByFavoriteCount(1)
		} else {
			err = errors.Join(row.Error, errors.New("err:可能已有数据"))
		}
	case 2:
		row := db.Exec("DELETE FROM user_favorite Where user_id = ? AND video_id = ?", uid, vid)
		if row.Error == nil && row.RowsAffected == 1 {
			val.HIncrByFavoriteCount(-1)
		} else {
			err = errors.Join(row.Error, errors.New("err:可能无该数据"))
		}
	default:
		err = errors.New("你看看你传的什么东西吧")
	}
	if err != nil {
		return err
	}
	return nil
}

// VideoLikeList 获取喜欢列表
func VideoLikeList(uid int64) ([]*model.Video, error) {
	var data []*model.Video
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Favorite").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// VideoList 获取作品列表
func VideoList(uid int64) ([]*model.Video, error) {
	var data []*model.Video
	err := db.Set("user_id", uid).Model(&model.User{Model: id(uid)}).Association("Videos").Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
