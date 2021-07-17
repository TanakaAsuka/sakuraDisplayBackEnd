package router

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

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
		log.Printf("IP:%s访问接口register", c.IP())
		err := handleRegister(c)
		if err != nil {
			fmt.Println(err)
			return errors.New("注册出错")
		}
		return nil
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		log.Printf("IP:%s访问接口login", c.IP())

		return handleLogin(c)
	})

	app.Get("/gallery", func(c *fiber.Ctx) error {
		return handleGallery(c)
	})
	app.Get("/pixiv", func(c *fiber.Ctx) error {
		return handlePixiv(c)
	})
	app.Get("/userauth", func(c *fiber.Ctx) error {
		return handleUserAuth(c)
	})
	app.Post("/upload", func(c *fiber.Ctx) error {
		log.Printf("IP:%s访问接口upload", c.IP())
		return handleUpload(c)
	})
	app.Post("/delete", func(c *fiber.Ctx) error {
		log.Printf("IP:%s访问接口delete", c.IP())
		return handleDelete(c)
	})
	app.Get("/test", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		name := sess.Get("username")

		return c.SendString(fmt.Sprintf("%v", name))

	})
	app.Static("/", "./public")
	// 服务静态文件
	app.Get("/*", func(c *fiber.Ctx) error {
		err := c.SendFile("./assets" + c.Path())
		if err != nil {
			return c.SendStatus(404)
		}
		return nil
	})

}
