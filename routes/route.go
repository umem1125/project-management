package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/umem1125/project-management/controllers"
)

// uc: UserController
func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .ev file")
	}
	app.Post("/v1/auth/register", uc.Register)
}