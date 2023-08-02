package server

import (
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/Ocyss/douyin/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func Init(r *gin.Engine) {
	r.MaxMultipartMemory = 16 << 20                // 16 MiB
	r.Use(middleware.Logger(log.StandardLogger())) // 使用logrus记录日志
	r.Use(gin.Recovery())                          // 恐慌恢复
	r.Use(cors.Default())                          // 跨域处理
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router := r.Group("douyin")
	tester := r.Group("douyin")
	tester.Use(middleware.Test())
	// 视频类接口
	{
		newRouter(router, "feed/", handlers.VideoGet).GET()                     // 获取视频流
		newRouter(router, "publish/action/", handlers.VideoAction).POST()       // 视频投稿
		newRouter(tester, "publish/actionUrl/", handlers.VideoActionUrl).POST() // 视频投稿(测试接口)
		newRouter(router, "publish/list/", handlers.VideoList).GET()            // 获取发布列表
	}
	// 用户类接口
	{
		newRouter(router, "user/register/", handlers.UserRegister).POST() // 用户注册
		newRouter(router, "user/login/", handlers.UserLogin).POST()       // 用户登录
		newRouter(router, "user/", handlers.UserInfo).GET()               // 获取用户信息
	}
	// 互动类接口
	{
		newRouter(router, "favorite/action/", handlers.FavoriteAction).POST() // 点赞操作
		newRouter(router, "favorite/list/", handlers.FavoriteList).GET()      // 获取喜欢列表
		newRouter(router, "comment/action/", nil).POST()                      // 评论操作
		newRouter(router, "comment/list/", nil).GET()                         // 获取评论列表
	}
	//社交类接口
	{
		newRouter(router, "relation/action/", nil).POST()       // 关注/取关 操作
		newRouter(router, "relatioin/follow/list/", nil).GET()  // 获取用户关注列表
		newRouter(router, "relation/follower/list/", nil).GET() // 获取用户粉丝列表
		newRouter(router, "relation/friend/list/", nil).GET()   // 获取用户好友列表
		// 消息类接口
		{
			newRouter(router, "message/chat/", nil).GET()    // 获取消息
			newRouter(router, "message/action/", nil).POST() // 发送消息
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
