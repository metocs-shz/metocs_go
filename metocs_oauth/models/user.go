package models

import (
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Phone    string `json:"phone" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Status   uint   `json:"status"`
}

// SerAuthorizationHandler 手机号验证码登录
func SerAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {

	return "123", nil
}
