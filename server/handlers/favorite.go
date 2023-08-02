package handlers

import (
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/server/common"
	"github.com/Ocyss/douyin/utils/tokens"
	"github.com/gin-gonic/gin"
)

type actionReqs struct {
	Token      string `form:"token"   json:"token" binding:"required"`            // 用户鉴权token
	VideoId    int64  `form:"video_id"   json:"video_id" binding:"required"`      // 视频id
	ActionType int    `form:"action_type"  json:"action_type" binding:"required"` // 1-点赞，2-取消点赞
}

// FavoriteAction 点赞
func FavoriteAction(c *gin.Context) {
	var (
		reqs actionReqs
	)
	// 参数绑定
	if err := common.Bind(c, &reqs); err != nil {
		common.ErrParam(c, err)
		return
	}
	claims, err := tokens.CheckToken(reqs.Token)

	if err != nil {
		common.Err(c, "Token 错误", err)
		return
	}
	err = db.VideoLike(claims.ID, reqs.VideoId, reqs.ActionType)
	if err != nil {
		common.Err(c, "网卡了,再试一次吧", err)
	} else {
		common.OK(c)
	}
}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) {
	// TODO: 点赞列表接口
}
