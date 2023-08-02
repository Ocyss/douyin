package handlers

import (
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/server/common"
	"github.com/Ocyss/douyin/utils"
	"github.com/Ocyss/douyin/utils/checks"
	"github.com/Ocyss/douyin/utils/tokens"
	"github.com/gin-gonic/gin"
)

type (
	userLogRegReqs struct {
		Name            string `json:"username" form:"username" binding:"required"` // 用户名称
		Pawd            string `json:"password" form:"password" binding:"required"` // 用户密码
		Avatar          string `json:"avatar" form:"avatar"`                        // 用户头像
		BackgroundImage string `json:"background_image" form:"background_image"`    // 用户个人页顶部大图
		Signature       string `json:"signature" form:"signature"`                  // 个人简介
	}
	userReqs struct {
		ID    int64  `json:"user_id" form:"user_id" binding:"required"` // 用户id
		Token string `json:"token" form:"token" binding:"required"`     // 用户鉴权token
	}
	userInfoResp struct {
		ID              int64  `json:"user_id"`          // 用户id
		Name            string `json:"name"`             // 用户名称
		FollowCount     int64  `json:"follow_count"`     // 关注总数
		FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
		IsFollow        bool   `json:"is_follow"`        // 是否关注
		Avatar          string `json:"avatar"`           // 用户头像
		BackgroundImage string `json:"background_image"` //用户个人页顶部大图
		Signature       string `json:"signature"`        //个人简介
		WorkCount       int64  `json:"work_count"`       // 作品数量
		TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
		FavoriteCount   int64  `json:"favorite_count"`   // 点赞数量
	}
)

// UserLogin 用户登陆
func UserLogin(c *gin.Context) (int, any) {
	var (
		reqs userLogRegReqs
	)
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return fail, ErrParam(err)
	}
	if msg := checks.ValidateInput(4, 32, reqs.Name, reqs.Pawd); len(msg) > 0 {
		return fail, Err("账户或者密码" + msg)
	}

	data, msg, err := db.Login(reqs.Name, reqs.Pawd)
	if err != nil {
		return fail, Err(msg, err)
	}
	token, err := tokens.GetToken(data.ID, data.Name)
	if err != nil {
		return fail, Err("抱歉，麻烦再试一次吧...", err)
	}
	return ok, H{"user_id": data.ID, "token": token}
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) (int, any) {
	var (
		reqs userLogRegReqs
		data model.User
	)
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		return fail, ErrParam(err)
	}
	if msg := checks.ValidateInput(4, 32, reqs.Name, reqs.Pawd); len(msg) > 0 {
		return fail, Err("账户或者密码" + msg)
	}
	_ = utils.Merge(&data, reqs)

	msg, err := db.Register(&data)
	if err != nil {
		return fail, Err(msg, err)
	}

	token, err := tokens.GetToken(data.ID, data.Name)
	if err != nil {
		return fail, Err("抱歉，麻烦再试一次吧...", err)
	}
	return ok, H{"user_id": data.ID, "token": token}
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) (int, any) {
	var (
		reqs userReqs
		resp userInfoResp
	)
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		return fail, ErrParam(err)
	}

	_, err := tokens.CheckToken(reqs.Token)

	if err != nil {
		return fail, Err("Token 错误", err)
	}
	data, msg, err := db.UserInfo(reqs.ID)
	if err != nil {
		return fail, Err(msg, err)
	}
	_ = utils.Merge(&resp, data)
	return ok, H{"user": resp}
}
