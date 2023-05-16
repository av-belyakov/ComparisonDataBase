package modulecomparison

import (
	"fmt"

	"github.com/av-belyakov/comparisondatabase/commonlibs"
	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/logging"
	"github.com/av-belyakov/comparisondatabase/modulemongodb"
	"github.com/av-belyakov/comparisondatabase/moduleredisearch"
)

const limitMaxSize = 150

func getListIndex(listElem []datamodels.ElementSTIXObject) []datamodels.IndexObject {
	listIndexObj := []datamodels.IndexObject{}

	for _, v := range listElem {
		listIndexObj = append(listIndexObj, moduleredisearch.GetIndex(v))
	}

	return listIndexObj
}

func ModuleInteraction(mdbChan *modulemongodb.MongoDBChannels, rsdbChan *moduleredisearch.RedisearchChannels, currentLog *logging.LoggingData) {
	var fullCountObj int64
	var offset int = 1

	//запрос на общее колличество STIX объектов в MongoDB
	mdbChan.ChanInput <- datamodels.ChannelInputMDB{ActionType: "get count object"}

	// принимаем найденное общее кол-во объектов
	data := <-mdbChan.ChanOutput
	fullCountObj, ok := data.Data.(int64)
	if !ok {
		currentLog.WriteLoggingData("невозможно выполнить преобразование типа в int64 для общего количества найденных STIX объектов", "error")

		mdbChan.ChanDown <- struct{}{}
		rsdbChan.ChanDown <- struct{}{}

		return
	}

	if fullCountObj == 0 {
		mdbChan.ChanDown <- struct{}{}
		rsdbChan.ChanDown <- struct{}{}

		fmt.Println("0 objects found, stopped")

		return
	}

	// получаем кол-во частей
	maxChunk := commonlibs.GetCountChunk(fullCountObj, limitMaxSize)

	fmt.Printf("Total objects found %d, parts of them %d\n", fullCountObj, maxChunk)

	mdbChan.ChanInput <- datamodels.ChannelInputMDB{
		ActionType:   "get a limited number of objects",
		Offset:       int64(offset),
		LimitMaxSize: limitMaxSize,
	}

	for data := range mdbChan.ChanOutput {
		listElem, ok := data.Data.([]datamodels.ElementSTIXObject)
		if ok {
			fmt.Println("Creating indexes for objects of part no. ", offset)

			rsdbChan.ChanInput <- datamodels.ChannelInputRSDB{
				ActionType: "set index",
				IndexList:  getListIndex(listElem),
			}
		}

		if offset == maxChunk {
			break
		}

		offset += 1
		mdbChan.ChanInput <- datamodels.ChannelInputMDB{
			ActionType:   "get a limited number of objects",
			Offset:       int64(offset),
			LimitMaxSize: limitMaxSize,
		}
	}

	mdbChan.ChanDown <- struct{}{}

	rsdbChan.ChanInput <- datamodels.ChannelInputRSDB{ActionType: "get count index"}
	tmp := <-rsdbChan.ChanOutput

	rsdbChan.ChanDown <- struct{}{}

	fmt.Println("Total indexes built ", tmp.IndexCount)
}
