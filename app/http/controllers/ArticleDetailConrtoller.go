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

type ArticleDetail struct {
	Controller
}

func (c *ArticleDetail) Index(ctx *gin.Context) {
	service := grpc.NewBlogService()
	id := ctx.Param("articleId")
	intId, err := strconv.ParseInt(id, 10, 64)
	rep, err := service.ArticleDetail(context.TODO(), &blog.ArticleDetailRequest{Id: intId})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, c.Format(rep, code.GetCode("RECORD_MICROSERVERERERR"), code.HandleServiceError(err).Detail))
		return
	}
	ctx.JSON(http.StatusOK, c.Format(rep, code.GetCode("RECORD_OK"), "success"))
}
