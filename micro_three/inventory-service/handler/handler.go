package handler

import (
	"context"
	"micor_three/inventory-service/service"
	"micor_three/proto/inventory/find"
)

type handler struct {
}

func (h handler) GetFind(ctx context.Context, request *find.FindRequest, response *find.FindResponse) error {
	if err := service.GetFind(request, response); err != nil {
		return err
	}
	return nil
}

func Handler() find.FindHandler {
	return handler{}
}
