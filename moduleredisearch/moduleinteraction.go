package moduleredisearch

import (
	"fmt"

	"github.com/RediSearch/redisearch-go/redisearch"

	"ComparisonDataBase/datamodels"
	"ComparisonDataBase/logging"
)

// RedisDBChannels содержит каналы для в заимодействия с базой данных Redis
type RedisearchChannels struct {
	ChanInput  chan datamodels.ChannelsDescriptionInput
	ChanOutput chan datamodels.ChannelsDescriptionOutput
	ChanDown   chan struct{}
}

func InteractionRedisearch(conf *datamodels.ConfRedisearch, currentLog *logging.LoggingData) (RedisearchChannels, error) {
	fmt.Println("func 'InteractionRedisearch' START")

	channels := RedisearchChannels{
		ChanInput:  make(chan datamodels.ChannelsDescriptionInput),
		ChanOutput: make(chan datamodels.ChannelsDescriptionOutput),
		ChanDown:   make(chan struct{}),
	}

	//client, err := CreateConnection(conf)
	conn, err := CreateConnection(*conf)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		return channels, err
	}

	go routing(channels.ChanOutput, conn, currentLog, channels.ChanDown, channels.ChanInput)

	return channels, nil
}

func CreateConnection(conf datamodels.ConfRedisearch) (*redisearch.Client, error) {
	fmt.Println("func 'CreateConnection', Redisearch, START...")

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

	/*if err := conn.CreateIndex(sc); err == nil {
		return conn, err
	}

	return conn, fmt.Errorf("error connecting to the Research database or error creating indexes")*/

	_ = conn.CreateIndex(sc)

	return conn, nil
}

func routing(
	chanOutput chan<- datamodels.ChannelsDescriptionOutput,
	conn *redisearch.Client,
	currentLog *logging.LoggingData,
	chanDown <-chan struct{},
	chanInput <-chan datamodels.ChannelsDescriptionInput) {
	fmt.Println("func 'routing' START")

	for {
		select {
		case req := <-chanInput:
			fmt.Println("func 'routing', REQUEST for redisearch database: ", req)
		case <-chanDown:
			fmt.Println("func 'routing', reseived STOP signal")

			return
		}
	}
}
