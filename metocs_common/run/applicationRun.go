package run

import (
	"github.com/gin-gonic/gin"
	"log"
	"metocs/common/application"
	"metocs/common/database"
	"net/http"
	"strconv"
	"time"
)

func Run() {
	// 获取配置数据
	config := application.Application
	// 创建mysql数据源
	database.MysqlInit()
	//创建redis数据园
	database.RedisInit()

	//启动服务器监听
	engine := gin.Default()

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.Server.Port) + config.Server.Path,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
