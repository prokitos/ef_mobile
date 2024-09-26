package tables

import (
	"modules/internal/models"
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

func (instance *User) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *User) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *User) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *User) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *User) GetId() int {
	return instance.User_id
}

func (instance *User) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.User_id = id
	return nil
}

func (instance *User) GetQueryParams(c *fiber.Ctx) error {
	instance.Login = c.Query("login", "")
	instance.Password = c.Query("password", "")
	return nil
}

func (instance *User) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
