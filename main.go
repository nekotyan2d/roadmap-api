package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/nekotyan2d/roadmap-api/controllers"
	"github.com/nekotyan2d/roadmap-api/data"
)

func main() {
	data.LoadRoadmaps()

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/roadmaps", controllers.GetAllRoadmaps)
	app.Get("/roadmaps/:id", controllers.GetRoadmap)

	log.Fatal(app.Listen(":3000"))
}
