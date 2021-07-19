package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"sakuradisplay/app/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// Store a session

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	router.Registe(app)

	log.Fatal(app.Listen(":3000"))
}
