package main

import (
	"github.com/light-chaser/server/config"
	"github.com/light-chaser/server/db"
	"github.com/light-chaser/server/global"
	"github.com/light-chaser/server/log"
	"github.com/light-chaser/server/server"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	//初始化日志
	global.LC_LOG = log.InitLogger()
	//初始化环境配置
	global.LC_ENV = config.LoadEvn()
	//初始化数据库连接
	global.LC_DB = db.InitDatabase()
	//启动服务
	server.RunServer()
}
