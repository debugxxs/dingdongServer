package main

import (
	"dindongwork/common"
	"dindongwork/router"
	"dindongwork/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	/*
	1.加载项目配置文件
	2.加载数据库初始化连接
	3.加载路由
	4.启动服务器
	*/
	cfg,err := tools.ParsingConfig("./config/config.json")
	common.ErrHandler("初始化配置失败",err)

	tools.InitDbEngine(cfg)

	app := gin.Default()

	router.LoadRouter(app)

	_=app.Run(cfg.AppHost +":"+ cfg.AppPort)
}
