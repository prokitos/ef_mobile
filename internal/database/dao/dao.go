package dao

import (
	"mymod/internal/database"
	"mymod/internal/models"
)

// Глобальные переменные которые хранят все существующие дао, а также конвертация интерфейса подключения к базе данных в конкретное подключение.

var GlobalSongDao *SongDao

// конвертация интерфейса базы данных в конкретную базу данных
func convertToPostgres(interf models.DatabaseCore) *database.PostgresDatabase {
	dbConnect, err := interf.(*database.PostgresDatabase)
	if !err {
		return nil
	}

	return dbConnect
}
