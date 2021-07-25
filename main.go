package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"sakuradisplay/app/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func createFolder(path string) {
	// 检查文件夹是否存在
	if _, err := os.Stat(path); err == nil {
		// 文件夹存在,略过
		fmt.Println("path exists", path)
	} else {
		// 文件夹不存在,则创建文件
		fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0711)

		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
		}
	}
}

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})
	createFolder("./public")

	router.Registe(app)

	log.Fatal(app.Listen(":80"))
}
