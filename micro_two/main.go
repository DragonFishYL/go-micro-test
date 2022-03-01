package main

import (
	"context"
	"fmt"
	"micro_two/handler"
	pb "micro_two/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "micro_two"
	version = "latest"
)

func runClient(services micro.Service) {
	// Create new greeter client
	microTwo := pb.NewMicroTwoService(service, services.Client())

	// Call the greeter
	rsp, err := microTwo.Call(context.TODO(), &pb.CallRequest{Name: service})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Msg)
}

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterMicroTwoHandler(srv.Server(), handler.NewMicroTwo())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
