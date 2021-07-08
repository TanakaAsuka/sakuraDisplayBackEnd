package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handleGallery(c *fiber.Ctx) error {
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
