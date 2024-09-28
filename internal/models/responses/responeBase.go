package responses

import (
	"errors"
	"mymod/internal/models"

	"github.com/gofiber/fiber/v2"
)

// основные респонсы к таблицам. разные таблицы просто меняют описание ответов.

type ResponseBase struct {
	Description string         `json:"description"       		 example:"description"`
	Code        int            `json:"code"               		 example:"status"`
	Data        []models.Table `json:"data,omitempty"       	 example:"...."`
}

func (instance ResponseBase) GoodCreate(tableName string) models.Response {
	instance.Code = 200
	instance.Description = tableName + " create success"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) BadCreate(tableName string) models.Response {
	instance.Code = 400
	instance.Description = tableName + " create error"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) GoodUpdate(tableName string) models.Response {
	instance.Code = 200
	instance.Description = tableName + " update success"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) BadUpdate(tableName string) models.Response {
	instance.Code = 400
	instance.Description = tableName + " update error"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) GoodDelete(tableName string) models.Response {
	instance.Code = 200
	instance.Description = tableName + " delete success"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) BadDelete(tableName string) models.Response {
	instance.Code = 400
	instance.Description = tableName + " delete error"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) GoodShow(curData []models.Table, tableName string) models.Response {
	instance.Code = 200
	instance.Description = tableName + " show success"
	instance.Data = curData
	return instance.GetResponse()
}
func (instance ResponseBase) BadShow(tableName string) models.Response {
	instance.Code = 400
	instance.Description = tableName + " show error"
	instance.Data = nil
	return instance.GetResponse()
}
func (instance ResponseBase) InternalError() models.Response {
	instance.Code = 400
	instance.Description = "internal error"
	instance.Data = nil
	return instance.GetResponse()
}

func (instance ResponseBase) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
func (instance ResponseBase) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}

func (instance ResponseBase) GetResponse() models.Response {
	var temp models.Response
	temp = instance
	return temp
}

func (instance ResponseBase) BaseServerError() error {
	result := "Code:400; Description: Internal Error"
	return errors.New(result)
}
func (instance ResponseBase) BaseExternalError() error {
	result := "Code:400; Description: External Server Error"
	return errors.New(result)
}
