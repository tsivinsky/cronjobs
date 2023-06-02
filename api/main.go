package main

import (
	"api/db"
	"api/env"
	"log"

	"github.com/gofiber/fiber/v2"
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

	log.Fatal(app.Listen(":5000"))
}
