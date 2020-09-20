package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/handler"
	"github.com/kaijian/gin-vue/pkg/errno"
	"github.com/kaijian/gin-vue/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
