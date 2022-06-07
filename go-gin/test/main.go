package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// GET无参数: curl http://localhost:8080/hello
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin !",
		})
	})

	// GET解析路径参数: /user/coulson
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// GETQuery参数: users?name=coulson&role=developer，role可选
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	r.GET("/api/v1/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"list":    []string{"home", "about", "docs", "footer"},
			"data":    "asas",
			"bar": []interface{}{
				&Bar{
					Title:   "条形图示例",
					Content: "直方图",
				},
			},
			"dateinfo": []interface{}{
				&DateInfo{
					dvp_1: "2022-04-09",
					dvp_2: "2022-04-09",
					dvp_3: "20220409",
				},
			},
			"msg":  "请求成功",
			"code": 1,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "这是post请求返回的数据")
	})

	// curl http://localhost:8080/form  -X POST -d 'username=geektutu&password=1234'
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合
	// curl "http://localhost:8080/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// curl -g "http://localhost:8080/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/postmap", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path":    c.FullPath(),
			"success": "ture",
		})
	}

	// group: v1
	// http://localhost:8080/api/v1/task
	v1 := r.Group("/api/v1")
	{
		v1.GET("/task", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/task", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000")
}

type Bar struct {
	Title   string
	Content string
}

type DateInfo struct {
	dvp_1 string
	dvp_2 string
	dvp_3 string
}
