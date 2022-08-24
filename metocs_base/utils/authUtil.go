package utils

import (
	"github.com/gin-gonic/gin"
	"metocs/base/middleware"
)

func GetClaim(c *gin.Context) *middleware.UserClaims {
	value, exists := c.Get("UserClaims")
	if exists {
		claims := value.(*middleware.UserClaims)
		return claims
	}
	c.Abort()
	Fail(c, 500, "无法获取上下文用户数据！")
	return nil
}
