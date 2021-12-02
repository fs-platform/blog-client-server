package routes

import (
	"client_server/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(router *gin.Engine) *gin.Engine {
	category := new(controllers.Category)
	article := new(controllers.Article)
	api := router.Group("/api")
	{
		api.GET("/category", category.Index)
		api.GET("/article/:id", article.Index)
	}
	return router
}
