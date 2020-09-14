package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kaijian/gin-vue/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	drivierName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	chartset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, chartset)

	db, err := gorm.Open(drivierName, args)
	if err != nil {
		panic("failed to connect database , err" + err.Error())
	}

	//自动创建表
	db.AutoMigrate(&model.User{})

	DB =db
	return db
}

func GetDB() *gorm.DB  {
	return DB
}
