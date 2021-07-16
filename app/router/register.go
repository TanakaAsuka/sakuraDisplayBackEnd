package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
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

	parternWrap := map[string]string{
		username: database.UserPattern,
		nickname: database.NickPattern,
		password: database.PassPattern,
	}

	fmt.Printf("username:%s,password:%s\n", username, password)

	// 验证用户名密码

	for k, v := range parternWrap {
		match, err := regexp.MatchString(v, k)
		if err != nil {
			return err
		}
		if !match {
			return c.JSON(&fiber.Map{
				"err": 1,
				"msg": "用户名或密码非法",
			})
		}
	}
	// 查询用户是否存在
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
		if err := rows.Scan(&user.UserName, &user.NickName, &user.Salt, &user.Password, &user.Role); err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("userResult:", user)
	if user.UserName == username {
		return c.JSON(&fiber.Map{
			"err": 1,
			"msg": "用户名已注册",
		})
	}

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
	// 管理员注册
	for _, admin := range database.Admin {

		fmt.Println("username:", username)
		if username == admin {
			res, err := database.DB.Query("INSERT INTO user_table (username, nickname, salt,password,role) VALUES ($1, $2, $3,$4,$5)", username, nickname, salt, passResult, "admin")
			if err != nil {
				return err
			}
			fmt.Println(res)
			return c.Status(200).JSON(&fiber.Map{
				"err":  0,
				"msg":  "注册成功",
				"role": "admin",
			})
		}
	}
	// 普通用户注册
	res, err := database.DB.Query("INSERT INTO user_table (username, nickname, salt,password,role) VALUES ($1, $2, $3,$4,$5)", username, nickname, salt, passResult, "visitor")
	if err != nil {
		return err
	}

	fmt.Println(res)

	return c.Status(200).JSON(&fiber.Map{
		"err":  0,
		"msg":  "注册成功",
		"role": "visitor",
	})
}
