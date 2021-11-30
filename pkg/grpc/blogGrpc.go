package grpc

import (
	"client_server/proto/blog"
	"github.com/micro/go-micro/v2"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
)

func NewBlogService() blog.BlogService {
	consul := consul2.NewRegistry()
	client := micro.NewService(micro.Registry(consul))
	service := blog.NewBlogService("go.micro.service.blog", client.Client())
	return service
}
