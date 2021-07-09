package router

import (
	"fmt"
	"log"
	"sakuradisplay/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Connect with database
func insertData() error {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// Add record into postgreSQL
	// New Image struct
	newUUID := uuid.New()
	img := database.Image{
		Author:         "焦茶",
		Ban:            false,
		UUID:           newUUID,
		Description:    "光影",
		WidthAndHeight: "1920-1080",
		Subject:        "人物",
		Tag:            []string{"女孩", "笑脸"},
		URL:            "http://sakuradisplay/img/" + newUUID.String() + ".jpg",
		Title:          "微笑的女孩",
	}

	// Insert Image into database
	res, err := database.DB.Query("INSERT INTO images_table (author, ban, uuid, description, subject, title, url, width_and_height) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", img.Author, img.Ban, img.UUID, img.Description, img.Subject, img.Title, img.URL, img.WidthAndHeight)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil

}

func handleGallery(c *fiber.Ctx) error {
	err := insertData()
	if err != nil {
		fmt.Println("插入数据出错", err)
	}

	return c.JSON(&fiber.Map{
		"imgObj": map[string]interface{}{
			"title":   "明日方舟德克萨斯",
			"desc":    "穿黑衣的德狗",
			"subject": "德克萨斯",
			"author":  "unknown",
			"ban":     false,
			"width":   1920,
			"height":  1080,
			"tag": []string{
				"明日方舟", "同人", "单马尾",
			},
			"url": "http://sakuradisplay/img/v3/fie5j5j30s0hgfkc0.jpg",
		},
		"imgObj2": map[string]interface{}{
			"title":   "明日方舟能天使",
			"desc":    "吃苹果派的能天使",
			"subject": "能天使",
			"author":  "unknown",
			"ban":     false,
			"width":   1920,
			"height":  1080,
			"tag": []string{
				"明日方舟", "同人", "红发",
			},
			"url": "http://sakuradisplay/img/v3/gfdhdhfg7543423s7u6u762v43.jpg",
		},
	})
}
func handlePixiv(c *fiber.Ctx) error {
	findType := c.Query("type")
	pixivID := c.Query("id")
	pixivURL := "https://www.pixiv.net/"

	switch findType {
	case "artworks":
		pixivURL = "https://www.pixiv.net/artworks/" + pixivID
	case "users":
		pixivURL = "https://www.pixiv.net/users/" + pixivID
	default:
		c.SendStatus(600)
		return c.SendString("你误入了结界，但这里什么也没有!")
	}

	fmt.Println(pixivURL)

	return c.Redirect(pixivURL)
}
func handleUpload(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		c.SendStatus(600)
		return err
	}
	files := form.File["uploadFile"]

	for _, file := range files {
		u4 := uuid.New()
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

		err := c.SaveFile(file, fmt.Sprintf("../../assets/%s", u4))

		if err != nil {
			return err
		}
	}
	fmt.Println("文件写入成功!")
	return c.SendStatus(200)
}
