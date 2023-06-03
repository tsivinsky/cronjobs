package main

import (
	"api/db"
	"api/env"
	"api/handle"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := env.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     env.Env.AllowOrigin,
		AllowCredentials: true,
	}))
	app.Use(recover.New())

	app.Get("/auth/github", handle.GitHubEntry)
	app.Get("/auth/github/callback", handle.GitHubLogin)
	app.Get("/auth/logout", handle.LogoutUser)

	app.Get("/user", handle.GetUser)

	log.Fatal(app.Listen(":5000"))
}
