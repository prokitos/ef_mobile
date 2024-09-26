package dao

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// (дао/круд) для таблицы Task. вызывается из соответсвующей таблицы.

type TaskDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов у новых дао.
func (currentlDB *TaskDao) curResponse() responses.ResponseTask {
	return responses.ResponseTask{}
}

func (currentlDB *TaskDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Create(&task); result.Error != nil {
		log.Debug("create record error!")
		return currentlDB.curResponse().BadCreate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *TaskDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := task.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	result := dbConnect.Instance.Delete(&task, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return currentlDB.curResponse().BadDelete()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *TaskDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Updates(&task); result.Error != nil {
		log.Debug("update record error!")
		return currentlDB.curResponse().BadUpdate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *TaskDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao get = ", data)

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.Task

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	results := dbConnect.Instance.Find(&finded, task)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return currentlDB.curResponse().BadShow()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodShow(finded)
}

// перево интерфейса таблицы в конкретную таблицу
func (currentlDB *TaskDao) getData(temp models.Table) (tables.Task, models.Response) {
	task, ok := temp.(*tables.Task)
	if ok == false {
		return tables.Task{}, currentlDB.curResponse().InternalError()
	}
	return *task, nil
}
