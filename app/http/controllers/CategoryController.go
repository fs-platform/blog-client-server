package controllers

import (
	"client_server/pkg/code"
	"client_server/pkg/grpc"
	"client_server/proto/blog"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Category struct {
	Controller
}

func (c *Category) Index(ctx *gin.Context) {
	service := grpc.NewBlogService()
	rep, err := service.Category(context.TODO(), &blog.CategoryIndexRequest{})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, c.Format(rep, code.GetCode("RECORD_MICROSERVERERERR"), "error"))
		return
	}
	ctx.JSON(http.StatusOK, c.Format(rep, code.GetCode("RECORD_OK"), "success"))
}
