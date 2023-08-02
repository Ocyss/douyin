package server

import (
	"github.com/Ocyss/douyin/cmd/flags"
	"github.com/Ocyss/douyin/server/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyHandler func(*gin.Context) (int, any)

// handler 装饰器
func handler() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(c *gin.Context) {
			code, data := h(c)
			req := gin.H{
				"status_code": code,
				"status_msg":  "",
			}
			if code == 0 {
				//判断数据类型
				if val, ok := data.(handlers.H); ok {
					for k, v := range val {
						req[k] = v
					}
				}
				req["status_msg"] = "ok!"
				c.JSON(200, req)
			} else {
				switch data.(type) {
				case string:
					req["status_msg"] = data
				case error:
					//判断是否debug模式，是的话返回错误信息
					if flags.Dev || flags.Debug {
						req["errmsg"] = data.(error).Error()
					}
				case handlers.MyErr:
					e := data.(handlers.MyErr)
					req["status_msg"] = e.Msg
					//判断是否debug模式，是的话返回错误信息
					if flags.Dev || flags.Debug {
						errs := make([]string, 0, 10)
						for i := range e.Errs {
							errs = append(errs, e.Errs[i].Error())
						}
						req["errmsg"] = errs
					}
				}

				c.JSON(http.StatusOK, req)
			}
		}
	}
}

type myrouter struct {
	Group    *gin.RouterGroup
	Path     string
	Handler  MyHandler
	Handlers []gin.HandlerFunc
}

func newRouter(group *gin.RouterGroup, path string, handler MyHandler, handlers ...gin.HandlerFunc) *myrouter {
	return &myrouter{
		group,
		path,
		handler,
		handlers}
}

func (r *myrouter) Handle(method string) *myrouter {
	if r.Handler == nil {
		// 防止空指针 gin报错
		return r
	}
	r.Group.Handle(method, r.Path, append(r.Handlers, handler()(r.Handler))...)
	return r
}
func (r *myrouter) GET() *myrouter {
	r.Handle("GET")
	return r
}

func (r *myrouter) POST() *myrouter {
	r.Handle("POST")
	return r
}

func (r *myrouter) PUT() *myrouter {
	r.Handle("PUT")
	return r
}

func (r *myrouter) DELETE() *myrouter {
	r.Handle("DELETE")
	return r
}
