package services

import (
	"mymod/internal/database"
	"mymod/internal/database/dao"
	"mymod/internal/models"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао.

func SongInsert(instance tables.Song) models.Response {
	log.Debug("services get = ", instance)

	EnrichtSong()

	return instance.RecordCreate(database.GlobalPostgres, dao.GlobalSongDao)
}

func SongShow(instance tables.Song, limit int, offset int) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordShow(database.GlobalPostgres, dao.GlobalSongDao, limit, offset)
}

func SongUpdate(instance tables.Song) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordUpdate(database.GlobalPostgres, dao.GlobalSongDao)
}

func SongDelete(instance tables.Song) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordDelete(database.GlobalPostgres, dao.GlobalSongDao)
}
