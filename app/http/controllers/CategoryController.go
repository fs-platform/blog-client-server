package controllers

import (
	"client_server/pkg/code"
	"client_server/proto/blog"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/service/grpc"
	"net/http"
)

type Category struct {
	Controller
}

func (c *Category) Index(ctx *gin.Context) {
	service := blog.NewBlogService("go.micro.service.blog", grpc.NewService().Client())
	rep, err := service.Category(context.TODO(), &blog.CategoryIndexRequest{})
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, c.Format(rep, code.GetCode("RECORD_OK"), "success"))
}
