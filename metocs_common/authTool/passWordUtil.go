package authTool

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"metocs/common/response"
)

var (
	ErrorPassWordDecoder = errors.New("账号或密码错误！")
)

type PanicPassword struct {
	code    int
	message string
}

// PassWordEncoder 加密密码
func PassWordEncoder(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(hash)
}

// PassWordDecoder  验证密码
func PassWordDecoder(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		mesasgae := PanicPassword{
			code:    response.AUTH_DEIND,
			message: "用户名密码错误！",
		}
		panic(mesasgae)
	}
	return true
}

func GetLoginUser(c *gin.Context) {

}
