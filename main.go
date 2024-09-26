package main

import (
	gin "github.com/gin-gonic/gin"
	"github.com/light-chaser/server/db"
	"github.com/light-chaser/server/global"
	"github.com/light-chaser/server/log"
	"github.com/light-chaser/server/router"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	engine := gin.Default()
	//初始化环境配置
	//初始化日志
	global.LC_LOG = log.InitLogger()
	//初始化数据库连接
	db.InitDataBase()
	//注册路由
	router.RegisterRouters(engine)
	//启动服务
	err := engine.Run(":8888")
	if err != nil {
		return
	}

}
