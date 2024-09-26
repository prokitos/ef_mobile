package database

import (
	"gorm.io/gorm"
)

// глобальные переменные для хранения всех подключений к бд.

var GlobalPostgres *PostgresDatabase

type PostgresDatabase struct {
	Instance *gorm.DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
