package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-todolist/api"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello todolist!")
	})

	r.GET("/api/v1/users", api.ShowUsers)
	// r.GET("api/v1/users/:id", api.ShowUser)
	r.GET("api/v1/users/name", api.GetUserByName)
	r.POST("/api/v1/users/add", api.AddUser)

	return r
}
