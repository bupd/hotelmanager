package api

import (
	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "kumar"})
}

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Kumar",
	}

	return c.JSON(u)
}
