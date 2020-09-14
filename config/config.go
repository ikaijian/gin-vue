package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error  {
	c := Config{
		Name: cfg,
	}

	//初始化配置配置文件
	if err :=c.initConfig(); err != nil {
		return err
	}
	
	//监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	}else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("config")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为APISERVER
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".","_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig();err != nil {// viper解析配置文件
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file change:%s",e.Name)
	})
}