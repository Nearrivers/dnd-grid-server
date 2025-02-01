package main

import (
	"log"

	"github.com/Nearrivers/dnd-grid-server/api/routes"
	"github.com/Nearrivers/dnd-grid-server/pkg/levels"
	"github.com/Nearrivers/dnd-grid-server/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := Setup()
	log.Fatal(app.Listen(":3000"))
}

func Setup() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	db := models.ConnectToDb()

	app.Static("/", "./assets")
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(cors.New())

	levelService := levels.NewService(db)
	routes.BookRouter(app, levelService)

	return app
}
