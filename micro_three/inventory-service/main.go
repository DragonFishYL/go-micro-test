package main

import (
	"go-micro.dev/v4"
	"micor_three/inventory-service/handler"
	"micor_three/proto/inventory/find"
)

func main() {
	service := micro.NewService(
		micro.Name("inventory-serice"),
		micro.Version("latest"),
	)
	service.Init()
	_ = find.RegisterFindHandler(service.Server(), handler.Handler())
	if err := service.Run(); err != nil {
		panic(err)
	}
}
