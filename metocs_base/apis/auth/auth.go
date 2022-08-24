package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"metocs/base/middleware"
	"metocs/base/models"
	"metocs/base/utils"
)

// Login  用户登录
func Login(c *gin.Context) {
	paramUser := &models.User{}
	err := c.BindJSON(paramUser)
	if err != nil {
		panic("数据解析失败！请检查请求参数！")
	}
	// 验证用户名密码
	user := paramUser.GetUserByPhone()
	passWordErr := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(paramUser.PassWord))
	if passWordErr != nil {
		c.Abort()
		utils.Fail(c, 401, "用户名或密码错误！")
	}

	claims := middleware.UserClaims{
		ID:     user.ID,
		Name:   user.Name,
		Phone:  user.Phone,
		Status: user.Status,
	}

	token := middleware.GenerateToken(&claims)

	utils.Success(c, token)
}
