package presenter

import (
	"github.com/Nearrivers/dnd-grid-server/pkg/models/repository"
	"github.com/gofiber/fiber/v2"
)

type Level struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required"`
	ImagePath   string `json:"image_path" validate:"required"`
	GridWidth   int64  `json:"grid_width" validate:"required,numeric"`
	GridColor   string `json:"grid_color" validate:"required,hexcolor"`
	GridSpacing int64  `json:"grid_spacing" validate:"required,numeric"`
}

func mapLevel(l repository.Levels) Level {
	return Level{
		ID:          l.ID,
		Name:        l.Name,
		ImagePath:   l.ImagePath,
		GridWidth:   l.GridWidth,
		GridColor:   l.GridColor,
		GridSpacing: l.GridSpacing.Int64,
	}
}

func mapLevels(l []repository.Levels) []Level {
	var levels []Level

	for _, level := range l {
		levels = append(levels, mapLevel(level))
	}

	return levels
}

func LevelSuccessResponse(data repository.Levels) *fiber.Map {
	level := mapLevel(data)
	return &fiber.Map{
		"status": true,
		"data":   level,
		"error":  nil,
	}
}

func LevelsSuccessResponse(data []repository.Levels) *fiber.Map {
	levels := mapLevels(data)
	return &fiber.Map{
		"status": true,
		"data":   &levels,
		"error":  nil,
	}
}

func EmptyLevelSucessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
		"error":  nil,
	}
}

func ImageUploadSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "image upload successful",
		"error":  nil,
	}
}

func ImageUploadErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "image upload failed",
		"error":  err.Error(),
	}
}

func LevelErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
