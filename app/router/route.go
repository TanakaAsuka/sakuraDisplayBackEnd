package router

import (
	"github.com/gofiber/fiber/v2"
)

// Registe all route
func Registe(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	app.Get("/gallery", func(c *fiber.Ctx) error {
		return handleGallery(c)
	})
	app.Get("/pixiv", func(c *fiber.Ctx) error {
		return handlePixiv(c)
	})
	app.Post("/upload", func(c *fiber.Ctx) error {
		return handleUpload(c)
	})
	app.Get("/err", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

}
