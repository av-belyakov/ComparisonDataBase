package datamodels

import (
	"encoding/json"
)

// CommonPropertiesObjectSTIX свойства общие, для всех объектов STIX
// Type - наименование типа шаблона (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//
//	Type должен содержать одно из следующих значений:
//	1. Для Domain Objects STIX
//
// - "attack-pattern"
// - "campaign"
// - "course-of-action"
// - "grouping"
// - "identity"
// - "indicator"
// - "infrastructure"
// - "intrusion-set"
// - "location"
// - "malware"
// - "malware-analysis"
// - "note"
// - "observed-data"
// - "opinion"
// - "report"
// - "threat-actor"
// - "tool"
// - "vulnerability"
//  2. Для Relationship Objects STIX
//
// - "relationship"
// - "sighting"
//  3. Для Cyber-observable Objects STIX
//
// - "artifact"
// - "autonomous-system"
// - "directory"
// - "domain-name"
// - "email-addr"
// - "email-message"
// - "file"
// - "ipv4-addr"
// - "ipv6-addr"
// - "mac-addr"
// - "mutex"
// - "network-traffic"
// - "process"
// - "software"
// - "url"
// - "user-account"
// - "windows-registry-key"
// - "x509-certificate"
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type CommonPropertiesObjectSTIX struct {
	Type string `bson:"type"`
	ID   string `bson:"id"`
}

// HandlerSTIXObject интерфейс реализующий обработчики для STIX объектов
type HandlerSTIXObject interface {
	DecoderJSONObject
	EncoderJSONObject
	GetterParametersSTIXObject
	ToBeautifulOutputConverter
	IndexingSTIXObject
}

// DecoderJSONObject интерфейс реализующий обработчик для декодирования JSON объекта в STIX объект
type DecoderJSONObject interface {
	DecodeJSON(*json.RawMessage) (interface{}, error)
}

// EncoderJSONObject интерфейс реализующий обработчик для кодирования STIX объекта в JSON объект
type EncoderJSONObject interface {
	EncodeJSON(interface{}) (*[]byte, error)
}

// GetterParametersSTIXObject интерфейс реализующий обработчик для получения ID STIX объекта
type GetterParametersSTIXObject interface {
	GetID() string
	GetType() string
}

// ToBeautifulOutputConverter интерфейс реализующий обработчик для красивого представления данных хранящихся в пользовательской структуре
type ToBeautifulOutputConverter interface {
	ToStringBeautiful() string
}

type IndexingSTIXObject interface {
	GeneratingDataForIndexing() map[string]string
}

// ElementSTIXObject может содержать любой из STIX объектов с указанием его типа
// DataType - тип STIX объекта
// Data - непосредственно сам STIX объект
type ElementSTIXObject struct {
	DataType string
	Data     HandlerSTIXObject
}
