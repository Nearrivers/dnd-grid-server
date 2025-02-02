package routes

import (
	"github.com/Nearrivers/dnd-grid-server/api/handlers"
	"github.com/Nearrivers/dnd-grid-server/pkg/levels"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service levels.Service) {
	app.Post("/levels", handlers.AddLevel(service))
	app.Post("/levels/image", handlers.UploadLevelImage(service))
	app.Get("/levels", handlers.GetLevels(service))
	app.Delete("/levels/:id", handlers.DeleteLevels(service))
}
