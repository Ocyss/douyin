package handlers

import (
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/server/common"
	"github.com/Ocyss/douyin/utils/tokens"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type (
	actionData struct {
		Data  []byte `json:"data"`
		Token string `json:"token"`
		Title string `json:"title"`
		Url   string `json:"url"`
		ID    int64  `json:"id"`
	}
)

// VideoGet 视频流获取
func VideoGet(c *gin.Context) {
	var latestTime int64
	//token := c.Query("token")
	if t := c.Query("latest_time"); t != "" {
		latestTime, _ = strconv.ParseInt(t, 10, 0)
	}
	if latestTime == 0 {
		latestTime = time.Now().UnixMilli()
	}
	data, err := db.Feed(latestTime)
	if err != nil {
		common.Err(c, "数据获取出错，请稍后再试.", err)
	} else {
		res := H{
			"video_list": data,
		}
		if len(data) > 0 {
			res["next_time"] = data[len(data)-1].CreatedAt.Unix()
		}
		common.OK(c, res)
	}
}

// VideoAction 视频投稿
func VideoAction(c *gin.Context) {
	// TODO: 视频投稿接口
}

// VideoActionUrl 视频投稿
// 测试接口可直接指定URL，或使用ID进行投稿
func VideoActionUrl(c *gin.Context) {
	var data actionData
	err := c.Bind(&data)
	if err != nil || (data.ID == 0 && data.Token == "") || (len(data.Data) == 0 && data.Url == "") {
		common.ErrParam(c, err)
		return
	}
	if data.Token != "" {
		token, err := tokens.CheckToken(data.Token)
		if err != nil {
			common.Err(c, "Token 错误", err)
			return
		}
		data.ID = token.ID
	}
	id, msg, err := db.Action(data.ID, data.Data, data.Url, data.Title)
	if err != nil {
		common.Err(c, msg, err)
	} else {
		common.OK(c, H{"vid": id})
	}
}

// VideoList 发布列表
func VideoList(c *gin.Context) {
	// TODO: 发布列表接口
}
