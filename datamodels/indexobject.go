package datamodels

// IndexObject объект индексирования
// ID - идентификатор
// Type - тип объекта
// Name - наименование
// Description - подробное описание
// StreetAddress - физический адрес
// ResultName - результат классификации или имя, присвоенное экземпляру вредоносного ПО инструментом анализа (сканером)
// используется в STIX объектах MalwareAnalysis
// Abstract - краткое изложение содержания записки используется в STIX объектах Node
// Content - основное содержание записки используется в STIX объектах Node
// URL - уникальный идентификатор ресурса
// Value - параметр value может содержать в себе сетевое доменное имя, email адрес, ip адрес, url в STIX объектах DomainName,
// EmailAddress, IPv4Address, IPv6Address, URL
type IndexObject struct {
	ID            string `bson:"id"`
	Type          string `bson:"type"`
	Name          string `bson:"name"`
	Description   string `bson:"description"`
	StreetAddress string `bson:"street_address"`
	ResultName    string `bson:"result_name"`
	Abstract      string `bson:"abstract"`
	Content       string `bson:"content"`
	URL           string `bson:"url"`
	Value         string `bson:"value"`
}
