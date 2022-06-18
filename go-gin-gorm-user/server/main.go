package main

import (
	"log"
	"server/routers"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("connect port error")
	}
}
