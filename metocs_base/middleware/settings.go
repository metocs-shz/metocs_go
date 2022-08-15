package middleware

import (
	"github.com/gin-gonic/gin"
)

func Settings(engine *gin.Engine) {

	engine.Use(Recover500)

	engine.NoRoute(Recover400)
	engine.NoMethod(Recover400)
}
