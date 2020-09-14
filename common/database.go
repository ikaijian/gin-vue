package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kaijian/gin-vue/model"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	chartset := viper.GetString("datasource.charset")
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
