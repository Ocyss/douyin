package handlers

import (
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/server/common"
	"github.com/Ocyss/douyin/utils/checks"
	"github.com/gin-gonic/gin"
)

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	// TODO: 用户登陆接口
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	signature := c.Query("signature")
	if msg, ok := checks.Username(username, password); !ok {
		common.Err(c, msg)
		return
	}
	id, token, msg, err := db.Register(username, password, signature)
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
