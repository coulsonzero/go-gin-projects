package api

import (
	"github.com/gin-gonic/gin"
	"go-gin-todolist/model"
	"net/http"
	"strconv"
)

// CreateUser handles HTTP POST request to create a new User
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Create()
	c.JSON(http.StatusCreated, user)
}

func AddUser(c *gin.Context) {
	name := c.PostForm("name")
	ageParam := c.PostForm("age")
	age, err := strconv.Atoi(ageParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Add(name, age)

	c.JSON(http.StatusOK, user)
}

// ShowUser handles HTTP GET request to find a User by id param
func ShowUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.FindUserByID(id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser handles HTTP PATCH request to find and update User details
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.FindUserByID(id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Update()
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles HTTP DELETE request to find and delete User record
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.FindUserByID(id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	user.Delete()
	c.JSON(http.StatusOK, gin.H{})
}

// ShowUsers handles HTTP GET request to return all User records
func ShowUsers(c *gin.Context) {
	users := model.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
