package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

// SELECT * from users;
func GetAllUsers() *[]User {
	var users []User
	DB.Find(&users)
	return &users
}

// SELECT * from users WHERE id = ?;
func FindUserByID(id int) *User {
	var user User
	DB.First(&user, id)
	return &user
}

func (u *User) Create() {
	DB.Create(&u)
}

// INSERT INTO users(name, age) VALUES(?, ?);
func (u *User) Add(name string, age int) {
	DB.Create(&User{
		Name: name,
		Age:  age,
	})
}

// UPDATE users SET name = ?, age = ?;
func (u *User) Update() {
	DB.Model(&u).Updates(User{
		Name: u.Name,
		Age:  u.Age,
	})
}

// DELETE FROM users;
func (u *User) Delete() {
	DB.Delete(&u)
}
