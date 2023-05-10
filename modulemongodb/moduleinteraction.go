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

// MongoDBChannels содержит каналы для взаимодействия с базой данных MongoDB
type MongoDBChannels struct {
	ChanInput  chan datamodels.ChannelsDescriptionInput
	ChanOutput chan datamodels.ChannelsDescriptionOutput
	ChanDown   chan struct{}
}

func IntarctionMongoDB(conf *datamodels.ConfMongoDB, currentLog *logging.LoggingData) (MongoDBChannels, error) {
	fmt.Println("func 'intarctionMongoDB' START")

	channels := MongoDBChannels{
		ChanInput:  make(chan datamodels.ChannelsDescriptionInput),
		ChanOutput: make(chan datamodels.ChannelsDescriptionOutput),
		ChanDown:   make(chan struct{}),
	}

	client, err := CreateConnection(conf)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		return channels, err
	}

	collection := client.Database(conf.DBname).Collection("stix_object_collection")

	go routing(channels.ChanOutput, collection, currentLog, channels.ChanDown, channels.ChanInput)

	return channels, nil
}

func CreateConnection(mdbs *datamodels.ConfMongoDB) (*mongo.Client, error) {
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
	chanOutput chan<- datamodels.ChannelsDescriptionOutput,
	collection *mongo.Collection,
	currentLog *logging.LoggingData,
	chanDown <-chan struct{},
	chanInput <-chan datamodels.ChannelsDescriptionInput) {
	fmt.Println("func 'routing' START")

	for {
		select {
		case req := <-chanInput:
			fmt.Println("func 'routing', REQUEST for mongo database: ", req)
		case <-chanDown:
			fmt.Println("func 'routing', reseived STOP signal")

			return
		}
	}
}
