package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/handler"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)

	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)

}
