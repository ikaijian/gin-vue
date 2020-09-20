package main

import (
	"github.com/lexkong/log"
	"net/http"
	"time"
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kaijian/gin-vue/config"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/routes"
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
	//数据库初始化
	model.DB.Init()
	defer model.DB.Close()//延迟关闭

	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	//路由
	routes.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Infof("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("server.addr"))
	log.Info(http.ListenAndServe(viper.GetString("server.addr"), g).Error())
	//r := gin.Default()
	//r = CollectRoute(r)
	//panic(r.Run())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("server.max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("server.url") + "/test/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

