package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/light-chaser/server/global"
	"github.com/light-chaser/server/model/database"
)

func GetProjectList(c *gin.Context) {
	var project database.DbProject
	global.LC_DB.First(&project, 27)
	c.JSON(200, gin.H{
		"msg":  "project list 大牛逼",
		"data": project,
	})
}

func GetProjectInfo(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"msg": "获取项目信息",
		"id":  id,
	})
}
