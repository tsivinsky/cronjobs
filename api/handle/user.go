package handle

import (
	"api/data"
	"api/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	u := c.Cookies("user_id")
	if u == "" {
		return c.Status(403).JSON(fiber.Map{})
	}

	userId, err := strconv.Atoi(u)
	if err != nil {
		return err
	}

	var user data.User
	if tx := db.Db.Where("id = ?", userId).First(&user); tx.Error != nil {
		return tx.Error
	}

	return c.JSON(user)
}
