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

func ApplicationInit() {
	// 创建mysql数据源
	database.MysqlInit()
	//创建redis数据园
	database.RedisInit()
}

func Run(engine *gin.Engine) {
	// 获取配置数据
	server := application.Application.Server
	//启动服务器监听
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(server.Port),
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
