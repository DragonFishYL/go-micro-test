package main

import (
	"go-micro.dev/v4"
	"micor_three/proto/rand"
	"micor_three/rand-service/handler"
)

func main() {
	//服务配置
	service := micro.NewService(
		micro.Name("rand-serice"),
		//micro.Address("127.0.0.1:8001"),
		micro.Version("latest"),
	)
	service.Init()
	//服务注册
	_ = rand.RegisterRandHandler(service.Server(), handler.Handler())
	//服务启动
	if err := service.Run(); err != nil {
		panic(err)
	}
}
