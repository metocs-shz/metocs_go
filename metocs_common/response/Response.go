package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS_CODE         = 0
	SERVER_ERROR         = 1
	BAD_REQUEST_PARAM    = 2
	METHOD_IS_NOT_ALLOWD = 3
	HANDLER_NOT_Found    = 4
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": "成功！"})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": SUCCESS_CODE, "message": "成功！", "data": data})
}

func BadRequestError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": BAD_REQUEST_PARAM, "message": "请求参数错误！"})
	c.Abort()
}

func ServerError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": SERVER_ERROR, "message": "哎呀！服务器出现错误了！"})
	c.Abort()
}

func ServerErrorMessage(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"code": SERVER_ERROR, "message": "哎呀！服务器出现错误了！", "error": err.Error()})
	c.Abort()
}
