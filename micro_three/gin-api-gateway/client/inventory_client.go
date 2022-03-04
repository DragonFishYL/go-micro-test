package client

import (
	"go-micro.dev/v4"
	"micor_three/proto/inventory/find"
)

var (
	findClient find.FindService
)

func GetFindClient() find.FindService {
	if findClient == nil {
		service := micro.NewService()
		findClient = find.NewFindService("inventory-serice", service.Client())
	}
	return findClient
}
