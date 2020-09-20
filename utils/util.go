package utils

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"github.com/teris-io/shortid"
)

//返回随机字符串
func RandomString(number int) string {

	var letters = []byte("abcdeASDFASDFfghiDSFDASEREGFGFjklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, number)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//
func GenShortId() (string, error) {
	return shortid.Generate()
}

func GentReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request_Id")
	if !ok {
		return ""
	}

	if requestId,ok := v.(string); ok{
		return requestId
	}
	return ""
}
