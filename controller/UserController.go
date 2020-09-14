package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kaijian/gin-vue/common"
	"github.com/kaijian/gin-vue/dto"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/response"
	"github.com/kaijian/gin-vue/utils"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}

	//	给名称一个随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}

	hassedPassword , err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err !=nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}
	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hassedPassword),
	}
	DB.Create(&newUser)
	//返回参数
	response.Success(c,nil,"用户注册成功")
}

func Login(c *gin.Context)  {
	DB := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据校验
	//数据验证
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}
	//判断密码是否正确
	if err :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err !=nil{
		response.Response(c,http.StatusBadRequest,400,nil,"密码错误")
		return
	}
	//发放token
	token,err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token generate error:%v",err)
		return
	}
	//返回结果
	response.Success(c,gin.H{"token":token},"登陆成功")
}

func User(c *gin.Context)  {
	user ,_ := c.Get("user")
	response.Success(c,gin.H{"user":dto.ToUserDto(user.(model.User))},"用户信息")
}


func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 {
		return true
	}
	return false
}
