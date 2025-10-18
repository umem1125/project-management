// Bussines Logic Here
package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/umem1125/project-management/models"
	"github.com/umem1125/project-management/repositories"
	"github.com/umem1125/project-management/utils"
)

type UserService interface {
	Register(user *models.User) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{}
}

// s: service
func (s *userService)Register(user *models.User) error {
	// cek email sudah tercaftar atau belum
	// hashing password
	// set role
	// simpan user

	existingUser , _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID != 0 {
		return errors.New("email already registered")
	}
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()
	return s.repo.Create(user)
}