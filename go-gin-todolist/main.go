package main

import (
	"go-gin-todolist/model"
	"go-gin-todolist/router"
)

func main() {
	db := model.SetupDB()
	defer model.CloseDB(db)

	r := router.SetupRouter()
	r.Run(":7080")
}
