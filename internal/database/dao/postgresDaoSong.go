package dao

import (
	"mymod/internal/models"
	"mymod/internal/models/responses"
	"mymod/internal/models/tables"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// (дао/круд) для таблицы Song. вызывается из соответсвующей таблицы.

type SongDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов если будут новые таблицы.
func (currentlDB *SongDao) curResponse() responses.ResponseSong {
	return responses.ResponseSong{}
}

// создание новой записи.
func (currentlDB *SongDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		log.Error("unexpected error when convert interface to data")
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		log.Error("dont connect to database")
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Create(&song); result.Error != nil {
		log.Error("create error")
		log.Debug("create error with data = ", song)
		return currentlDB.curResponse().BadCreate()
	}

	log.Info("Create complete")
	return currentlDB.curResponse().GoodCreate()
}

// удаление записи и связанной записи
func (currentlDB *SongDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		log.Error("unexpected error when convert interface to data")
		return resp
	}
	id := song.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		log.Error("dont connect to database")
		return currentlDB.curResponse().InternalError()
	}

	result := dbConnect.Instance.Select(clause.Associations).Delete(&song)
	if result.RowsAffected == 0 || result.Error != nil {
		log.Error("delete error")
		log.Debug("delete error with data = ", song)
		return currentlDB.curResponse().BadDelete()
	}
	dbConnect.Instance.Delete(&tables.Song{}, id)

	log.Info("Delete complete")
	return currentlDB.curResponse().GoodDelete()
}

// обновление записи и связанных записей
func (currentlDB *SongDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		log.Error("unexpected error when convert interface to data")
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		log.Error("dont connect to database")
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Updates(&song); result.Error != nil {
		log.Error("update error")
		log.Debug("update error with data = ", song)
		return currentlDB.curResponse().BadUpdate()
	}
	for i := 0; i < len(song.Text); i++ {
		dbConnect.Instance.Model(tables.Verse{}).Where("verse_id = ? and owner_id = ?", song.Text[i].VerseId, song.Text[i].OwnerId).Updates(&song.Text[i])
	}

	log.Info("Update complete")
	return currentlDB.curResponse().GoodUpdate()
}

// показать записи и связанные записи
func (currentlDB *SongDao) ShowData(data models.Table, core models.DatabaseCore, setting models.TableSettings) models.Response {
	log.Debug("dao get = ", data)

	song, resp := currentlDB.getData(data)
	if resp != nil {
		log.Error("unexpected error when convert interface to data")
		return resp
	}
	var finded []tables.Song

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		log.Error("dont connect to database")
		return currentlDB.curResponse().InternalError()
	}

	var results *gorm.DB
	if setting.GetSpecData() == 0 {
		results = dbConnect.Instance.Limit(setting.GetLimit()).Offset(setting.GetOffset()).Preload("Text").Find(&finded, song)
	} else {
		results = dbConnect.Instance.Limit(setting.GetLimit()).Offset(setting.GetOffset()).Preload("Text", "verse_id = ? ", setting.GetSpecData()).Find(&finded, song)
	}

	if results.Error != nil || results.RowsAffected == 0 {
		log.Error("nothing to show")
		log.Debug("show error with data = ", song)
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
