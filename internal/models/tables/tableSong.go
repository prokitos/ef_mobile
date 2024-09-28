package tables

import (
	"mymod/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// таблица Song. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type Song struct {
	SongId      int     `json:"id,omitempty" gorm:"unique;primaryKey;autoIncrement"`
	Group       string  `json:"group" example:"aria"`
	Song        string  `json:"song" example:"some song"`
	ReleaseDate string  `json:"release_date" example:"01.01.2000"`
	Link        string  `json:"link" example:"http://whatever"`
	Text        []Verse `json:"text" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:OwnerId;references:SongId"`
}

type Verse struct {
	OwnerId   int64  `json:"-" example:"" gorm:"primaryKey;autoIncrement"`
	VerseId   int    `json:"verse_id" example:"1" gorm:"primaryKey"`
	VerseText string `json:"verse" example:"first verse"`
}

// таблица для получения данных из внешнего сервера
type ExternalSong struct {
	ReleaseDate string `json:"release_date" example:"01.01.2010"`
	Link        string `json:"link" example:"http://some.com"`
	Text        string `json:"text" example:"get lorem , dogs"`
}

// сюда передаётся экземпляр базы данных и дао, после этого мы вызываем методы пришедшего DAO и передаём туда нашу таблицу.
// можно было из сервиса вызывать напрямую дао, но если следовать ООП, то мы вызываем работу с данными сущности song, и поэтому вызов в коде идёт сначала через неё.
func (instance *Song) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Song) RecordShow(db models.DatabaseCore, dao models.DatabaseDao, settings models.TableSettings) models.Response {
	err := dao.ShowData(instance, db, settings)
	return err
}
func (instance *Song) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Song) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Song) GetId() int {
	return instance.SongId
}

func (instance *Song) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.SongId = id
	return nil
}

func (instance *Song) GetQueryParams(c *fiber.Ctx) error {
	instance.Group = c.Query("group", "")
	instance.Song = c.Query("song", "")
	instance.ReleaseDate = c.Query("release_date", "")
	instance.Link = c.Query("link", "")

	log.Debug("get query param = ", instance)
	return nil
}

func (instance *Song) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}

	log.Debug("get body param = ", instance)
	return nil
}

// получение limit offst verse из query
func (instance *Song) GetSettings(c *fiber.Ctx) models.TableSettings {
	tempLimit := c.Query("limit", "")
	tempOffset := c.Query("offset", "")
	tempVerse := c.Query("verse", "")

	var setting SongSettings

	limit, err := strconv.Atoi(tempLimit)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(tempOffset)
	if err != nil {
		offset = 0
	}
	verse, err := strconv.Atoi(tempVerse)
	if err != nil {
		verse = 0
	}

	setting.Limit = limit
	setting.Offset = offset
	setting.VerseId = verse

	log.Debug("get settings param = ", instance)
	return setting
}

type SongSettings struct {
	VerseId int
	Offset  int
	Limit   int
}

func (instance SongSettings) GetLimit() int {
	return instance.Limit
}

func (instance SongSettings) GetOffset() int {
	return instance.Offset
}
func (instance SongSettings) GetSpecData() int {
	return instance.VerseId
}
