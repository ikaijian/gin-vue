package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/handler"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/pkg/errno"
	"github.com/kaijian/gin-vue/utils"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("update function called.", lager.Data{"X-Request-Id": utils.GentReqID(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint(userId)

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
