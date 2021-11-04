package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@/golangauth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
