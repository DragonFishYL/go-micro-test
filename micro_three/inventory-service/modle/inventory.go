package modle

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnInventory(dbName string, collName string) *mongo.Collection {
	con := ClientPool{}
	conn, err := con.GetClient()
	if err != nil {
		panic(err)
	}
	inv := conn.Client.Database(dbName).Collection(collName)
	return inv
}
