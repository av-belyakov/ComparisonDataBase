package datamodels

// IndexObject объект индексирования
// ID - идентификатор
// Type - тип объекта
// Name - наименование
// Description - подробное описание
// StreetAddress - физический адрес
// Abstract - краткое изложение содержания записки используется в STIX объектах Node
// Aliases - альтернативные имена
// Content - основное содержание записки используется в STIX объектах Node
// Value - параметр value может содержать в себе сетевое доменное имя, email адрес, ip адрес, url в STIX объектах DomainName,
// EmailAddress, IPv4Address, IPv6Address, URL
type IndexObject struct {
	ID            string `bson:"id"`
	Type          string `bson:"type"`
	Name          string `bson:"name"`
	Description   string `bson:"description"`
	StreetAddress string `bson:"street_address"`
	Abstract      string `bson:"abstract"`
	Aliases       string `bson:"aliases"`
	Content       string `bson:"content"`
	Value         string `bson:"value"`
}
