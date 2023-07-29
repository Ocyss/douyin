package handlers

import (
	"errors"
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/server/common"
	"github.com/Ocyss/douyin/utils"
	"github.com/Ocyss/douyin/utils/checks"
	"github.com/gin-gonic/gin"
)

type userReqs struct {
	Name            string `json:"username" form:"username" binding:"required"` // 用户名称
	Pawd            string `json:"password" form:"password" binding:"required"` // 用户密码
	Avatar          string `json:"avatar" form:"avatar"`                        // 用户头像
	BackgroundImage string `json:"background_image" form:"background_image"`    // 用户个人页顶部大图
	Signature       string `json:"signature" form:"signature"`                  // 个人简介
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	var (
		reqs userReqs
	)
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		if err2 := c.ShouldBindJSON(&reqs); err2 != nil {
			common.ErrParam(c, errors.Join(err, err2))
			return
		}
	}
	if msg := checks.ValidateInput(4, 32, reqs.Name, reqs.Pawd); len(msg) > 0 {
		common.Err(c, "账户或者密码"+msg)
		return
	}

	id, token, msg, err := db.Login(reqs.Name, reqs.Pawd)
	if err != nil {
		common.Err(c, msg, err)
	} else {
		common.OK(c, H{"user_id": id, "token": token})
	}
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var (
		reqs userReqs
		data model.User
	)
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		if err2 := c.ShouldBindJSON(&reqs); err2 != nil {
			common.ErrParam(c, errors.Join(err, err2))
			return
		}
	}
	if msg := checks.ValidateInput(4, 32, reqs.Name, reqs.Pawd); len(msg) > 0 {
		common.Err(c, "账户或者密码"+msg)
		return
	}
	_ = utils.Merge(&data, reqs)
	id, token, msg, err := db.Register(data)
	if err != nil {
		common.Err(c, msg, err)
	} else {
		common.OK(c, H{"user_id": id, "token": token})
	}
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	// TODO: 用户信息接口
}
