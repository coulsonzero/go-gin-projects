package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/models"
)

type Manager interface {
	AddUser(user *models.User)
	ListUser() []models.User
}
type manager struct {
	db *gorm.DB
}

// AddUser 添加用户
func (m *manager) AddUser(user *models.User) {
	// TODO implement me
	m.db.Create(user)
}

// ListUser 查找用户
func (m manager) ListUser() []models.User {
	// TODO implement me
	var users []models.User
	m.db.Find(&users)
	return users
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to create a connection to database")
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
}
