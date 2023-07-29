package db

import (
	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/utils/tokens"
	"gopkg.in/hlandau/passlib.v1"
)

func Register(data model.User) (id int64, token string, msg string, err error) {
	hash, err := passlib.Hash(data.Pawd)
	if err != nil {
		msg = "请更换您的密码再试一次"
		return
	}
	data.Pawd = hash
	err = db.Create(&data).Error
	if err != nil {
		msg = "抱歉，请稍后再试..."
		return
	}
	token, err = tokens.GetToken(data.ID, data.Name)
	if err != nil {
		msg = "抱歉，麻烦再试一次吧..."
		return
	}
	return data.ID, token, "", nil
}

func Login(user, pawd string) (id int64, token string, msg string, err error) {
	var data model.User
	//根据用户名获取对应的全部数据
	err = db.Where("name=?", user).Find(&data).Error
	if err != nil {
		msg = "没有此用户名~"
		return
	}
	//进行哈希值效验密码是否正确
	newHash, err := passlib.Verify(pawd, data.Pawd)
	if err != nil {
		msg = "用户名或者密码不正确!"
		return
	}
	if newHash != "" {
		//登陆成功，判断是否需要更换哈希值
		db.Where(data).Update("pawd", newHash)
	}
	token, err = tokens.GetToken(data.ID, data.Name)
	if err != nil {
		msg = "抱歉，麻烦再试一次吧..."
		return
	}
	return data.ID, token, "", nil
}
