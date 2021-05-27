package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const DB_URL = "potato:potato@tcp(127.0.0.1:3306)/demo_golang?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB
var DBError error

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitDabase() {
	DB, DBError = gorm.Open(mysql.Open(DB_URL), &gorm.Config{})
	if DBError != nil {
		log.Println(DBError)
		os.Exit(-1)
	}

	DB.AutoMigrate(&User{})

	log.Println("Database Initialized")
}
