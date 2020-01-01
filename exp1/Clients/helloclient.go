package main

import (
	"context"
	"fmt"

	proto "github.com/golang-study-day/go-micro/exp1/proto"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	//service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"127.0.0.1:8500",}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("helloworldServer001"), // 服务名称
	)

	service.Init()
	helloservice := proto.NewHelloClient("hellooo", service.Client())
	res, err := helloservice.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	if res==nil{
		fmt.Println("nil return")
	}else {
		fmt.Println(res.Msg)
	}
}