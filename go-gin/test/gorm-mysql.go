package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接mysql数据库失败！")
	}

	db.AutoMigrate(&Employees{})

	// db.Create(&Employees{EmpNo: 10001, BirthDate: "1953-09-02", Gender: "M"})
	// db.Create(&Employees{EmpNo: 10002, BirthDate: "1956-12-10", Gender: "F"})
	// db.Create(&Employees{EmpNo: 10003, BirthDate: "1978-02-26", Gender: "M"})

	// Read
	var emp []Employees
	// db.First(&emp, 10001)
	// db.First(&emp, "EmpNo = ?", 10001)
	// res := db.First(&emp, 10001)

	// res := db.Find(&emp)
	// res := db.Find(&emp, "emp_no = ?", "10002")
	res := db.Raw("SELECT * FROM employees WHERE gender = ? order by birth_date desc", "M").Scan(&emp)

	// Update
	// db.Model(&emp).Update("gender", "F")
	// Update - 更新多个字段
	// db.Model(&emp).Updates(Employees{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&emp).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	// db.Delete(&product, 1)

	fmt.Println(res.RowsAffected)
	fmt.Println(emp)
}

type Employees struct {
	// gorm.Model
	EmpNo     int
	BirthDate string
	Gender    string
}
