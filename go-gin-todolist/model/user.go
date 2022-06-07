package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func GetAllUsers() *[]User {
	var users []User
	DB.Find(&users)
	return &users
}

func FindUserByID(id int) *User {
	var user User
	DB.First(&user, id)
	return &user
}

func (u *User) Create() {
	DB.Create(&u)
}

func (u *User) Add(name string, age int) {
	DB.Create(&User{
		Name: name,
		Age:  age,
	})
}

func (u *User) Update() {
	DB.Model(&u).Updates(User{
		Name: u.Name,
		Age:  u.Age,
	})
}

func (u *User) Delete() {
	DB.Delete(&u)
}
