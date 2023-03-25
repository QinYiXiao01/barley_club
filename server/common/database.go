package common

import (
	"fmt"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// driverName := "mysql"
	host := "localhost"
	port := "3307"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database, err: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
