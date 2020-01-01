package main

import (
	"log"
	"time"

	hello "github.com/golang-study-day/go-micro/exp3/proto"
	"github.com/micro/go-micro"
	grpc "google.golang.org/grpc"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"context"
	"net"
	"fmt"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request)(rsp *hello.Response,err error) {
	log.Print("Received Say.Hello request")
	r:=new(hello.Response)
	r.Msg = "Hello " + req.Name
	return r,nil
}

func main() {
	reg:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs=[]string{"127.0.0.1:8500",}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":9090"),
	)

	// optionally setup command line usage
	service.Init()

	s:=grpc.NewServer()
	// Register Handlers
	hello.RegisterSayServer(s, new(Say))
	//hello.RegisterSayHandler(service.Server(), new(Say))

	port := ":9090"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)

	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	s.Serve(l)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
