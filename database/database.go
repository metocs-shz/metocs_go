package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"metocs_go/config"
)

var MysqlDataBase *gorm.DB

func NewMysqlBase() {
	app := config.App
	mysqlConfig := app.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlConfig.UserName,
		mysqlConfig.PassWord,
		mysqlConfig.Ip,
		mysqlConfig.Port,
		mysqlConfig.DataBase,
		"10s")
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	MysqlDataBase = open
}
