package middleware

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		if flags.Dev {
			c.Next()
		} else {
			c.Abort()
		}
	}
}
