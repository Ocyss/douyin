package server

import (
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init(r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20
	r.Use(gin.LoggerWithWriter(log.StandardLogger().Out), gin.RecoveryWithWriter(log.StandardLogger().Out))
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Any("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router := r.Group("douyin")
	router.GET("feed", handlers.FeedGet)
}
