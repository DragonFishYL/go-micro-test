package handler

import (
	"context"
	"micor_three/proto/rand"
	"micor_three/rand-service/service"
)

type handler struct {
}

func (h handler) GetRand(ctx context.Context, request *rand.RandRequest, response *rand.RandResponse) error {
	max := request.GetMax()
	response.Result = service.GetRand(max)
	return nil
}

func Handler() rand.RandHandler {
	return handler{}
}
