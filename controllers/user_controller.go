package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umem1125/project-management/models"
	"github.com/umem1125/project-management/services"
	"github.com/umem1125/project-management/utils"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal", err.Error())
	}

	return utils.Success(ctx, "Register Success", user)
}