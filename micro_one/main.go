package main

import (
	"micro_one/handler"
	pb "micro_one/proto"

	log "go-micro.dev/v4/logger"
)

var (
	service = "micro_one"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterMicro_oneHandler(srv.Server(), new(handler.Micro_one))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
