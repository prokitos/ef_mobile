package transport

import (
	"mymod/internal/models/tables"
	"mymod/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)
	curSong.GetQueryParams(c)
	limit, offset := curSong.GetLimitOffset(c)

	return services.SongShow(curSong, limit, offset).GetError(c)
}

func insertSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetBodyParams(c)

	return services.SongInsert(curSong).GetError(c)
}

func deleteSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)

	return services.SongDelete(curSong).GetError(c)
}

func updateSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)
	curSong.GetBodyParams(c)

	return services.SongUpdate(curSong).GetError(c)
}
