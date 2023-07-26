package handlers

import (
	"github.com/Ocyss/douyin/server/common"
	"github.com/gin-gonic/gin"
)

// FeedGet 视频流获取
func FeedGet(c *gin.Context) {
	// TODO: 视频流获取接口
	var data gin.H
	var err error
	if err == nil {
		common.Err(c, "Err...")
	} else {
		common.OKData(c, data)
	}
}

// FeedAction 视频投稿
func FeedAction(c *gin.Context) {
	// TODO: 视频投稿接口
}

// FeedList 发布列表
func FeedList(c *gin.Context) {
	// TODO: 发布列表接口
}
