package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

// ответы для таблицы User

type ResponseUser struct {
	Description string        `json:"description"        example:"description"`
	Code        int           `json:"code"               example:"status"`
	Users       []tables.User `json:"users,omitempty"    example:"...."`
}

func (instance ResponseUser) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("user")
}
func (instance ResponseUser) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("user")
}
func (instance ResponseUser) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("user")
}
func (instance ResponseUser) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("user")
}
func (instance ResponseUser) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("user")
}
func (instance ResponseUser) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("user")
}
func (instance ResponseUser) GoodShow(curUser []tables.User) models.Response {
	var items []models.Table
	for i := 0; i < len(curUser); i++ {
		items = append(items, &curUser[i])
	}
	return ResponseBase{}.GoodShow(items, "user")
}
func (instance ResponseUser) BadShow() models.Response {
	return ResponseBase{}.BadShow("user")
}
func (instance ResponseUser) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseUser) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}

func (instance ResponseUser) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
