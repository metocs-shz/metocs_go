package utils

import (
	"gorm.io/gorm"
	"metocs_go/database"
)

func Page(pageSize, pageNum int) *gorm.DB {
	return database.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize)
}
