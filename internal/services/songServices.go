package services

import (
	"mymod/internal/database"
	"mymod/internal/database/dao"
	"mymod/internal/models"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао.

func SongInsert(instance tables.Song) models.Response {
	log.Debug("services layer get = ", instance)

	// отправка в сервис обогащения данных
	instance, err := EnrichtSong(instance)
	// можно убрать, если мы хотим чтобы insert работал без доступа к внешнему серверу
	if err != nil {
		return responses.ResponseSong{}.ExternalError()
	}
	log.Debug("after Enricht = ", instance)

	return instance.RecordCreate(database.GlobalPostgres, &dao.SongDao{})
}

func SongShow(instance tables.Song, settings models.TableSettings) models.Response {
	log.Debug("services layer get = ", instance)

	return instance.RecordShow(database.GlobalPostgres, &dao.SongDao{}, settings)
}

func SongUpdate(instance tables.Song) models.Response {
	log.Debug("services layer get = ", instance)

	return instance.RecordUpdate(database.GlobalPostgres, &dao.SongDao{})
}

func SongDelete(instance tables.Song) models.Response {
	log.Debug("services layer get = ", instance)

	return instance.RecordDelete(database.GlobalPostgres, &dao.SongDao{})
}
