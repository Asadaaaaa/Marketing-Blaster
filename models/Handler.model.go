package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	var DB_HOST string = os.Getenv("DB_HOST")
	var DB_USER string = os.Getenv("DB_USER")
	var DB_PASS string = os.Getenv("DB_PASS")

	queryConnect := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ")/marketing_blaster?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(queryConnect), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
