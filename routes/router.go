package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/controller/sd"
	"github.com/kaijian/gin-vue/controller/user"
	"github.com/kaijian/gin-vue/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"the incorrect API route.")
	})

	//用户路由设置
	u := g.Group("/v1/user")
	{
		u.POST("",user.Create) //创建用户
		u.DELETE("/:id",user.Delete) //删除用户
		u.PUT("/:id",user.Update) //更新用户
		u.GET("",user.List) //用户列表
		u.GET("/:username",user.Get) //指定用户详情
	}

	svcd :=g.Group("/sd")
	{
		svcd.GET("/health",sd.HealthCheck)
		svcd.GET("/disk",sd.DiskCheck)
		svcd.GET("/cpu",sd.CPUCheck)
		svcd.GET("/ram",sd.RAMCheck)
	}

	return g
}
