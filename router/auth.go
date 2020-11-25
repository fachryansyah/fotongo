package router

import (
	"fotongo/handler"
	"fotongo/middleware"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is for routing auth handler
func AuthRouter(app *fiber.App) error {

	api := app.Group("/auth")

	api.Post("/login", handler.LoginHandler)
	api.Post("/register", handler.RegisterHandler)
	api.Get("/token", middleware.ProtectedUser, handler.CheckTokenHandler)
	api.Get("/refresh-token", handler.RequestTokenHandler)

	return nil
}
