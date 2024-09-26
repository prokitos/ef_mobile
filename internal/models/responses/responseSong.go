package responses

import (
	"mymod/internal/models"
	"mymod/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

// ответы для таблицы Song

type ResponseSong struct {
	Description string        `json:"description"        example:"description"`
	Code        int           `json:"code"               example:"status"`
	Users       []tables.Song `json:"songs,omitempty"    example:"...."`
}

func (instance ResponseSong) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("song")
}
func (instance ResponseSong) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("song")
}
func (instance ResponseSong) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("song")
}
func (instance ResponseSong) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("song")
}
func (instance ResponseSong) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("song")
}
func (instance ResponseSong) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("song")
}
func (instance ResponseSong) GoodShow(curSong []tables.Song) models.Response {
	var items []models.Table
	for i := 0; i < len(curSong); i++ {
		items = append(items, &curSong[i])
	}
	return ResponseBase{}.GoodShow(items, "song")
}
func (instance ResponseSong) BadShow() models.Response {
	return ResponseBase{}.BadShow("song")
}
func (instance ResponseSong) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseSong) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}

func (instance ResponseSong) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
