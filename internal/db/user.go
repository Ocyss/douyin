package db

import (
	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/utils/tokens"
	"gopkg.in/hlandau/passlib.v1"
)

func Register(username, password, signature string) (int64, string, string, error) {
	hash, err := passlib.Hash(password)
	if err != nil {
		return 0, "", "请更换您的密码再试一次", err
	}
	data := model.User{Name: username, Pawd: hash, Signature: signature}
	err = db.Create(&data).Error
	if err != nil {
		return 0, "", "抱歉，请稍后再试...", err
	}
	token, err := tokens.GetToken(data.ID, username)
	if err != nil {
		return 0, "", "抱歉，麻烦再试一次吧...", err
	}
	return data.ID, token, "", nil
}
