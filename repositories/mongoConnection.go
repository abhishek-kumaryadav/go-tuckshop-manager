package repositories

import (
	"context"
	"go-tuckshop-manager/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func GetConnection() (*mongo.Collection, *mongo.Client) {
	if client != nil && collection != nil {
		return collection, client
	}

	properties := services.GetEnvProperties()

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().
		ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	}
	collection = client.Database(properties["database"]).Collection(properties["collection"])
	return collection, client
}
