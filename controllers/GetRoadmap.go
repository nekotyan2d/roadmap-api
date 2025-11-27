package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nekotyan2d/roadmap-api/data"
)

func GetRoadmap(c fiber.Ctx) error {
	id := c.Params("id")

	roadmaps := data.RoadmapsFull

	var roadmap *map[string]data.RoadmapData
	for key, r := range roadmaps {
		if key == id {
			roadmap = &r
			break
		}
	}

	if roadmap == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Roadmap not found",
		})
	}

	return c.JSON(fiber.Map{
		"response": roadmap,
	})
}
