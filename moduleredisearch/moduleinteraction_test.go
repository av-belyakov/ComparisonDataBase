package moduleredisearch_test

import (
	"fmt"

	"github.com/RediSearch/redisearch-go/redisearch"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/moduleredisearch"
)

var _ = Describe("Moduleinteraction", Ordered, func() {
	var (
		rErr error
		conn *redisearch.Client
	)

	BeforeAll(func() {
		conn, rErr = moduleredisearch.CreateConnection(datamodels.ConfRedisearch{
			Host: "192.168.9.208",
			Port: 6379,
		})
	})

	Context("Тест 1. Проверка подключения к БД Redisearch", func() {
		It("При подключении к БД Redisearch ошибки быть не должно", func() {
			Expect(rErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка наличия определенного количества документов", func() {
		It("Должно быть найдено документов больше 0", func() {
			docList, docNum, err := conn.Search(redisearch.NewQuery("*").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				).
				SetReturnFields("id"))

			fmt.Printf("______FULL SEARCH DOCUMENTS docNum: %d\n docList: %v\n", docNum, docList)

			Expect(docNum).Should(BeNumerically(">", 0))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Проверка наличия индексов", func() {
		It("Должно быть найдено некоторое кол-во индексов", func() {
			listIndex, err := conn.List()

			fmt.Println("listIndex = ", listIndex)

			Expect(listIndex[0]).Should(Equal("isems-index"))
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
