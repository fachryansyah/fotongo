package main

import (
	"fotongo/model"
	"fotongo/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	model.InitDB()

	app.Use(cors.New())

	router.AuthRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test api")
	})

	app.Listen(":1337")
}
