package category

import (
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	Create(cat *models.Category) error
	Update(cat *models.Category) error
	Delete(id uint) error
	FindAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
	FindByName(name string) (*models.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(cat *models.Category) error {
	return r.db.Create(cat).Error
}

func (r *repository) Update(cat *models.Category) error {
	return r.db.Save(cat).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

func (r *repository) FindAll() ([]models.Category, error) {
	var cats []models.Category
	err := r.db.Order("name ASC").Find(&cats).Error
	return cats, err
}

func (r *repository) FindByID(id uint) (*models.Category, error) {
	var cat models.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *repository) FindByName(name string) (*models.Category, error) {
	var cat models.Category
	if err := r.db.Where("name = ?", name).First(&cat).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}
