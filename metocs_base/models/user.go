package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	database "metocs_go/database"
	"metocs_go/utils"
	"time"
)

type User struct {
	ID       string `gorm:"primarykey"`
	Phone    string `json:"phone" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Status   int    `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) GetUserByID(param string) {
	database.DB.Where("ID = ?", param).Take(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := utils.GetID()
	u.ID = id
	password, err := bcrypt.GenerateFromPassword([]byte(u.PassWord), bcrypt.MinCost)
	if err != nil {
		panic("用户密码加密失败！")
	}
	u.PassWord = string(password)
	return
}

func (u *User) GetUserByPhone() *User {
	user := &User{}
	database.DB.Where("phone = ?", u.Phone).Take(user)
	return user
}

func CreateUser(user *User) {
	tx := database.DB.Create(user)
	if tx.Error != nil {
		panic(tx.Error.Error())
	}
}
