package modulemongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"ComparisonDataBase/datamodels"
	"ComparisonDataBase/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ChanOption struct {
	ActionType string
	Data       interface{}
}

type MongoDBChannels struct {
	ChanInput, ChanOutput chan ChanOption
	ChanDown              chan struct{}
}

func IntarctionMongoDB(conf *datamodels.ConfMongoDB, currentLog *logging.LoggingData) (MongoDBChannels, error) {
	fmt.Println("func 'intarctionMongoDB' START")

	channels := MongoDBChannels{
		ChanInput:  make(chan ChanOption),
		ChanOutput: make(chan ChanOption),
		ChanDown:   make(chan struct{}),
	}

	// подключаемся к базе данных MongoDB
	mclient, err := createConnection(conf)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		return channels, err
	}

	// инициализируем маршрутизатор запросов
	go routing(channels.ChanOutput, mclient, currentLog, channels.ChanDown, channels.ChanInput)

	return channels, err
}

func createConnection(mdbs *datamodels.ConfMongoDB) (*mongo.Client, error) {
	fmt.Println("func 'createConnection' START")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 3000*time.Second)
	defer func() {
		ctxCancel()
	}()

	clientOption := options.Client()
	clientOption.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    mdbs.DBname,
		Username:      mdbs.User,
		Password:      mdbs.Passwd,
	})

	client, err := mongo.NewClient(clientOption.ApplyURI("mongodb://" + mdbs.Host + ":" + strconv.Itoa(mdbs.Port) + "/" + mdbs.DBname))
	if err != nil {
		return nil, err
	}

	client.Connect(ctx)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	fmt.Println("func 'createConnection' END")

	return client, nil
}

func routing(
	chanOutput chan<- ChanOption,
	mclient *mongo.Client,
	currentLog *logging.LoggingData,
	chanDown <-chan struct{},
	chanInput <-chan ChanOption) {
	fmt.Println("func 'routing' START")

	for {
		select {
		case req := <-chanInput:
			fmt.Println("func 'routing', request for mongo database: ", req)
		case <-chanDown:
			return
		}
	}
}
