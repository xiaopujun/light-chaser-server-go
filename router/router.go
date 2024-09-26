package router

import "github.com/gin-gonic/gin"

func RegisterRouters(engine *gin.Engine) {
	ProjectRouters(engine)
}
