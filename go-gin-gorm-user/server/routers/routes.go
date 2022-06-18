package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/api"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 跨域访问
	myCORS := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})

	r.Use(myCORS)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/users", api.ListUser)
	r.POST("/add", api.AddUser)
	return r
}
