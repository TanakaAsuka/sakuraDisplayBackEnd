package router

import (
	"github.com/gofiber/fiber/v2"
)

// Registe all route
func Registe(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	// get gallery
	app.Get("/gallery", func(c *fiber.Ctx) error {
		return handleGallery(c)
	})
	// get pixiv
	app.Get("/pixiv", func(c *fiber.Ctx) error {
		return handlePixiv(c)
	})
	app.Get("/err", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

}
