package router

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Registe all route
func Registe(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {
		// 允许跨域
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET")
		c.Set("Access-Control-Allow-Methods", "POST")

		return c.Next()
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		err := handleRegister(c)
		if err != nil {
			fmt.Println(err)
			return errors.New("注册出错")
		}
		return nil
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return handleLogin(c)
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
	app.Static("/", "./public")
	// 服务静态文件
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./assets" + c.Path())
	})

}
