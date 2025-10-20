package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtwhere "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/umem1125/project-management/config"
	"github.com/umem1125/project-management/controllers"
	"github.com/umem1125/project-management/utils"
)

// uc: UserController
func Setup(app *fiber.App, uc *controllers.UserController, bc *controllers.BoardController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .ev file")
	}
	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	// membuat API Group
	// JWT proctected routes
	api := app.Group("/api/v1", jwtwhere.New(jwtwhere.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c, "error unauthorized", err.Error())
		},
	}))

	userGroup := api.Group("/users")
	userGroup.Get("/page", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser) // path ----> /api/v1/users/:id
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)

	boardGroup := api.Group("/boards")
	boardGroup.Post("/", bc.CreateBoard)
}
