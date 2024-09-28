package transport

import (
	"github.com/gofiber/fiber/v2"

	_ "mymod/docs"

	swagger "github.com/gofiber/swagger"
)

// здесь хранятся хэндлеры.

func SetHandlers(instance *fiber.App) {

	instance.Get("/song", getSong)
	instance.Post("/song", insertSong)
	instance.Delete("/song", deleteSong)
	instance.Put("/song", updateSong)

	instance.Get("/swagger/*", swagger.HandlerDefault)
}
