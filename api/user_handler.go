package api

import (
	"errors"

	"github.com/bupd/hotelmanager/db"
	"github.com/bupd/hotelmanager/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) HandleDelUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": userID})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	// TO-DO createusers with store
	// basically how all these works is create a route in the main function
	// attach a handler to the route which takes the params and set Context
	// to call the correct method for the route
	// handler calls the method to update the db
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	errors := params.Validate()
	if len(errors) > 0 {
		return c.JSON(errors)
	}

	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"msg": "user not found."})
		}
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}
