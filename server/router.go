package server

import (
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/Ocyss/douyin/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
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
	tester := r.Group("douyin")
	tester.Use(middleware.Test())
	// 视频类接口
	{
		router.GET("feed/", handlers.VideoGet)                     // 获取视频流
		router.POST("publish/action")                              // 视频投稿
		tester.POST("publish/actionUrl/", handlers.VideoActionUrl) // 视频投稿(测试接口)
		//router.GET("publish/list")                                // 获取发布列表
	}
	// 用户类接口
	{
		router.POST("user/register/", handlers.UserRegister) // 用户注册
		router.POST("user/login/", handlers.UserLogin)       // 用户登录
		router.GET("user/", handlers.UserInfo)               // 获取用户信息
	}
	// 互动类接口
	{
		router.POST("favorite/action/", handlers.FavoriteAction) // 点赞操作
		//router.GET("favorite/list")    // 获取喜欢列表
		router.POST("comment/action/") // 评论操作
		router.GET("comment/list/")    // 获取评论列表
	}
	//社交类接口
	{
		router.POST("relation/action")       // 关注/取关 操作
		router.GET("relatioin/follow/list")  // 获取用户关注列表
		router.GET("relation/follower/list") // 获取用户粉丝列表
		router.GET("relation/friend/list")   // 获取用户好友列表
		// 消息类接口
		{
			router.GET("message/chat")    // 获取消息
			router.POST("message/action") // 发送消息
		}
	}
	// 挂载 web 服务
	r.Use(static.Serve("/", static.LocalFile("web", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("web/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})
}
