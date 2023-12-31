package middleware

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/gin-gonic/gin"
)

// Test 测试中间件,仅开发模式下使用
func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		if flags.Dev || flags.Tst {
			c.Next()
		} else {
			c.Abort()
		}
	}
}
