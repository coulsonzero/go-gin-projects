package main

import (
	"gin-grom-mysql/models"
	"gin-grom-mysql/router"
)

func main() {
	db := models.SetupDB()
	defer models.CloseDB(db)
	r := router.SetupRouter()
	r.Run(":8090")
}
