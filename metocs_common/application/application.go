package application

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Server      *server
	MysqlConfig *mysqlConfig
	Auth        *auth
	Redis       *redis
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

type auth struct {
	AccessTokenExp    int
	RefreshTokenExp   int
	IsGenerateRefresh bool
}

type redis struct {
	Ip       string
	Password string
	Port     int
	DataBase int
}

func ApplicationInit() *AppConfig {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("read config failed: %v", err)
	}

	return &AppConfig{
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
		Auth: &auth{
			AccessTokenExp:    viper.GetInt("auth.accessTokenExp"),
			RefreshTokenExp:   viper.GetInt("auth.refreshTokenExp"),
			IsGenerateRefresh: viper.GetBool("auth.isGenerateRefresh"),
		},
		Redis: &redis{
			Ip:       viper.GetString("redis.ip"),
			Port:     viper.GetInt("redis.port"),
			Password: viper.GetString("redis.password"),
			DataBase: viper.GetInt("redis.database"),
		},
	}
}

var Application *AppConfig

func init() {
	Application = ApplicationInit()
}
