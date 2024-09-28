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

func main() {
	log.SetLevel(log.DebugLevel)
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
