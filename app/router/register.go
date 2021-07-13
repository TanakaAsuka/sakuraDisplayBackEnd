package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sakuradisplay/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func handleRegister(c *fiber.Ctx) error {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	// Add record into postgreSQL
	username := c.FormValue("username")
	nickname := c.FormValue("nickname")
	password := c.FormValue("password")

	fmt.Printf("username:%s,password:%s", username, password)

	// 验证用户名密码

	salt := time.Now().Unix()
	m5 := md5.New()
	// 密码
	m5.Write([]byte(password))
	// 加盐
	m5.Write([]byte(fmt.Sprintf("%d", salt)))
	st := m5.Sum(nil)
	passResult := hex.EncodeToString(st)

	// fmt.Println(st, hex.EncodeToString(st))

	// Insert user into database
	res, err := database.DB.Query("INSERT INTO user_table (username, nickname, salt,password) VALUES ($1, $2, $3,$4)", username, nickname, salt, passResult)

	if err != nil {
		return err
	}

	fmt.Println(res)

	c.SendStatus(200)

	return c.SendString("注册成功")
}
