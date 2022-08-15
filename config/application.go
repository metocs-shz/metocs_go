package config

import (
	"github.com/spf13/viper"
	"log"
)

type Application struct {
	Server      *server
	MysqlConfig *mysqlConfig
}

type server struct {
	Port int
	Path string
}

type mysqlConfig struct {
	Ip       string
	Port     int
	UserName string
	PassWord string
	DataBase string
}

func ApplicationInit() *Application {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	return &Application{
		Server: &server{
			Port: viper.GetInt("server.port"),
			Path: viper.GetString("server.path"),
		},
		MysqlConfig: &mysqlConfig{
			Ip:       viper.GetString("mysql.ip"),
			Port:     viper.GetInt("mysql.port"),
			UserName: viper.GetString("mysql.username"),
			PassWord: viper.GetString("mysql.password"),
			DataBase: viper.GetString("mysql.database"),
		},
	}
}

var App = new(Application)

func init() {
	//注入配置参数
	App = ApplicationInit()
}
