package models

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load(".env")

	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local

	db_config := DB_USERNAME + ":" + DB_PASSWORD + "@(" + DB_HOST + ")/" + DB_NAME + "?charset=utf8mb4&loc=Local"

	db, err := gorm.Open("mysql", db_config)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{})

	DB = db
}