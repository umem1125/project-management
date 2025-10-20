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
	Login(email, password string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByPublicID(id string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// s: service
func (s *userService) Register(user *models.User) error {
	// cek email sudah tercaftar atau belum
	// hashing password
	// set role
	// simpan user

	existingUser, _ := s.repo.FindByEmail(user.Email)
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

func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credential")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid Credentials")
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByPublicID(id string) (*models.User, error) {
	return s.repo.FindByPublicID(id)
}
