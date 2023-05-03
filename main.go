package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"comparison_data_two_DB/logging"
)

func readFileYAML(conf Configuration) (Configuration, error) {

	return conf, nil
}

type ConfMongoDB struct {
	host   string
	port   int
	dbname string
	user   string
	passwd string
}

type ConfReactDB struct {
	host string
	port int
}

type Configuration struct {
	MongoDB ConfMongoDB
	ReactDB ConfReactDB
}

func NewAppConf() (*Configuration, error) {
	const strHelp = `application option:
	Параметры для подключения к MongoDB:
	 -mh	--mhost	ip адрес или доменное имя
	 -mp	--mport	сетевой порт
	 -mndb	--mnamedb	наименование базы данных
	 -mu	--muser	имя пользователя
	 -mpass	--mpasswd	пароль пользователя
	Параметры для подключения к Redis DB:
	 -rh	--rhost	ip адрес или доменное имя
	 -rp	--rport	сетевой порт
	
	 При использовании конфигурационного файла для подключения к СУБД (формат YAML):
	 -c		--config имя файла
	Пример конфигурационного файла в формате YAML:
	
	MongoDB:
	 host: 127.0.0.1
	 port: 27017
	 namedb: nameDataBase
	 user: userName
	 passwd: userPasswd
	Redis:
	 host: 127.0.0.1
	 port: 6379`

	const (
		mhStr = "ip адрес или доменное имя СУБД MongoDB"
		mpStr = "сетевой порт СУБД MongoDB"
	)

	conf := &Configuration{
		MongoDB: ConfMongoDB{
			host: "127.0.0.1",
			port: 27017,
		},
		ReactDB: ConfReactDB{
			host: "127.0.0.1",
			port: 6379,
		},
	}

	if len(os.Args) == 1 {
		fmt.Println("Недостаточно аргументов для запуска приложения. Добавьте '-h' или '--help' для просмотра доступных параметров.")

		return conf, errors.New("there are not enough arguments to launch the application")
	}

	for _, v := range os.Args {
		if v == "-h" || v == "--help" {
			fmt.Print(strHelp)

			return conf, nil
		}

		if v == "-c" || v == "--conf" {

			return conf, nil
		}
	}

	fs := flag.NewFlagSet(strHelp, flag.ContinueOnError)

	fs.StringVar(&conf.MongoDB.host, "mh", conf.MongoDB.host, mhStr)
	fs.StringVar(&conf.MongoDB.host, "mhost", conf.MongoDB.host, mhStr)
	fs.IntVar(&conf.MongoDB.port, "mp", conf.MongoDB.port, mpStr)
	fs.IntVar(&conf.MongoDB.port, "mport", conf.MongoDB.port, mpStr)
	fs.StringVar(&conf.MongoDB.dbname, "mndb", conf.MongoDB.dbname, mhStr)
	fs.StringVar(&conf.MongoDB.dbname, "mnamedb", conf.MongoDB.dbname, mhStr)
	fs.StringVar(&conf.MongoDB.user, "mu", conf.MongoDB.user, mhStr)
	fs.StringVar(&conf.MongoDB.user, "muser", conf.MongoDB.user, mhStr)
	fs.StringVar(&conf.MongoDB.passwd, "mpass", conf.MongoDB.passwd, mhStr)
	fs.StringVar(&conf.MongoDB.passwd, "mpasswd", conf.MongoDB.passwd, mhStr)
	fs.StringVar(&conf.ReactDB.host, "rh", conf.ReactDB.host, mhStr)
	fs.StringVar(&conf.ReactDB.host, "rhost", conf.ReactDB.host, mhStr)
	fs.IntVar(&conf.ReactDB.port, "rp", conf.ReactDB.port, mpStr)
	fs.IntVar(&conf.ReactDB.port, "rport", conf.ReactDB.port, mpStr)

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
	}

	return conf, nil
}

var (
	appConf    *Configuration
	currentLog logging.LoggingData
)

func init() {
	const errMsg = "the '%s' argument is missing to launch the application"
	var err error
	if currentLog, err = logging.NewLoggingData("logs_app", "", []string{"error", "information"}); err != nil {
		log.Fatal(err)
	}

	appConf, err = NewAppConf()
	if err != nil {
		currentLog.WriteLoggingData(fmt.Sprint(err), "error")
	}

	if appConf.MongoDB.dbname == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "namedb"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать наименование базы данных в MongoDB"))
	}

	if appConf.MongoDB.user == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "user name"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать имя пользователя для доступа к MongoDB"))
	}

	if appConf.MongoDB.passwd == "" {
		currentLog.WriteLoggingData(fmt.Sprintf(errMsg, "passwd"), "error")
		log.Fatal(errors.New("для запуска приложения необходимо указать пароль пользователя для доступа к MongoDB"))
	}
}

func main() {
	fmt.Println("comparisonDataBase application is START")
	fmt.Println("___ appConf = ", appConf)

	currentLog.WriteLoggingData("start comparisonDataBase application", "information")

	currentLog.ClosingFiles()
}
