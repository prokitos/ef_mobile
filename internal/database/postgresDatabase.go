package database

import (
	"fmt"
	"mymod/internal/config"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// запуск, соединение и миграция для postgresDB.

func (currentlDB *PostgresDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

func (currentlDB *PostgresDatabase) StartMigration() {
	currentlDB.Instance.AutoMigrate(tables.Song{})
	log.Debug("migration complete")
}

func (currentlDB *PostgresDatabase) OpenConnection(config config.MainConfig) {

	err := currentlDB.checkDatabaseCreated(config)
	if err != nil {
		panic("not connection to db")
	}

	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresDB.User, config.PostgresDB.Pass, config.PostgresDB.Host, config.PostgresDB.Port, config.PostgresDB.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db

}

// проверка если есть база данных. если нет, то создать.
func (currentlDB *PostgresDatabase) checkDatabaseCreated(config config.MainConfig) error {

	// открытие соеднение с базой по стандарту
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresDB.User, config.PostgresDB.Pass, config.PostgresDB.Host, config.PostgresDB.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Error("database don't open")
		return responses.ResponseBase{}.BaseServerError()
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	// проверка если есть нужная нам база данных
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.PostgresDB.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Error("error, dont read bd")
		return responses.ResponseBase{}.BaseServerError()
	}

	// если нет, то создать
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.PostgresDB.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Error("error, dont create a database")
			responses.ResponseBase{}.BaseServerError()
		}
	}

	return nil
}

func (currentlDB *PostgresDatabase) GlobalSet() {
	GlobalPostgres = currentlDB
}
