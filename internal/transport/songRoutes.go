package transport

import (
	"module/internal/models/tables"
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryParams(c)
	curSong.GetQueryId(c)
	return services.UserShow(curSong).GetError(c)
}

func insertSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryParams(c)
	return services.UserInsert(curSong).GetError(c)
}

func deleteSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)
	return services.UserDelete(curSong).GetError(c)
}

func updateSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryParams(c)
	curSong.GetQueryId(c)
	return services.UserUpdate(curSong).GetError(c)
}
