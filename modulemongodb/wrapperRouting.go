package modulemongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/av-belyakov/comparisondatabase/datamodels"
)

// wrapperRoutingGetFullCount выполняет поиск общего количества объектов в коллекции
func wrapperRoutingGetFullCount(collection *mongo.Collection) (int64, error) {
	fmt.Println("func 'wrapperRoutingGetFullCount', reseived full count object")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Count().SetMaxTime(2 * time.Second)
	count, err := collection.CountDocuments(ctx, bson.D{}, opts)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func wrapperRoutingGetLimitObject(collection *mongo.Collection, offset, limitMaxSize int64) ([]datamodels.ElementSTIXObject, error) {
	fmt.Println("func 'wrapperRoutingGetLimitObject', reseived full count object")

	sortOrder := -1
	resultObject := make([]datamodels.ElementSTIXObject, 0, limitMaxSize)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	options := options.Find().SetAllowDiskUse(true).SetSort(bson.D{{Key: "_id", Value: sortOrder}, {Key: "commonpropertiesobjectstix.id", Value: sortOrder}}).SetSkip(offset).SetLimit(limitMaxSize)

	cur, err := collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return resultObject, err
	}

	for cur.Next(context.Background()) {
		elem, err := GetListElementSTIXObject(cur)
		if err != nil {
			return resultObject, err
		}

		resultObject = append(resultObject, elem)
	}

	return resultObject, nil
}
