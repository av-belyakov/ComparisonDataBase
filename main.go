package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/av-belyakov/comparisondatabase/datamodels"
	"github.com/av-belyakov/comparisondatabase/logging"
	"github.com/av-belyakov/comparisondatabase/modulecomparison"
	"github.com/av-belyakov/comparisondatabase/modulemongodb"
	"github.com/av-belyakov/comparisondatabase/moduleredisearch"
)

const strHelp = `application option:
	Параметры для подключения к MongoDB:
	 -mhost	ip адрес или доменное имя
	 -mport	сетевой порт
	 -mnamedb	наименование базы данных
	 -mcollection	коллекция
	 -muser	имя пользователя
	 -mpasswd	пароль пользователя
	Параметры для подключения к Redis DB:
	 -rhost	ip адрес или доменное имя
	 -rport	сетевой порт
	
	При использовании конфигурационного файла для подключения к СУБД (формат YAML):
	 -config имя файла
	Пример конфигурационного файла в формате YAML:
	
	MongoDB:
	 host: 127.0.0.1
	 port: 27017
	 namedb: nameDataBase
	 mcollection: collectionDataBase
	 user: userName
	 passwd: userPasswd
	Redisearch:
	 host: 127.0.0.1
	 port: 6379`

func readFileYAML(fileName string, conf datamodels.AppConfiguration) (datamodels.AppConfiguration, error) {
	if _, err := os.Stat(fileName); err != nil {
		return conf, err
	}

	viper.SetConfigFile(fileName)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return conf, err
	}

	if viper.IsSet("MongoDB.host") {
		conf.MongoDB.Host = viper.GetString("MongoDB.host")
	}

	if viper.IsSet("MongoDB.port") {
		conf.MongoDB.Port = viper.GetInt("MongoDB.port")
	}

	if viper.IsSet("MongoDB.namedb") {
		conf.MongoDB.DBname = viper.GetString("MongoDB.namedb")
	}

	if viper.IsSet("MongoDB.collection") {
		conf.MongoDB.Collection = viper.GetString("MongoDB.collection")
	}

	if viper.IsSet("MongoDB.user") {
		conf.MongoDB.User = viper.GetString("MongoDB.user")
	}

	if viper.IsSet("MongoDB.passwd") {
		conf.MongoDB.Passwd = viper.GetString("MongoDB.passwd")
	}

	if viper.IsSet("Redisearch.host") {
		conf.Redisearch.Host = viper.GetString("Redisearch.host")
	}

	if viper.IsSet("Redisearch.port") {
		conf.Redisearch.Port = viper.GetInt("Redisearch.port")
	}

	return conf, nil
}

func NewAppConf() (datamodels.AppConfiguration, error) {
	const (
		mhStr       = "ip адрес или доменное имя СУБД MongoDB"
		mpStr       = "сетевой порт СУБД MongoDB"
		fileNameStr = "конфигурационный файл приложения"
		errMsg      = "'-h' or '--help' to view the available options"
	)

	var fileName string
	conf := datamodels.AppConfiguration{
		MongoDB: datamodels.ConfMongoDB{
			Host: "127.0.0.1",
			Port: 27017,
		},
		Redisearch: datamodels.ConfRedisearch{
			Host: "127.0.0.1",
			Port: 6379,
		},
	}

	if len(os.Args) == 1 {
		return conf, errors.New(errMsg)
	}

	fs := flag.NewFlagSet(errMsg, flag.ContinueOnError)
	fs.StringVar(&fileName, "conf", "", fileNameStr)
	fs.StringVar(&conf.MongoDB.Host, "mhost", conf.MongoDB.Host, "ip адрес или доменное имя СУБД MongoDB")
	fs.IntVar(&conf.MongoDB.Port, "mport", conf.MongoDB.Port, "сетевой порт СУБД MongoDB")
	fs.StringVar(&conf.MongoDB.DBname, "mnamedb", conf.MongoDB.DBname, "имя базы данных для подключения к БД MongoDB")
	fs.StringVar(&conf.MongoDB.Collection, "mcollection", conf.MongoDB.Collection, "имя коллекции БД MongoDB")
	fs.StringVar(&conf.MongoDB.User, "muser", conf.MongoDB.User, "имя пользователя для подключения к БД MongoDB")
	fs.StringVar(&conf.MongoDB.Passwd, "mpasswd", conf.MongoDB.Passwd, "пароль для подключения к БД MongoDB")
	fs.StringVar(&conf.Redisearch.Host, "rhost", conf.Redisearch.Host, "ip адрес или доменное имя СУБД Redis")
	fs.IntVar(&conf.Redisearch.Port, "rport", conf.Redisearch.Port, "сетевой порт для подключения СУБД Redis")

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
	}

	if fileName != "" {
		return readFileYAML(fileName, conf)
	}

	return conf, nil
}

var (
	appConf    datamodels.AppConfiguration
	currentLog logging.LoggingData
)

func init() {
	var err error
	const errMsg = "the '%s' argument is missing to launch the application"

	if currentLog, err = logging.NewLoggingData("logs_app", "", []string{"error", "information"}); err != nil {
		log.Fatal(err)
	}

	for _, v := range os.Args {
		if v == "-h" || v == "--help" {
			fmt.Print(strHelp)

			return
		}
	}

	appConf, err = NewAppConf()
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")
		log.Fatal(err)
	}

	if appConf.MongoDB.DBname == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "data base name"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать наименование базы данных в MongoDB"))
	}

	if appConf.MongoDB.Collection == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "data base collection"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать коллекцию в базе данных в MongoDB"))
	}

	if appConf.MongoDB.User == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "user name"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать имя пользователя для доступа к MongoDB"))
	}

	if appConf.MongoDB.Passwd == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "password"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать пароль пользователя для доступа к MongoDB"))
	}
}

func main() {
	fmt.Println("Comparison data base application is START")

	currentLog.WriteLoggingData("start comparisonDataBase application", "information")

	//инициализация соединения с MongoDB
	mdbChan, err := modulemongodb.IntarctionMongoDB(&appConf.MongoDB, &currentLog)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatal(err)
	}

	//инициализация соединения с Redisearch
	rsdbChan, err := moduleredisearch.InteractionRedisearch(&appConf.Redisearch, &currentLog)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatal(err)
	}

	fmt.Println("action channals MongoDB: ", mdbChan, " send request to MongoDB")
	mdbChan.ChanInput <- datamodels.ChannelInputMDB{
		ActionType: "test request",
		Data:       "any data",
	}

	modulecomparison.ModuleInteraction(&mdbChan, &rsdbChan, &currentLog)

	//	mdbchan.ChanDown <- struct{}{}
	//	time.Sleep(3 * time.Second)

	currentLog.ClosingFiles()
}
