package handle

import (
	"api/data"
	"api/db"
	"api/env"
	"api/github"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const SessionLifeTime = time.Second * 60 * 60 * 24 * 30

func GitHubEntry(c *fiber.Ctx) error {
	uri := fmt.Sprintf("https://github.com/login/oauth/authorize?scope=read:user&client_id=%s", env.Env.GHClientId)

	return c.Redirect(uri)
}

func GitHubLogin(c *fiber.Ctx) error {
	code := c.Query("code")

	accessToken, err := github.GetAccessToken(code)
	if err != nil {
		return err
	}

	ghUser, err := github.GetUserData(accessToken)
	if err != nil {
		return err
	}

	var user *data.User
	if tx := db.Db.Where("github_login = ?", ghUser.Login).First(&user); tx.Error != nil {
		user.GitHubLogin = ghUser.Login
		user.Avatar = ghUser.Avatar
		user.Email = ghUser.Email
		db.Db.Create(&user)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(int(user.ID)),
		Path:     "/",
		Domain:   env.Env.COOKIE_DOMAIN,
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(SessionLifeTime),
	})

	return c.Redirect(env.Env.WEB_APP_URL)
}
