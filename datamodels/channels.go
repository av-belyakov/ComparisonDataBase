package datamodels

// ChannelsOutputMDB предназначен для вывода результатов поиска из БД MongoDB
// DataType - тип данных
// Data - данные
type ChannelOutputMDB struct {
	DataType string
	Data     interface{}
}

// ChannelOutputRSDB предназначен для передачи данных из БД Redisearch
// DataType - тип данных
// IndexCount - количество индексов
type ChannelOutputRSDB struct {
	DataType   string
	IndexCount int
}

// ChannelInputMDB предназначен для передачи запросов в БД MongoDB
// ActionType - тип запроса
// Offset - смещение
// LimitMaxSize - максимальный лимит для выборки
// Data - данные
type ChannelInputMDB struct {
	ActionType   string
	Offset       int64
	LimitMaxSize int64
	Data         interface{}
}

// ChannelInputRSDB предназначен для передачи запросов в БД Redisearch
// ActionType - тип запроса
// IndexList - список индексов
type ChannelInputRSDB struct {
	ActionType string
	IndexList  []IndexObject
}
