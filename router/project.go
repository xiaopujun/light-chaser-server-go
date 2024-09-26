package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ProjectRouters(engine *gin.Engine) {
	group := engine.Group("/api/project")
	group.GET("/list", ProjectList)
}

func ProjectList(c *gin.Context) {
	fmt.Println("project list")
	c.JSON(200, gin.H{
		"msg": "project list 大牛逼",
	})
}
