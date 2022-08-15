package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"metocs_go/database"
	"metocs_go/middleware"
	"metocs_go/router"
	"net/http"
	"time"
)

func ApplicationStart(engine *gin.Engine) {

	database.NewMysqlBase()
	// 最后执行gin 启动服务器
	s := &http.Server{
		Addr:           ":8888",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 添加中间件
	middleware.Settings(engine)
	// 添加全部端点
	router.GroupInit(engine)

	//启动服务器
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
