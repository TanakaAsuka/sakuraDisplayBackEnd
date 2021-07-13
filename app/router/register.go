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

	// Add record into postgreSQL
	username := "yixuan4"
	nickname := "逸轩"
	password := "123456"

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

	return c.SendString("注册成功")
}
