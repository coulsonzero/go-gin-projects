package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create a connection to database")
	}
	fmt.Println("连接mysql数据库成功")
	DB = db
	DB.AutoMigrate(&User{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close connection from database")
	}
	dbSQL.Close()
}
