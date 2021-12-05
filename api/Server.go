package api

import (
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	_ "fotongo/docs"
	// _ "github.com/arsmn/fiber-swagger/v2/example/docs"
)

// @BasePath /api
func InitializeServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fotongo 1.1",
	})

	// app.Use(middleware.Log())

	if os.Getenv("APP_ENV") == "local" {
		app.Get("/swagger/*", swagger.Handler) // default

		app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
			DeepLinking: false,
			// Expand ("list") or Collapse ("none") tag groups by default
			DocExpansion: "none",
			// Prefill OAuth ClientId on Authorize popup
			OAuth: &swagger.OAuthConfig{
				AppName:  "OAuth Provider",
				ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
			},
			// Ability to change OAuth2 redirect uri location
			OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
		}))
	}

	// api := app.Group("/api")

	// v1 := api.Group("/v1")

	// authController, err := InitializeAuthController()
	// if err != nil {
	// 	panic(err)
	// }
	// authController.Route(v1)

	return app
}
