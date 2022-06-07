package router

import (
	"github.com/gin-gonic/gin"
	"go-gin/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", controller.HelloController)

	return r
}
