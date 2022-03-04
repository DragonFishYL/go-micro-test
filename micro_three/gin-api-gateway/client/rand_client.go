package client

import (
	"go-micro.dev/v4"
	"micor_three/proto/rand"
)

var (
	randClient rand.RandService
)

func GetRandClient() rand.RandService {
	if randClient == nil {
		service := micro.NewService()
		randClient = rand.NewRandService("rand-serice", service.Client())
	}
	return randClient
}
