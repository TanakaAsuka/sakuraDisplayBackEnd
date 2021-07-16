package router

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"sakuradisplay/database"

	"github.com/gofiber/fiber/v2"
)

func handleLogin(c *fiber.Ctx) error {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Add record into postgreSQL
	username := c.FormValue("username")
	password := c.FormValue("password")
	parternWrap := map[string]string{
		username: database.UserPattern,
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

	// 查询数据库用户是否存在
	queryStr := fmt.Sprintf("SELECT * FROM user_table WHERE username='%s'", username)

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
	// 用户名为空
	if user.UserName == "" {
		return c.JSON(&fiber.Map{
			"err": 1,
			"msg": "用户名或密码不正确",
		})
	}

	m5 := md5.New()
	// 密码
	m5.Write([]byte(password))
	// 加盐
	m5.Write([]byte(fmt.Sprintf("%s", user.Salt)))
	st := m5.Sum(nil)
	passResult := hex.EncodeToString(st)
	fmt.Println("passResult:", passResult)
	// 用户密码不对
	if passResult != user.Password {
		return c.JSON(&fiber.Map{
			"err": 1,
			"msg": "用户名或密码不正确",
		})
	}

	// 持久化
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if name := sess.Get("username"); name == nil {
		sess.Set("username", user.UserName)

		// Save session
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"err":  0,
			"msg":  "登录成功！",
			"name": username,
		})

	}

	return c.JSON(&fiber.Map{
		"err": 0,
		"msg": "已经登录过了！",
	})

}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
