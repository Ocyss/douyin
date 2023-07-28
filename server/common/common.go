package common

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

const statusOk = 0
const statusErr = 1

//func OK(c *gin.Context) {
//	OKData(c, nil)
//}

func OK(c *gin.Context, data ...handlers.H) {
	res := gin.H{
		"status_code": statusOk,
		"status_msg":  "Success",
	}
	for d := range data {
		for k, v := range data[d] {
			res[k] = v
		}
	}
	c.JSON(http.StatusOK, res)
}

func Err(c *gin.Context, msg string, err ...error) {
	res := gin.H{
		"status_code": statusErr,
		"status_msg":  msg,
	}
	// 调试与开发模式，返回错误消息。
	if (flags.Dev || flags.Debug) && len(err) > 0 {
		errs := make([]string, len(err))
		for i, e := range err {
			errs[i] = e.Error()
		}
		res["errmsg"] = errs
	}
	c.JSON(http.StatusOK, res)
}

func ErrParam(c *gin.Context, err ...error) {
	Err(c, "参数不正确", err...)
}
