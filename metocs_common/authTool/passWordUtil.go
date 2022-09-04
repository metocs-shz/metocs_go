package authTool

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorPassWordDecoder = errors.New("账号或密码错误！")
)

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
		fmt.Println(err.Error())
		return false
	}
	return true
}
