package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"luenci.web.com/service/getCaptcha/handler"

	getCaptcha "luenci.web.com/service/getCaptcha/proto/getCaptcha"
)

func main() {
	// 初始化 consul 对象
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Address("127.0.0.1:"),// 防止随机生成port
		micro.Name("go.micro.service.getCaptcha"),
		micro.Registry(consulReg), // 注册服务发现
		micro.Version("latest"),
	)

	// Register Handler
	err := getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))
	if err != nil {
		log.Println(err)
		return
	}


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
