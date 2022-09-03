package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"metocs/common/application"
)

var DB *gorm.DB

func MysqlInit() {
	mysqlConfig := application.Application.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlConfig.UserName,
		mysqlConfig.PassWord,
		mysqlConfig.Ip,
		mysqlConfig.Port,
		mysqlConfig.DataBase,
		"10s")
	fmt.Println(dsn)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, _ := open.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	DB = open
}
