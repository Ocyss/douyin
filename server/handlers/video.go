package handlers

import (
	"github.com/Ocyss/douyin/internal/db"
	"github.com/Ocyss/douyin/internal/model"
	"github.com/Ocyss/douyin/utils/tokens"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type (
	actionData struct {
		Data  multipart.File `json:"data" form:"data"`
		Token string         `json:"token" form:"token"`
		Title string         `json:"title" form:"title"`
		Url   string         `json:"url" form:"url"`
		ID    int64          `json:"id" form:"id"`
	}
)

// VideoGet 视频流获取
func VideoGet(c *gin.Context) (int, any) {
	//token := c.Query("token")
	t := c.Query("latest_time")
	data, err := db.Feed(t)
	if err != nil {
		return fail, Err("数据获取出错，请稍后再试.", err)
	}
	res := H{
		"video_list": data,
	}
	if len(data) > 0 {
		res["next_time"] = data[len(data)-1].ID
	}
	return ok, res
}

// VideoAction 视频投稿
func VideoAction(c *gin.Context) (int, any) {
	var data actionData
	file, _, err := c.Request.FormFile("data")
	data.Data = file
	data.Token = c.PostForm("token")
	data.Title = c.PostForm("title")
	if err != nil || data.Token == "" {
		return fail, ErrParam(err)
	}
	token, err := tokens.CheckToken(data.Token)
	if err != nil {
		return fail, Err("Token 错误", err)
	}
	id, msg, err := db.VideoUpload(token.ID, data.Data, "", data.Title)
	if err != nil {
		return fail, Err(msg, err)
	}
	return ok, H{"vid": id}
}

// VideoActionUrl 视频投稿
// 测试接口可直接指定URL，或使用ID进行投稿
func VideoActionUrl(c *gin.Context) (int, any) {
	var data actionData

	err := c.ShouldBindJSON(&data)
	if err != nil || (data.ID == 0 && data.Token == "") || (data.Data == nil && data.Url == "") {
		return fail, ErrParam(err)
	}
	if data.Token != "" {
		token, err := tokens.CheckToken(data.Token)
		if err != nil {
			return fail, Err("Token 错误", err)
		}
		data.ID = token.ID
	}
	id, msg, err := db.VideoUpload(data.ID, data.Data, data.Url, data.Title)
	if err != nil {
		return fail, Err(msg, err)
	}
	return ok, H{"vid": id}
}

// VideoList 发布列表
func VideoList(c *gin.Context) (int, any) {
	var (
		data []*model.Video
		reqs userReqs
	)
	// 参数绑定
	if err := c.ShouldBindQuery(&reqs); err != nil {
		return fail, ErrParam(err)
	}
	_, err := tokens.CheckToken(reqs.Token)

	if err != nil {

		return fail, Err("Token 错误", err)
	}

	data, err = db.VideoList(reqs.ID)
	if err != nil {
		return fail, Err("网卡了,再试一次吧", err)
	}

	return ok, H{"video_list": data}
}
