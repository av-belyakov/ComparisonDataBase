package datamodels

// ConfMongoDB хранит настройки СУБД MongoDB
// Host - ip адрес или доменное имя
// Port - сетевой порт
// DBname - наименование базы данных
// User - имя пользователя
// Passwd - пароль пользователя
type ConfMongoDB struct {
	Host       string
	Port       int
	DBname     string
	Collection string
	User       string
	Passwd     string
}

// ConfRedis хранит настройки СУБД Redis
// Host - ip адрес или доменное имя
// Port - сетевой порт
type ConfRedis struct {
	Host string
	Port int
}

// AppConfiguration хранит настройки приложения
type AppConfiguration struct {
	MongoDB ConfMongoDB
	Redis   ConfRedis
}
