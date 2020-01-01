package main

import (
	"context"
	"fmt"

	proto "github.com/golang-study-day/go-micro/exp1/proto"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}
func main() {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"127.0.0.1:8500",}
	})
	service := micro.NewService(
		micro.Registry(reg),
	//service := micro.NewService(
		micro.Name("hellooo"), // 服务名称
	)
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}