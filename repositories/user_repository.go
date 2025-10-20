package repositories

import (
	"strings"

	"github.com/umem1125/project-management/config"
	"github.com/umem1125/project-management/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByPublicID(publicID string) (*models.User, error)
	FindAllPagination(filter, sort string, limit, ofset int) ([]models.User, int64, error)
	Update(user *models.User) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email =  ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByPublicID(publicID string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("public_id = ?", publicID).First(&user).Error
	return &user, err
}

func (r *userRepository) FindAllPagination(filter, sort string, limit, ofset int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := config.DB.Model(&models.User{})

	// filterinng
	if filter != "" {
		filterPattern := "%" + filter + "%"
		// Ilike --> akan memfilter incasesensitive
		db = db.Where("name Ilike ? OR email Ilike ?", filterPattern, filterPattern)
	}

	// menghitung total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting
	if sort != "" {
		// sort = name (asc), sort =-name (desc)
		if sort == "-id" {
			sort = "-internal_id"

		} else if sort == "id" {
			sort = "internal_id"
		}

		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort += " ASC"
		}
		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(ofset).Find(&users).Error
	return users, total, err
}

func (r *userRepository) Update(user *models.User) error {
	return config.DB.Model(&models.User{}).Where("public_id = ?", user.PublicID).Updates(map[string]interface{}{
		"name": user.Name,
	}).Error
}
