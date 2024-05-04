package api

import (
	"context"

	"github.com/bupd/hotelmanager/db"
	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Kumar",
		ID:        "potta",
	}

	return c.JSON(u)
}
