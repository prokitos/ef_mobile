package dao

import (
	"mymod/internal/database"
	"mymod/internal/models"
)

// конвертация интерфейса базы данных в конкретную базу данных
func convertToPostgres(interf models.DatabaseCore) *database.PostgresDatabase {
	dbConnect, err := interf.(*database.PostgresDatabase)
	if !err {
		return nil
	}

	return dbConnect
}
