package config

import (
	"os"

	"github.com/joho/godotenv"
)

// структура конфигов, а также их считывание из файла, и загрузка в эти стуктуры.

type MainConfig struct {
	Server     string
	External   string
	PostgresDB PostgresConfig
}

type PostgresConfig struct {
	User   string
	Pass   string
	Host   string
	Name   string
	Reload bool
	Port   string
}

func (cfg *MainConfig) ConfigMustLoad(name string) {

	path := "./config/" + name + ".yaml"
	godotenv.Load(path)

	cfg.PostgresDB.User := os.Getenv("PostgresUser")
	cfg.PostgresDB.Pass := os.Getenv("PostgresPass")
	cfg.PostgresDB.Host := os.Getenv("PostgresHost")
	cfg.PostgresDB.Port := os.Getenv("PostgresPort")
	cfg.PostgresDB.Name := os.Getenv("PostgresName")
	cfg.PostgresDB.Reload := os.Getenv("PostgresReload")
	cfg.Server := os.Getenv("ServerPort")
	cfg.External := os.Getenv("ExtAddress")
}
