package api

import (
	"fmt"

	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	return c.JSON(map[string]string{"user": id})
}

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Kumar",
		ID:        "potta",
	}

	return c.JSON(u)
}
