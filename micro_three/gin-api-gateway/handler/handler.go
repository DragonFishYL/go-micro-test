package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"micor_three/proto/inventory/find"
	"micor_three/proto/rand"
	"net/http"
)

type ApiHandler struct {
	randClient rand.RandService
	findClient find.FindService
}

func GetApiHandler(randClient rand.RandService, findClient find.FindService) *ApiHandler {
	return &ApiHandler{randClient: randClient, findClient: findClient}
}

func (s *ApiHandler) Rand(ctx *gin.Context) {
	request := rand.RandRequest{}
	ctx.ShouldBindQuery(&request)
	response, _ := s.randClient.GetRand(context.Background(), &request)
	ctx.JSON(http.StatusOK, gin.H{
		"result": response.GetResult(),
	})
}

func (s *ApiHandler) Find(ctx *gin.Context) {
	request := find.FindRequest{}
	ctx.ShouldBindQuery(&request)
	response, _ := s.findClient.GetFind(context.Background(), &request)
	fmt.Println("------------------------------")
	fmt.Printf("%+v", request)
	fmt.Printf("%+v", response)
	fmt.Println("------------------------------")
	ctx.JSON(http.StatusOK, gin.H{
		"qty":    response.GetQty(),
		"item":   response.GetItem(),
		"status": response.GetStatus(),
	})
}
