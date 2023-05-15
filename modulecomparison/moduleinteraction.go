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

func ModuleInteraction(mdbChan *modulemongodb.MongoDBChannels, rsdbChan *moduleredisearch.RedisearchChannels, currentLog *logging.LoggingData) {
	fmt.Println("func 'ModuleInteraction', START...")

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

	fmt.Println("All search object is equal: ", fullCountObj)

	// получаем кол-во частей
	maxChunk := commonlibs.GetCountChunk(fullCountObj, limitMaxSize)

	fmt.Println("func 'ModuleInteraction', maxChank = ", maxChunk)

	mdbChan.ChanInput <- datamodels.ChannelInputMDB{
		ActionType:   "get a limited number of objects",
		Offset:       int64(offset),
		LimitMaxSize: limitMaxSize,
	}

	for data := range mdbChan.ChanOutput {
		fmt.Printf("func 'ModuleInteraction', RESEIVED data from MongoDB: %v\n", data)

		if offset == maxChunk {
			fmt.Println("func 'ModuleInteraction', STOP FOR")

			return
		}

		offset += 1
		mdbChan.ChanInput <- datamodels.ChannelInputMDB{
			ActionType:   "get a limited number of objects",
			Offset:       int64(offset),
			LimitMaxSize: limitMaxSize,
		}
	}

	/*
		Сделал что бы генерировались запросы на получение всех кусочков, однако приложение завершает работу (и скорее всего
		раньше чем будут добавлены данные в Redisearch) потому что блокирующая функция завершает цикл в связи с закрытием
		канала. Этот момент надо продумать, что бы приложение не завершало работу до того как данные не будут гарантированно
		добавлены в Redisearch

	*/

	//? если Redisearch ничего не будет возвращать (вроде не должна так как стороит только индексы) то достаточно этого
	/*
		for {
			select {
			case data := <-mdbChan.ChanOutput:
				fmt.Printf("func 'ModuleInteraction', RESEIVED data from MongoDB: %v\n", data)

			}
		}
		/*
			Функция должна принимать два канала для взаимодействия с модулями MongoDB и Redisearch, кроме
			того метод для логирования хода выполнения

			1. Получить общее кол-во данных в MongoDB
			1.1 Вывести в консоль полученные данные
			2. Разделить кол-во полученных данных на сегменты ?
			3. Запрос сегмента
			4. Передача данных, через канал, в модуль ответственный за взаимодействие с Redisearch
	*/
}
