package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"micor_three/inventory-service/modle"
	"micor_three/proto/inventory/find"
)

func GetFind(req *find.FindRequest, resp *find.FindResponse) error {
	findObj := modle.ConnInventory("local", "inventory")
	reObj, err := findObj.Find(
		context.TODO(),
		bson.M{
			"qty": req.GetQty(),
		},
		options.Find().SetLimit(10),
		options.Find().SetSort(
			bson.M{
				"qty": -1,
			},
		),
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		panic(err)
	}
	var (
		inventory      find.FindResponse
		inventorySlice = make([]find.FindResponse, 0)
	)
	for reObj.Next(context.TODO()) {
		if err = reObj.Decode(&inventory); err != nil {
			panic(err)
		}
		inventorySlice = append(inventorySlice, inventory)
	}
	if len(inventorySlice) > 0 {
		*resp = inventorySlice[0]
	}
	resp.Item = "aaaaa"
	resp.Qty = "AA"
	return nil
}
