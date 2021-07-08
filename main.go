package main

import (
	"sakuradisplay/app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Registe(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
