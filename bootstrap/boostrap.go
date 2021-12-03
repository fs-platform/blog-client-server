package bootstrap

import (
	"client_server/config"
	config2 "client_server/pkg/config"
	"client_server/pkg/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var router *gin.Engine

func Initialization() {
	//初始化配置
	config.Initialization()
	//注册gin mode
	if config2.GetString("app.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	//注册grpc
	grpc.New()
	//连接数据库
	SetupDB()
	//注册web端路由
	router = SetupRoute()
}

func Run() {
	s := &http.Server{
		Addr:           ":10025",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
