package handle

import (
	"api/env"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "user_id",
		Value:    "",
		Path:     "/",
		Domain:   env.Env.COOKIE_DOMAIN,
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * -9999),
	})

	return c.JSON(fiber.Map{})
}
