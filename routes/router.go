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

	u := g.Group("/v1/user")
	{
		u.POST("/:username",user.Create)
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
