package router

import (
	"fmt"
	"image"
	"log"
	"os"
	"sakuradisplay/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Connect with database
func insertData(u4 uuid.UUID, url string, WidthAndHeight string) error {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// Add record into postgreSQL

	// Insert Image into database
	res, err := database.DB.Query("INSERT INTO images_table (uuid, url, width_and_height) VALUES ($1, $2, $3)", u4, url, WidthAndHeight)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil

}

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
func handleUpload(c *fiber.Ctx) error {

	baseHost := c.BaseURL()

	form, err := c.MultipartForm()
	if err != nil {
		c.SendStatus(600)
		log.Println(err)
		return err
	}
	files := form.File["uploadFile"]

	for _, file := range files {
		// 获取图片宽高
		myfile, err := file.Open()
		config, format, err := image.DecodeConfig(myfile)
		if err != nil {
			log.Println(err)
			return err
		}
		widthAndHeight := fmt.Sprintf("%d-%d", config.Width, config.Height)

		u4 := uuid.New()
		fmt.Println(file.Filename, file.Size)
		// 获取当前年月日
		timeStr := time.Now().Format("20060102")

		// 检查文件夹是否存在
		path := "./assets/" + timeStr
		if _, err := os.Stat(path); err == nil {
			// 文件夹存在,略过
			fmt.Println("path exists 1", path)
		} else {
			// 文件夹不存在,则创建文件
			fmt.Println("path not exists ", path)
			err := os.MkdirAll(path, 0711)

			if err != nil {
				log.Println("Error creating directory")
				log.Println(err)
				return err
			}
		}
		// ./assets/20270707/xxxx.jpg
		basePath := fmt.Sprintf("%s/%s.%s", path, u4, format)
		// http://www.sakuradisplay/20210707/xxxxx.jpg
		// 去掉头部的字符 "./assets",先暂时这么写
		url := baseHost + basePath[8:]

		err = insertData(u4, url, widthAndHeight)

		if err != nil {
			log.Println(err)
			fmt.Println("插入数据出错", err)
			return err
		}
		err = c.SaveFile(file, basePath)
		if err != nil {
			return err
		}
	}

	fmt.Println("文件写入成功!")
	return c.SendStatus(200)
}
