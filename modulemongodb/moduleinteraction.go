package modulemongodb

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDBChannels содержит каналы для взаимодействия с базой данных MongoDB
type MongoDBChannels struct {
	ChanInput  chan datamodels.ChannelInputMDB
	ChanOutput chan datamodels.ChannelOutputMDB
	ChanDown   chan struct{}
}

func IntarctionMongoDB(conf *datamodels.ConfMongoDB, currentLog *logging.LoggingData, wg *sync.WaitGroup) (MongoDBChannels, error) {
	channels := MongoDBChannels{
		ChanInput:  make(chan datamodels.ChannelInputMDB),
		ChanOutput: make(chan datamodels.ChannelOutputMDB),
		ChanDown:   make(chan struct{}),
	}

	client, err := CreateConnection(conf)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		return channels, err
	}

	collection := client.Database(conf.DBname).Collection("stix_object_collection")

	go routing(channels.ChanOutput, collection, currentLog, channels.ChanInput)
	go func() {
		<-channels.ChanDown
		fmt.Println("func IntarctionMongoDB, groutina CLOSE channels ChanInput and ChanOutput")

		close(channels.ChanInput)
		close(channels.ChanOutput)

		wg.Done()
	}()

	return channels, nil
}

func CreateConnection(mdbs *datamodels.ConfMongoDB) (*mongo.Client, error) {
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

	return client, nil
}

func routing(
	chanOutput chan<- datamodels.ChannelOutputMDB,
	collection *mongo.Collection,
	currentLog *logging.LoggingData,
	chanInput <-chan datamodels.ChannelInputMDB) {

	for req := range chanInput {
		switch req.ActionType {
		case "get count object":
			result, err := wrapperRoutingGetFullCount(collection)
			if err != nil {
				currentLog.WriteLoggingData(fmt.Sprint(err), "error")
			}

			chanOutput <- datamodels.ChannelOutputMDB{
				DataType: "full count object",
				Data:     result,
			}

		case "get a limited number of objects":
			result, err := wrapperRoutingGetLimitObject(collection, req.Offset, req.LimitMaxSize)

			if err != nil {
				currentLog.WriteLoggingData(fmt.Sprint(err), "error")
			}

			chanOutput <- datamodels.ChannelOutputMDB{
				DataType: "limited number of object",
				Data:     result,
			}
		}
	}
	/*for {
		select {
		case req := <-chanInput:
			switch req.ActionType {
			case "get count object":
				result, err := wrapperRoutingGetFullCount(collection)
				if err != nil {
					currentLog.WriteLoggingData(fmt.Sprint(err), "error")
				}

				chanOutput <- datamodels.ChannelOutputMDB{
					DataType: "full count object",
					Data:     result,
				}

			case "get a limited number of objects":
				result, err := wrapperRoutingGetLimitObject(collection, req.Offset, req.LimitMaxSize)

				if err != nil {
					currentLog.WriteLoggingData(fmt.Sprint(err), "error")
				}

				chanOutput <- datamodels.ChannelOutputMDB{
					DataType: "limited number of object",
					Data:     result,
				}
			}
		case <-chanDown:
			fmt.Println("func 'routing' MongoDB, reseived STOP signal")

			close(chanOutput)

			return
		}
	}*/
}
