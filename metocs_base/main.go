package main

import (
	"github.com/gin-gonic/gin"
	"metocs_go/app"
)

func main() {
	//application := config.App
	router := gin.Default()
	app.ApplicationStart(router)

	//router.GET("/ping", func(context *gin.Context) {
	//
	//	context.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//
	//})
	//
	//err := router.Run(":" + strconv.Itoa(application.Server.Port))
	//if err != nil {
	//	return
	//}
}
