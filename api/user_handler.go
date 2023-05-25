package api

import (
	"github.com/AshirwadPradhan/hotelresapi/db"
	"github.com/AshirwadPradhan/hotelresapi/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var id = c.Params("id")

	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
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

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var reqparams types.CreateUserParams
	if err := c.BodyParser(&reqparams); err != nil {
		return err
	}

	if errors := reqparams.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}

	user, err := types.NewUserFromParams(reqparams)
	if err != nil {
		return err
	}

	createdUser, err := h.userStore.PostUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(createdUser)
}
