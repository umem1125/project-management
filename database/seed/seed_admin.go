package seed

import (
	"log"

	"github.com/umem1125/project-management/config"
	"github.com/umem1125/project-management/models"
	"github.com/umem1125/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name: "Super Admin",
		Email: "admin@example.com",
		Password: password,
		Role: "admin",
	}
	if err := config.DB.FirstOrCreate(&admin,models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("User admin seeded")
	}
}