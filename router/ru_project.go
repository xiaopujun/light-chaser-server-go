package router

import (
	"github.com/gin-gonic/gin"
	"github.com/light-chaser/server/controller"
)

func ProjectRouters(engine *gin.Engine) {
	group := engine.Group("/api/project")

	group.GET("/list", controller.GetProjectList)
	group.GET("/:id", controller.GetProjectInfo)

}
