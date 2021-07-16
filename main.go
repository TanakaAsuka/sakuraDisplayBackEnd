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
	app := fiber.New()
	// store := session.New()
	// app.Get("/test", func(c *fiber.Ctx) error {

	// 	sess, err := store.Get(c)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	name := sess.Get("name")

	// 	if name == nil {
	// 		sess.Set("name", "john")
	// 		sess.Set("info", "已登录")
	// 		// Save session
	// 		if err := sess.Save(); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// 	return c.SendString(fmt.Sprintf("Welcome %v,%v", name, sess.Get("info")))
	// })

	router.Registe(app)

	log.Fatal(app.Listen(":3000"))
}
