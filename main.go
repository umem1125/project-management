package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/umem1125/project-management/config"
	"github.com/umem1125/project-management/controllers"
	"github.com/umem1125/project-management/database/seed"
	"github.com/umem1125/project-management/repositories"
	"github.com/umem1125/project-management/routes"
	"github.com/umem1125/project-management/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is runninng on port : ", port)
	log.Fatal(app.Listen(":" + port))

}