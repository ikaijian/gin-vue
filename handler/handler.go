package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijian/gin-vue/pkg/errno"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func SendResponse(c *gin.Context,err error,data interface{})  {
	code,message := errno.DecodeErr(err)

	c.JSON(http.StatusOK,Response{
		Code: code,
		Data: data,
		Message: message,
	})
}