package moduleredisearch

import (
	"fmt"
	"sync"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/logging"
)

// RedisDBChannels содержит каналы для в заимодействия с базой данных Redis
type RedisearchChannels struct {
	ChanInput  chan datamodels.ChannelInputRSDB
	ChanOutput chan datamodels.ChannelOutputRSDB
	ChanDown   chan struct{}
}

func getRedisearchDocument(listIndex []datamodels.IndexObject) []redisearch.Document {
	redisearchDoc := make([]redisearch.Document, 0, len(listIndex))

	for _, v := range listIndex {
		tmp := redisearch.NewDocument(v.ID, 1.0)
		tmp.Set("type", v.Type)
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

func InteractionRedisearch(conf *datamodels.ConfRedisearch, currentLog *logging.LoggingData, wg *sync.WaitGroup) (RedisearchChannels, error) {
	channels := RedisearchChannels{
		ChanInput:  make(chan datamodels.ChannelInputRSDB),
		ChanOutput: make(chan datamodels.ChannelOutputRSDB),
		ChanDown:   make(chan struct{}),
	}

	conn, err := CreateConnection(*conf)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		return channels, err
	}

	go routing(channels.ChanOutput, conn, currentLog, channels.ChanInput)
	go func() {
		<-channels.ChanDown
		fmt.Println("func InteractionRedisearch, groutina CLOSE channels ChanInput and ChanOutput")

		close(channels.ChanInput)
		close(channels.ChanOutput)

		wg.Done()
	}()

	return channels, nil
}

func CreateConnection(conf datamodels.ConfRedisearch) (*redisearch.Client, error) {
	conn := redisearch.NewClient(fmt.Sprintf("%v:%v", conf.Host, conf.Port), "isems-index")
	if _, err := conn.Info(); err != nil {
		return conn, err
	}

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("type")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("description")).
		//физический адрес
		AddField(redisearch.NewTextField("street_address")).
		//результат классификации или имя, присвоенное экземпляру вредоносного ПО инструментом анализа (сканером)
		// используется в STIX объектах MalwareAnalysis
		AddField(redisearch.NewTextField("result_name")).
		//краткое изложение содержания записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("abstract")).
		//основное содержание записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("content")).
		AddField(redisearch.NewTextField("url")).
		//параметр value может содержать в себе сетевое доменное имя,
		// email адрес, ip адрес, url в STIX объектах DomainName, EmailAddress,
		// IPv4Address, IPv6Address, URL
		AddField(redisearch.NewTextField("value"))

	_ = conn.CreateIndex(sc)

	return conn, nil
}

func routing(
	chanOutput chan<- datamodels.ChannelOutputRSDB,
	conn *redisearch.Client,
	currentLog *logging.LoggingData,
	chanInput <-chan datamodels.ChannelInputRSDB) {

	for req := range chanInput {
		switch req.ActionType {
		case "set index":
			if err := conn.IndexOptions(
				redisearch.IndexingOptions{
					Replace: true,
					Partial: true,
				}, getRedisearchDocument(req.IndexList)...); err != nil {
				currentLog.WriteLoggingData(fmt.Sprint(err), "error")
			}

		case "get count index":
			_, docNum, err := conn.Search(redisearch.NewQuery("*").
				AddFilter(
					redisearch.Filter{
						Field: "name",
					},
				).
				SetReturnFields("id"))

			if err != nil {
				currentLog.WriteLoggingData(fmt.Sprint(err), "error")
			}

			chanOutput <- datamodels.ChannelOutputRSDB{
				DataType:   "send count index",
				IndexCount: docNum,
			}
		}
	}
}
