package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"metocs/common/authTool"
	"metocs/common/database"
	"net/http"
	"time"
)

type User struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Phone    string `json:"phone" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Status   int    `json:"status"`
}

// SerAuthorizationHandler 手机号验证码登录
func SerAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	phone := r.FormValue("phone")
	formValue := r.FormValue("password")
	//查询当前用户是否存在
	var user User
	tx := database.DB.Where("phone = ?", phone).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	decoder := authTool.PassWordDecoder(user.PassWord, formValue)
	if !decoder {
		return "", authTool.ErrorPassWordDecoder
	}
	return user.ID, nil
}

func PasswordAuthorizationHandler(ctx context.Context, clientID, username, password string) (userID string, err error) {
	//查询当前用户是否存在
	var user User
	tx := database.DB.Where("phone = ?", username).First(&user)
	if tx.Error != nil {
		return "", tx.Error
	}
	decoder := authTool.PassWordDecoder(user.PassWord, password)
	if !decoder {
		return "", errors.New("账号或密码错误！")
	}
	return user.ID, nil

}
