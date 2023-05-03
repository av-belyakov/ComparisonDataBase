package modulemongodb_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ComparisonDataBase/datamodels"
	"ComparisonDataBase/modulemongodb"
)

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

var _ = Describe("Moduleinteraction", Ordered, func() {
	var (
		mdbErr     error
		mdbConnect *mongo.Client
		collection *mongo.Collection
	)

	conf := datamodels.ConfMongoDB{
		Host:       "192.168.9.208",
		Port:       37017,
		DBname:     "isems-mrsict",
		Collection: "stix_object_collection",
		User:       "module-isems-mrsict",
		Passwd:     "vkL6Znj$Pmt1e1",
	}

	BeforeAll(func() {
		mdbConnect, mdbErr = modulemongodb.CreateConnection(&conf)
		collection = mdbConnect.Database(conf.DBname).Collection(conf.Collection)
	})

	Context("Тест 1. Проверка подключения к БД", func() {
		It("При подключении к БД ошибки быть не должно", func() {
			Expect(mdbErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка выполнения запроса для получения общего количества объектов в БД", func() {
		It("Должно быть найдено определенной количество объектов STIX", func() {
			cur, err := collection.Find(context.Background(), bson.D{})

			fmt.Println(cur)
			listID := []definingTypeSTIXObject{}
			for cur.Next(context.Background()) {
				var modelType definingTypeSTIXObject
				if err := cur.Decode(&modelType); err != nil {
					continue
				}

				listID = append(listID, modelType)
			}

			//fmt.Println("listID = ", listID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(listID)).Should(Equal(541))
		})
	})

	/*
		//для подсчета количества объектов

		cur, err := collection.Aggregate(
			context.TODO(),
			mongo.Pipeline{
				bson.D{bson.E{Key: "$match", Value: bson.D{
					bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
				}}},
				bson.D{
					bson.E{Key: "$group", Value: bson.D{
						bson.E{Key: "_id", Value: fmt.Sprintf("$outside_specification.%s", outsideSpecificationField)},
						bson.E{Key: "count", Value: bson.D{
							bson.E{Key: "$sum", Value: 1},
						}},
					}}}},
			opts)
		if err != nil {
			return fn, err
		}

		err = cur.All(context.TODO(), &tmpResults)
		if err != nil {
			return fn, err
		}

		for _, v := range tmpResults {
			name, ok := v["_id"].(string)
			if !ok {
				continue
			}

			if count, ok := v["count"].(int32); ok {
				rsiSTIXObject.ListComputerThreat[name] = count
			}
		}
	*/
})
