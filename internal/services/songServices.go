package services

import (
	"modules/internal/database"
	"modules/internal/database/dao"
	"modules/internal/models"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао. Вызывается из роутов.

func UserInsert(instance tables.User) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordCreate(database.GlobalPostgres, dao.GlobalUserDao)
}

func UserShow(instance tables.User) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordShow(database.GlobalPostgres, dao.GlobalUserDao)
}

func UserUpdate(instance tables.User) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordUpdate(database.GlobalPostgres, dao.GlobalUserDao)
}

func UserDelete(instance tables.User) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordDelete(database.GlobalPostgres, dao.GlobalUserDao)
}
