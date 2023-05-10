package datamodels

// ChannelsDescriptionOutput предназначен для вывода результатов поиска из БД
// DataType - тип данных
// Data - данные
type ChannelsDescriptionOutput struct {
	DataType string
	Data     interface{}
}

// ChannelsDescriptionInput предназначен для передачи запросов в БД
// ActionType - тип запроса
// Data - данные
type ChannelsDescriptionInput struct {
	ActionType string
	Data       interface{}
}
