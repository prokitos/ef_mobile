package transport

import (
	"mymod/internal/models/tables"
	"mymod/internal/services"

	"github.com/gofiber/fiber/v2"
)

// Users godoc
// @Summary get Song
// @Description get Song by params and offset and limit
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id query string false "Show by id"
// @Param group query string false "Show by group"
// @Param song query string false "Show by song"
// @Param release_date query string false "Show by release_date"
// @Param link query string false "Show by link"
// @Param offset query string false "Show by offset"
// @Param limit query string false "Show by limit"
// @Param verse query string false "Show by verse"
// @Success 200 "successful operation"
// @Router /song [get]
func getSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)
	curSong.GetQueryParams(c)
	settings := curSong.GetSettings(c)

	return services.SongShow(curSong, settings).GetError(c)
}

// Users godoc
// @Summary insert Song
// @Description insert Song by body params and erichment
// @Tags Song
// @Accept  json
// @Produce  json
// @Param orderBook body tables.Song true "insert song"
// @Success 200 "successful operation"
// @Router /song [post]
func insertSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetBodyParams(c)

	return services.SongInsert(curSong).GetError(c)
}

// Users godoc
// @Summary delete Song
// @Description delete Song by id
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id query string false "deleted by id"
// @Success 200 "successful operation"
// @Router /song [delete]
func deleteSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetQueryId(c)

	return services.SongDelete(curSong).GetError(c)
}

// Users godoc
// @Summary update Song
// @Description update Song by body params and id
// @Tags Song
// @Accept  json
// @Produce  json
// @Param id query string false "update by id"
// @Param orderBook body tables.Song true "update order"
// @Success 200 "successful operation"
// @Router /song [put]
func updateSong(c *fiber.Ctx) error {
	var curSong tables.Song
	curSong.GetBodyParams(c)
	curSong.GetQueryId(c)

	return services.SongUpdate(curSong).GetError(c)
}
