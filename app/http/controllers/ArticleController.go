package controllers

import (
	"client_server/pkg/code"
	"client_server/pkg/grpc"
	"client_server/proto/blog"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Article struct {
	Controller
}

func (c *Article) Index(ctx *gin.Context) {
	service := grpc.NewBlogService()
	id := ctx.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, c.Format("", code.GetCode("RECORD_PARAMERR"), "error"))
		return
	}
	rep, err := service.Article(context.TODO(), &blog.ArticleRequest{Id: intId})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, c.Format(rep, code.GetCode("RECORD_MICROSERVERERERR"), code.HandleServiceError(err).Detail))
		return
	}
	ctx.JSON(http.StatusOK, c.Format(rep, code.GetCode("RECORD_OK"), "success"))
}
