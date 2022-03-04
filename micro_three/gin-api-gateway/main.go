package main

import (
	"github.com/gin-gonic/gin"
	"micor_three/gin-api-gateway/client"
	"micor_three/gin-api-gateway/handler"
)

func main() {
	r := gin.Default()
	//定义路由
	route := r.Group("/three")

	//获取监听的服务
	randClient := client.GetRandClient()
	findClient := client.GetFindClient()
	randHandler := handler.GetApiHandler(randClient, findClient)
	//定义路由
	route.GET("rand", randHandler.Rand)
	route.GET("find", randHandler.Find)
	//服务启动
	if err := r.Run(); err != nil {
		panic(err)
	}
}
