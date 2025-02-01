package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Nearrivers/dnd-grid-server/api/presenter"
	"github.com/Nearrivers/dnd-grid-server/api/validators"
	"github.com/Nearrivers/dnd-grid-server/pkg/levels"
	"github.com/Nearrivers/dnd-grid-server/pkg/models/repository"
	"github.com/gofiber/fiber/v2"
)

func AddLevel(service levels.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var nl presenter.Level
		err := c.BodyParser(&nl)
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(presenter.LevelErrorResponse(err))
		}

		err = validators.ValidateLevel(nl)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.LevelErrorResponse(err))
		}

		err = service.NewLevel(repository.NewLevelParams{
			Name:      nl.Name,
			ImagePath: nl.ImagePath,
			GridWidth: nl.GridWidth,
			GridColor: nl.GridColor,
			GridSpacing: sql.NullInt64{
				Int64: nl.GridSpacing,
			},
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.LevelErrorResponse(err))
		}

		c.Status(http.StatusCreated)
		return c.JSON(presenter.EmptyLevelSucessResponse())
	}
}

func UploadLevelImage(service levels.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		img, err := c.FormFile("image")
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(presenter.ImageUploadErrorResponse(err))
		}

		err = c.SaveFile(img, fmt.Sprintf("./assets/%s", img.Filename))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ImageUploadErrorResponse(err))
		}

		c.Status(http.StatusCreated)
		return c.JSON(presenter.ImageUploadSuccessResponse())
	}
}

func GetLevels(service levels.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		levels, err := service.GetLevels()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.LevelErrorResponse(err))
		}

		if len(levels) == 0 {
			c.Status(http.StatusNotFound)
			return c.JSON(presenter.LevelErrorResponse(fiber.ErrNotFound))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.LevelsSuccessResponse(levels))
	}
}
