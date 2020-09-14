package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kaijian/gin-vue/model"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("db.driverName")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	database := viper.GetString("db.name")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	chartset := viper.GetString("db.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, chartset)

	db, err := gorm.Open(driverName, args)
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
