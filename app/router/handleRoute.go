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
func getData(c *fiber.Ctx) (database.Images, error) {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	// get record from database
	rows, err := database.DB.Query("SELECT * FROM images_table ORDER BY random() LIMIT 10")
	defer rows.Close()

	if err != nil {
		return database.Images{}, err
	}

	images := database.Images{}

	for rows.Next() {
		img := database.Image{}
		if err := rows.Scan(&img.UUID, &img.URL, &img.WidthAndHeight); err != nil {
			return database.Images{}, err // Exit if we get an error
		}

		// Append Employee to Employees
		images.ImagesList = append(images.ImagesList, img)
	}
	// Return Employees in JSON format
	return images, nil
}

func handleGallery(c *fiber.Ctx) error {

	imgs, err := getData(c)

	if err != nil {
		return nil
	}
	return c.JSON(imgs)
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
func handleUserAuth(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if name := sess.Get("username"); name == nil {

		return c.Status(200).JSON(&fiber.Map{
			"err": 1,
			"msg": "请先登录",
		})

	}
	return c.Status(200).JSON(&fiber.Map{
		"err": 0,
		"msg": "用户已登录",
	})
}
func handleUpload(c *fiber.Ctx) error {
	// 如果用户没登录的话
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if name := sess.Get("username"); name == nil {

		return c.Status(200).JSON(&fiber.Map{
			"err": 1,
			"msg": "请先登录",
		})

	}
	// 处理登录用户上传的图片

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
		fmt.Println("format:", format)
		if err != nil {
			c.SendString("请上传图片格式!")
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
