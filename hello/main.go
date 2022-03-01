package main

import (
	"hello/handler"
	pb "hello/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "hello"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Address("127.0.0.1:8088"),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterHelloHandler(srv.Server(), new(handler.Hello))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
