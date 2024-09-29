package server

import (
	"github.com/gin-gonic/gin"
	"github.com/light-chaser/server/global"
	"github.com/light-chaser/server/router"
	"strconv"
)

func RunServer() {
	gin.SetMode(global.LC_ENV.Server.Mode)
	engine := gin.Default()
	//注册路由
	router.RegisterRouters(engine)
	port := ":" + strconv.Itoa(global.LC_ENV.Server.Port)
	//启动服务
	err := engine.Run(port)
	if err != nil {
		panic(err)
	}
}
