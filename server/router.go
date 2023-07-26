package server

import (
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/Ocyss/douyin/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init(r *gin.Engine) {
	r.MaxMultipartMemory = 16 << 20                // 16 MiB
	r.Use(middleware.Logger(log.StandardLogger())) // 使用logrus记录日志
	r.Use(gin.Recovery())                          // 恐慌恢复
	r.Use(cors.Default())                          // 跨域处理
	r.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router := r.Group("douyin")
	router.GET("feed", handlers.FeedGet)
}
