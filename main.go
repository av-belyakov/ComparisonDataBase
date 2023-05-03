package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	"ComparisonDataBase/datamodels"
	"ComparisonDataBase/logging"
	"ComparisonDataBase/modulemongodb"
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
	Redis:
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

	if viper.IsSet("MongoDB.mcollection") {
		conf.MongoDB.Collection = viper.GetString("MongoDB.mcollection")
	}

	if viper.IsSet("MongoDB.user") {
		conf.MongoDB.User = viper.GetString("MongoDB.user")
	}

	if viper.IsSet("MongoDB.passwd") {
		conf.MongoDB.Passwd = viper.GetString("MongoDB.passwd")
	}

	if viper.IsSet("Redis.host") {
		conf.Redis.Host = viper.GetString("Redis.host")
	}

	if viper.IsSet("Redis.port") {
		conf.Redis.Port = viper.GetInt("Redis.port")
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
		Redis: datamodels.ConfRedis{
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
	fs.StringVar(&conf.Redis.Host, "rhost", conf.Redis.Host, "ip адрес или доменное имя СУБД Redis")
	fs.IntVar(&conf.Redis.Port, "rport", conf.Redis.Port, "сетевой порт для подключения СУБД Redis")

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
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "dbname"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать наименование базы данных в MongoDB"))
	}

	if appConf.MongoDB.Collection == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "dbcollection"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать коллекцию в базе данных в MongoDB"))
	}

	if appConf.MongoDB.User == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "user name"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать имя пользователя для доступа к MongoDB"))
	}

	if appConf.MongoDB.Passwd == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "passwd"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать пароль пользователя для доступа к MongoDB"))
	}
}

func main() {
	fmt.Println("comparisonDataBase application is START")
	fmt.Println("___ appConf = ", appConf)

	currentLog.WriteLoggingData("start comparisonDataBase application", "information")

	mdbchan, err := modulemongodb.IntarctionMongoDB(&appConf.MongoDB, &currentLog)
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")
	}

	fmt.Println("action channals MongoDB: ", mdbchan, " send request to MongoDB")
	mdbchan.ChanInput <- modulemongodb.ChanOption{
		ActionType: "test request",
		Data:       "any data",
	}

	fmt.Println("send data to chan STOP")
	mdbchan.ChanDown <- struct{}{}

	currentLog.ClosingFiles()
}
