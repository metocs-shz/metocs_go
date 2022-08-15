package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "成功", "data": data})
}

func SuccessNoMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "成功"})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
