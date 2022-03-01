package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4"
	proto "hello/proto"
)

func main() {
	// 创建新服务
	service := micro.NewService(micro.Name("hello.client"))
	// 服务初始化
	service.Init()

	// 创建RPC的客户端实例
	greeter := proto.NewHelloService("hello", service.Client())

	// 发起RPC调用
	rsp, err := greeter.Call(context.TODO(), &proto.CallRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// 打印返回值
	fmt.Println(rsp.Msg)
}
