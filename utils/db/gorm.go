package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3307)/crm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
