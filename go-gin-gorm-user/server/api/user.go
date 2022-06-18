package api

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"server/models"
	"server/response"
)

func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}

	dao.Mgr.AddUser(&user)
	response.Success("添加成功", user, c)
}

func ListUser(c *gin.Context) {
	users := dao.Mgr.ListUser()
	response.Success("查询成功", users, c)
}
