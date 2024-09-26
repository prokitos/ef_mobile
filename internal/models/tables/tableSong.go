package tables

import (
	"mymod/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// таблица User. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type Song struct {
	Song_id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Group       string `json:"group" example:"aria"`
	Song        string `json:"song" example:"some song"`
	ReleaseDate string `json:"release_date" example:"01.01.2000"`
	Text        string `json:"text" example:"some text"`
	Link        string `json:"link" example:"http://whatever"`
}

func (instance *Song) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Song) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Song) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Song) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Song) GetId() int {
	return instance.Song_id
}

func (instance *Song) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.Song_id = id
	return nil
}

func (instance *Song) GetQueryParams(c *fiber.Ctx) error {
	instance.Group = c.Query("group", "")
	instance.Song = c.Query("song", "")
	instance.ReleaseDate = c.Query("release_date", "")
	instance.Text = c.Query("text", "")
	instance.Link = c.Query("link", "")
	return nil
}

func (instance *Song) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
