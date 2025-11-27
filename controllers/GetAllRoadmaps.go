package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nekotyan2d/roadmap-api/data"
)

func GetAllRoadmaps(c fiber.Ctx) error {
	roadmaps := data.Roadmaps

	return c.JSON(fiber.Map{
		"response": roadmaps,
	})
}
