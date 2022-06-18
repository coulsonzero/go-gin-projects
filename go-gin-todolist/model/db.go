package model

import (
	"go-gin-todolist/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	// func SetupDB(dsn string) {
	/* ====== Step-2: 连接mysql数据库 ====== */
	dsn := conf.LoadIntConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	DB = db

	DB.AutoMigrate(&User{})
	// DB.Create(&User{
	// 	Name: "shville",
	// 	Age:  25,
	// })
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
