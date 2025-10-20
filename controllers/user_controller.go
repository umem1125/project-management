package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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

// c: controller
func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal", err.Error())
	}

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)
	return utils.Success(ctx, "Register Success", userResp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "invalid request", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "login failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)
	return utils.Success(ctx, "login success", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          userResp,
	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "data not found", err.Error())
	}

	var userResp models.UserResponse
	err = copier.Copy(&userResp, &user)
	if err != nil {
		return utils.BadRequest(ctx, "error bad request: internal server error", err.Error())
	}
	return utils.Success(ctx, "data berhasil ditemukan", userResp)
}
