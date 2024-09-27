package dao

import (
	"mymod/internal/models"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// (дао/круд) для таблицы Song. вызывается из соответсвующей таблицы.

type SongDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов если будут новые таблицы.
func (currentlDB *SongDao) curResponse() responses.ResponseSong {
	return responses.ResponseSong{}
}

func (currentlDB *SongDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Create(&song); result.Error != nil {
		log.Debug("create record error!")
		return currentlDB.curResponse().BadCreate()
	}

	log.Info("Create complete")
	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *SongDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := song.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	result := dbConnect.Instance.Delete(&song, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return currentlDB.curResponse().BadDelete()
	}

	log.Info("Delete complete")
	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *SongDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Updates(&song); result.Error != nil {
		log.Debug("update record error!")
		return currentlDB.curResponse().BadUpdate()
	}

	log.Info("Update complete")
	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *SongDao) ShowData(data models.Table, core models.DatabaseCore, limit int, offset int) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.Song

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	results := dbConnect.Instance.Limit(limit).Offset(offset).Find(&finded, song)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return currentlDB.curResponse().BadShow()
	}

	log.Info("Show complete")
	return currentlDB.curResponse().GoodShow(finded)
}

// перевод интерфейса таблицы в конкретную таблицу
func (currentlDB *SongDao) getData(temp models.Table) (tables.Song, models.Response) {
	task, ok := temp.(*tables.Song)
	if ok == false {
		return tables.Song{}, currentlDB.curResponse().InternalError()
	}
	return *task, nil
}
