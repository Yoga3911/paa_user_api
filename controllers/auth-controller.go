package controllers

import (
	"user_api/models"
	"user_api/services"
	"user_api/helper"
	"user_api/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthC interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authC struct {
	authS services.AuthS
}

func NewAuthC(authS services.AuthS) AuthC {
	return &authC{authS: authS}
}

func (a *authC) Login(c *fiber.Ctx) error {
	var user models.Login

	err := c.BodyParser(&user)
	if err != nil {
		return helper.Response(c, fiber.StatusNotAcceptable, nil, err.Error(), false)
	}

	if errors := middleware.StructValidator(user); errors != nil {
		return helper.Response(c, fiber.StatusConflict, nil, errors, false)
	}

	token, usr, err := a.authS.VerifyCredential(c.Context(), user)
	if err != nil {
		return helper.Response(c, fiber.StatusConflict, nil, err.Error(), false)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    usr,
		"token":   token,
		"status":  true,
		"message": "Login success!",
	})
}

func (a *authC) Register(c *fiber.Ctx) error {
	var user models.Register

	err := c.BodyParser(&user)
	if err != nil {
		return helper.Response(c, fiber.StatusNotAcceptable, nil, err.Error(), false)
	}
	
	if errors := middleware.StructValidator(user); errors != nil {
		return helper.Response(c, fiber.StatusConflict, nil, errors, false)
	}

	if err = a.authS.CreateUser(c.Context(), user); err != nil {
		return helper.Response(c, fiber.StatusConflict, nil, err.Error(), false)
	}

	return helper.Response(c, fiber.StatusOK, nil, "Register success!", true)
}
