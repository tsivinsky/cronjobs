package handle

import (
	"api/data"
	"api/db"
	"api/env"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GitHubAccessTokenBody struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type GitHubAccessTokenResult struct {
	AccessToken string `json:"access_token"`
}

type GitHubUser struct {
	Email  *string `json:"email"`
	Avatar *string `json:"avatar_url"`
	Login  string  `json:"login"`
}

func GitHubEntry(c *fiber.Ctx) error {
	uri := fmt.Sprintf("https://github.com/login/oauth/authorize?scope=read:user&client_id=%s", env.Env.GHClientId)

	return c.Redirect(uri)
}

func GitHubLogin(c *fiber.Ctx) error {
	code := c.Query("code")

	ghTokenBody := GitHubAccessTokenBody{
		ClientId:     env.Env.GHClientId,
		ClientSecret: env.Env.GHClientSecret,
		Code:         code,
	}
	b, err := json.Marshal(&ghTokenBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res GitHubAccessTokenResult
	err = json.Unmarshal(r, &res)
	if err != nil {
		return err
	}

	if res.AccessToken == "" {
		return errors.New("couldn't get access token from github")
	}

	uReq, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return err
	}
	uReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", res.AccessToken))
	uReq.Header.Set("Accept", "application/json")

	uResp, err := http.DefaultClient.Do(uReq)
	if err != nil {
		return err
	}
	defer uResp.Body.Close()

	ur, err := io.ReadAll(uResp.Body)
	if err != nil {
		return err
	}

	var ghUser GitHubUser
	err = json.Unmarshal(ur, &ghUser)
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

	return c.Redirect(env.Env.WEB_APP_URL)
}
