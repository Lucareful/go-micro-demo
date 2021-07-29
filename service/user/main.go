package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"user/handler"

	user "user/proto/user"
)

func main() {

	// 初始化 consul 对象
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(consulReg), // 注册服务发现
		micro.Version("latest"),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
