package main

import (
	"github.com/gin-gonic/gin"
	"log"
	controller "luenci.web.com/web/controller"
	"luenci.web.com/web/utils"
)

func main() {
	// gin 服务三板斧

	// 1.初始化路由
	router := gin.Default()

	// 2.路由匹配
	router.Static("/home", "./html")

	// 2.1.路由分组
	v1 := router.Group("/api/v1.0")
	{
		v1.GET("/session", controller.GetSession)
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSmsCd)
		v1.POST("/users", controller.CreateUser)

	}

	// 开启 redis 连接池
	utils.PoolInitRedis("121.199.72.50:6379", "luenci")


	// 3.启动运行
	log.Fatal(router.Run(":8091"))
}
