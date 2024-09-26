package transport

import (
	"github.com/gofiber/fiber/v2"
)

// здесь хранятся хэндлеры.

func SetHandlers(instance *fiber.App) {

	instance.Get("/song", getSong)
	instance.Post("/song", insertSong)
	instance.Delete("/song", deleteSong)
	instance.Put("/song", updateSong)
}
