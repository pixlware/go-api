package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: Config.Env != "local" && Config.Env != "default",
	})

	app.Use(cors.New(cors.Config{
		AllowMethods: Config.CorsMethods,
		AllowOrigins: Config.CorsOrigins,
	}))

	setupRoutes(app)

	log.Println("Running '" + Config.Env + "' environment on port: " + Config.Port)
	app.Listen(":" + Config.Port)
}

func setupRoutes(app *fiber.App) {
	app.Get("/", apiHandler)
}

func apiHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
