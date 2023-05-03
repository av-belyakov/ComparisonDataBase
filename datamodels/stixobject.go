package datamodels

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
	Type string `json:"type" bson:"type" required:"true"`
	ID   string `json:"id" bson:"id" required:"true"`
}
