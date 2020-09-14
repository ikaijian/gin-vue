package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kaijian/gin-vue/common"
	"github.com/kaijian/gin-vue/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config","c","","apiserver config file path.")
)

func main() {
	pflag.Parse()

	//初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 设置gin运行模式.
	gin.SetMode(viper.GetString("server.runmode"))

	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}

