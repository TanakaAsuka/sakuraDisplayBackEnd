package router

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"sakuradisplay/database"

	"github.com/gofiber/fiber/v2"
)

func handleLogin(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Add record into postgreSQL
	username := "yixuan4"
	// nickname := "逸轩"
	password := "123456"

	queryStr := fmt.Sprintf("SELECT * FROM user_table WHERE username='%s'", username)
	fmt.Println("queryStr:", queryStr)

	rows, err := database.DB.Query(queryStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	user := database.User{}
	for rows.Next() {
		if err := rows.Scan(&user.UserName, &user.NickName, &user.Salt, &user.Password); err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("userResult:", user)

	m5 := md5.New()
	// 密码
	m5.Write([]byte(password))
	// 加盐
	m5.Write([]byte(fmt.Sprintf("%s", user.Salt)))
	st := m5.Sum(nil)
	passResult := hex.EncodeToString(st)
	fmt.Println("passResult:", passResult)

	if passResult != user.Password {
		return errors.New("用户名或密码不正确")
	}

	return c.SendString("登录成功")

}
