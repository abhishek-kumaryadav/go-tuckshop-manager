package repositories

import (
	"context"
	"go-tuckshop-manager/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateShopDatabase(foods []model.Food) {
	coll, _ := GetConnection()

	var insertionList []interface{}

	for _, food := range foods {
		filter := bson.D{{"label", food.Label}}

		cursor, cursorErr := coll.Find(context.TODO(), filter)
		if cursorErr != nil {
			panic(cursorErr)
		}
		defer func(cursor *mongo.Cursor, ctx context.Context) {
			err := cursor.Close(ctx)
			if err != nil {
				panic(err)
			}
		}(cursor, context.TODO())

		var results []model.Food
		if err := cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		// Check if any results were found
		if len(results) > 0 {
			update := bson.D{{"$set", bson.D{{"price", food.Price}, {"id", food.ID}}}}
			_, updateErr := coll.UpdateOne(context.TODO(), filter, update)
			if updateErr != nil {
				panic(updateErr)
			}
		} else {
			insertionList = append(insertionList, food)
		}
	}

	_, _ = coll.InsertMany(context.TODO(), insertionList)
}
