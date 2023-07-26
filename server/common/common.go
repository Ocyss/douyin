package common

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/gin-gonic/gin"
	"net/http"
)

const STATUS_OK = 0
const STATUS_ERR = 1

func OK(c *gin.Context) {
	OKData(c, nil)
}

func OKData(c *gin.Context, data gin.H) {
	res := gin.H{
		"status_code": STATUS_OK,
		"status_msg":  "Success",
	}
	for k, v := range data {
		res[k] = v
	}
	c.JSON(http.StatusOK, res)
}

func ErrCode(c *gin.Context, msg string, err ...error) {
	res := gin.H{
		"status_code": STATUS_ERR,
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

func Err(c *gin.Context, msg string, err ...error) {
	ErrCode(c, msg, err...)
}

func ErrParam(c *gin.Context, err ...error) {
	ErrCode(c, "参数不正确", err...)
}
