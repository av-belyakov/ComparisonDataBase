package modulemongodb_test

import (
	"context"
	"fmt"

	"github.com/RediSearch/redisearch-go/redisearch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/modulemongodb"
	"github.com/av-belyakov/comparisondatabase/moduleredisearch"
)

func getListIndex(cur *mongo.Cursor) ([]datamodels.IndexObject, error) {
	listIndexObj := []datamodels.IndexObject{}
	for cur.Next(context.Background()) {
		elem, err := modulemongodb.GetListElementSTIXObject(cur)
		if err != nil {
			return listIndexObj, err
		}

		listIndexObj = append(listIndexObj, moduleredisearch.GetIndex(elem))
	}

	return listIndexObj, nil
}

func getRedisearchDocument(cur *redisearch.Client, listIndex []datamodels.IndexObject) []redisearch.Document {
	redisearchDoc := make([]redisearch.Document, 0, len(listIndex))

	for _, v := range listIndex {
		tmp := redisearch.NewDocument(v.ID, 1.0)
		tmp.Set("name", v.Name)
		tmp.Set("description", v.Description)
		tmp.Set("street_address", v.StreetAddress)
		tmp.Set("abstract", v.Abstract)
		tmp.Set("content", v.Content)
		tmp.Set("value", v.Value)

		redisearchDoc = append(redisearchDoc, tmp)
	}

	return redisearchDoc
}

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

var _ = Describe("Moduleinteraction", Ordered, func() {
	var (
		mdbErr, mdbFindErr error
		createListIndexErr error
		rsdbErr, rsdbAddIndexErr/*, rsdbFindErr*/ error
		mdbConnect   *mongo.Client
		collection   *mongo.Collection
		cursor       *mongo.Cursor
		connRDB      *redisearch.Client
		listIndexObj []datamodels.IndexObject
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
		//создаем соединение с СУБД MongoDB
		mdbConnect, mdbErr = modulemongodb.CreateConnection(&conf)
		collection = mdbConnect.Database(conf.DBname).Collection(conf.Collection)

		//создаем соединение с СУБД Redisearch
		connRDB, rsdbErr = moduleredisearch.CreateConnection(datamodels.ConfRedisearch{
			Host: "192.168.9.208",
			Port: 6379,
		})

		//выполняем поиск
		cursor, mdbFindErr = collection.Find(context.Background(), bson.D{})
		//формируем список индексов
		listIndexObj, createListIndexErr = getListIndex(cursor)
		//listIndexObj, createListIndexErr = moduleredisearch.GetListIndex(cursor)

		rsdbAddIndexErr = connRDB.IndexOptions(
			redisearch.IndexingOptions{
				Replace: true,
				Partial: true,
			}, getRedisearchDocument(connRDB, listIndexObj)...)
	})

	Context("Тест 1. Проверка подключения к БД MongoDB", func() {
		It("При подключении к БД MongoDB ошибки быть не должно", func() {
			Expect(mdbErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка выполнения запроса для получения общего количества объектов в БД", func() {
		It("Должно быть найдено определенной количество объектов STIX", func() {
			cur, err := collection.Find(context.Background(), bson.D{})

			listID := []definingTypeSTIXObject{}
			for cur.Next(context.Background()) {
				var modelType definingTypeSTIXObject
				if err := cur.Decode(&modelType); err != nil {
					continue
				}

				listID = append(listID, modelType)
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(listID)).Should(Equal(541))
		})
	})

	Context("Тест 3. Проверка выполнения поиска в MongoDB", func() {
		It("Поиск должен быть выполнен без ошибок", func() {
			Expect(mdbFindErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Проверка создания списка индексов", func() {
		It("При создании списка индексов ошибок быть не должно", func() {
			Expect(createListIndexErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 5. Проверка наличия заданного количества индексов в списке", func() {
		It("Должен быть создан спискок индексов с заданным количеством", func() {
			Expect(len(listIndexObj)).Should(Equal(541))
		})
	})

	Context("Тест 6. Проверка подключения к СУБД Redisearch", func() {
		It("При подключении к СУБД Redisearch ошибки быть не должно", func() {
			Expect(rsdbErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 7. Проверка добавления индексов в Redisearch", func() {
		It("При добавления списка индексов ошибок быть не должно", func() {
			Expect(rsdbAddIndexErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 8. Проверка поиска информации по индексам в Redisearch", func() {
		It("Должно быть добавленно (541) новых индексов", func() {
			_, num, err := connRDB.Search(redisearch.NewQuery("*").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				).
				SetReturnFields("id"))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(num).Should(Equal(541))
		})

		It("При выполнении поиска должно быть найден один ID объекта где поле 'name' равно 'frogfrog_list.txt'", func() {
			listName, numName, err := connRDB.Search(redisearch.NewQuery("frogfrog_list.txt").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				).
				SetReturnFields("id"))

			fmt.Printf("______FULL SEARCH DOCUMENTS name = frogfrog_list.txt: %v\n", listName[0].Id)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(listName[0].Id).Should(Equal("file--9b98cee2-06af-4f2e-a23a-3762c3e40bb9"))
			Expect(numName).Should(Equal(1))
		})

		It("При выполнении поиска должно быть найден один ID объекта где поле 'description' содержит 'testy to try'", func() {
			listDesc, numDesc, err := connRDB.Search(redisearch.NewQuery("testy to try").
				AddFilter(
					redisearch.Filter{
						Field: "description",
					},
				).
				SetReturnFields("id"))

			fmt.Printf("______FULL SEARCH DOCUMENTS description contains 'testy to try': %v\n", listDesc[0].Id)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(listDesc[0].Id).Should(Equal("report--0c6a75be-d979-4646-a92b-58ae4ab2d95d"))
			Expect(numDesc).Should(Equal(1))
		})

		It("При выполнении поиска должно быть найден один ID объекта где поле 'description' содержит 'Лягушки бывают разные.'", func() {
			listDesc, numDesc, err := connRDB.Search(redisearch.NewQuery("Лягушки бывают разные.").
				AddFilter(
					redisearch.Filter{
						Field: "description",
					},
				).
				SetReturnFields("id", "name", "description", "street_address", "abstract", "content", "value"))

			fmt.Printf("______FULL SEARCH DOCUMENTS description contains 'Лягушки бывают разные.': %s\n", listDesc[0].Id)
			fmt.Printf("______ALL INDEX: %v", listDesc[0].Properties)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(listDesc[0].Id).Should(Equal("report--193e108f-6cb3-474f-b4d3-fb7a86ebdca1"))
			Expect(numDesc).Should(Equal(1))
		})
	})
})
