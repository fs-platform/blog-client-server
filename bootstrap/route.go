package bootstrap

import (
	"client_server/app/http/middlewares"
	"client_server/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	//修改logger格式
	r.Use(LoggerFormat(), middlewares.Cors())
	//注册web端路由
	routes.RegisterWebRouter(r)
	return r
}
