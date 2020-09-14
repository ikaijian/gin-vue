package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/controller"
	"github.com/kaijian/gin-vue/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine  {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/user",middleware.AuthMiddleware(),controller.User)
	return  r
}
