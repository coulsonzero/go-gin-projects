package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// 新增数据
func NewUser(name string, email string) (*User, error) {
	var err error
	if name == "" {
		err = errors.New("name can not be blank")
		return nil, err
	}

	if email == "" {
		err = errors.New("email can not be blank")
		return nil, err
	}

	return &User{Name: name, Email: email}, err
}

// 查询数据 by id
func FindUserByID(id int) *User {
	var user User
	DB.First(&user, id)
	return &user
}

// 查询全部数据
func GetAllUsers() *[]User {
	var users []User
	DB.Find(&users)
	return &users
}

// Create creates a new User record in the database
func (u *User) Create() {
	DB.Create(&u)
}

// 更改数据
func (u *User) Update() {
	DB.Model(&u).Updates(User{Name: u.Name, Email: u.Email})
}

// 删除数据
func (u *User) Delete() {
	DB.Delete(&u)
}
