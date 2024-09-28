package models

import (
	"mymod/internal/config"

	"github.com/gofiber/fiber/v2"
)

// все интерфейсы

type Response interface {
	GetError(c *fiber.Ctx) error
	Validate() bool
}

type Table interface {
	RecordCreate(DatabaseCore, DatabaseDao) Response
	RecordDelete(DatabaseCore, DatabaseDao) Response
	RecordShow(DatabaseCore, DatabaseDao, TableSettings) Response
	RecordUpdate(DatabaseCore, DatabaseDao) Response
}

type TableSettings interface {
	GetLimit() int
	GetOffset() int
	GetSpecData() int
}
type DatabaseCore interface {
	OpenConnection(config.MainConfig)
	StartMigration()
	GlobalSet()
}
type DatabaseDao interface {
	CreateData(Table, DatabaseCore) Response
	DeleteData(Table, DatabaseCore) Response
	UpdateData(Table, DatabaseCore) Response
	ShowData(Table, DatabaseCore, TableSettings) Response
}
