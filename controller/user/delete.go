package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/handler"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/pkg/errno"
	"strconv"
)

func Delete(c *gin.Context)  {
	userId ,_ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint(userId));err !=nil{
		handler.SendResponse(c,errno.ErrDatabase,nil)
		return
	}

	handler.SendResponse(c,nil,nil)
}
