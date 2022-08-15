package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Recover400(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  404,
		"message": "404 Not Found" + c.Request.Method + " " + c.FullPath(),
	})
}

func Recover500(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(200, gin.H{
				"code":    500,
				"message": r,
			})
		}
	}()
	c.Next()
}
