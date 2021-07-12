package main

import (
	_ "image/png"
	"log"
	"sakuradisplay/app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Registe(app)

	log.Fatal(app.Listen(":3000"))
}
