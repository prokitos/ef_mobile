package main

import (
	"mymod/internal/app"
	"mymod/internal/config"
	"mymod/internal/database"
	"mymod/internal/services"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// запуск логов; загрузка конфигов; запуск бд и сервера.

// @title Test API
// @version 1.0
// @description This is a sample service for managing songs
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @BasePath /
func main() {
	log.SetLevel(log.DebugLevel) // чтобы отображать debug логи
	//log.SetLevel(log.InfoLevel)  // чтобы не отображать debug логи
	log.SetFormatter(&log.JSONFormatter{}) // для json формата, чтобы потом удобно было логи выгружать в файл.
	//enableLogToFile()                      // записывать логи в файл
	log.Info("log is loaded")

	var cfg config.MainConfig
	cfg.ConfigMustLoad("local")
	services.SongExternalAddress = cfg.External
	log.Info("config is loaded")

	var PGDB database.PostgresDatabase
	PGDB.Run(cfg)
	log.Info("database is loaded")

	var application app.App
	go application.NewServer(cfg.Server)
	log.Info("server is loaded")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop()
}

func enableLogToFile() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}
