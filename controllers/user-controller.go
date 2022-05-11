package controllers

import (
	"user_api/helper"
	"user_api/models"
	"user_api/services"

	"github.com/gofiber/fiber/v2"
)

type UserC interface {
	GetByToken(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type userC struct {
	userS services.UserS
}

func NewUserC(userS services.UserS) UserC {
	return &userC{userS: userS}
}

func (u *userC) GetByToken(c *fiber.Ctx) error {
	user, err := u.userS.GetOne(c.Context(), c.Get("Authorization"))
	if err != nil {
		return helper.Response(c, fiber.StatusBadRequest, nil, err.Error(), false)
	}

	return helper.Response(c, fiber.StatusOK, user, "Get user success!", true)
}

func (u *userC) UpdateUser(c *fiber.Ctx) error {
	var update models.Update

	err := c.BodyParser(&update)
	if err != nil {
		return helper.Response(c, fiber.StatusNotAcceptable, nil, err.Error(), false)
	}

	if err = u.userS.CheckEmail(c.Context(), update.Email, c.Get("Authorization")); err != nil {
		return helper.Response(c, fiber.StatusNotAcceptable, nil, err.Error(), false)
	}

	token, err := u.userS.Update(c.Context(), update, c.Get("Authorization"))
	if err != nil {
		return helper.Response(c, fiber.StatusConflict, nil, err.Error(), false)
	}

	user, err := u.userS.GetOne(c.Context(), c.Get("Authorization"))
	if err != nil {
		return helper.Response(c, fiber.StatusBadRequest, nil, err.Error(), false)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"token":   token,
		"status":  true,
		"message": "Update user success",
	})
}
