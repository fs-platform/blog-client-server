package bootstrap

import (
	"client_server/config"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Initialization() {
	//初始化配置
	config.Initialization()
	//连接数据库
	SetupDB()
	//注册web端路由
	router = SetupRoute()
}

func Run()  {
	router.Run(":8080")
}