package dao

import (
	"modules/internal/database"
	"modules/internal/models"
)

// Глобальные переменные которые хранят все существующие дао, а также конвертация интерфейса подключения к базе данных в конкретное подключение.

var GlobalUserDao *UserDao
var GlobalTaskDao *TaskDao
var GlobalComputerDao *ComputerDao
var GlobalCarDao *CarDao
var GlobalProductDao *ProductDao

// конвертация интерфейса базы данных в конкретную базу данных
func convertToMongo(interf models.DatabaseCore) *database.MongoDatabase {
	dbConnect, err := interf.(*database.MongoDatabase)
	if !err {
		return nil
	}

	return dbConnect
}

// конвертация интерфейса базы данных в конкретную базу данных
func convertToPostgres(interf models.DatabaseCore) *database.PostgresDatabase {
	dbConnect, err := interf.(*database.PostgresDatabase)
	if !err {
		return nil
	}

	return dbConnect
}

// конвертация интерфейса базы данных в конкретную базу данных
func convertToSqlite(interf models.DatabaseCore) *database.SqliteDatabase {
	dbConnect, err := interf.(*database.SqliteDatabase)
	if !err {
		return nil
	}

	return dbConnect
}
