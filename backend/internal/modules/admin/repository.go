package admin

import (
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateProduct(p *models.Product) error
	UpdateProduct(id uint, data map[string]interface{}) error
	DeleteProduct(id uint) error
	FindAllOrders(page, limit int) ([]models.Order, int64, error)
	UpdateOrderStatus(orderID uint, status string) error
	FindCategoryByName(name string) (*models.Category, error)
	FindCategoryByID(id uint) (*models.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateProduct(p *models.Product) error {
	return r.db.Create(p).Error
}

func (r *repository) UpdateProduct(id uint, data map[string]interface{}) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).Updates(data).Error
}

func (r *repository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *repository) FindAllOrders(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	r.db.Model(&models.Order{}).Count(&total)

	offset := (page - 1) * limit
	err := r.db.
		Preload("Items.Product.Category").
		Preload("User").
		Order("created_at DESC").
		Offset(offset).Limit(limit).
		Find(&orders).Error

	return orders, total, err
}

func (r *repository) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).
		UpdateColumn("status", status).Error
}

func (r *repository) FindCategoryByName(name string) (*models.Category, error) {
	var cat models.Category
	if err := r.db.Where("name = ?", name).First(&cat).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *repository) FindCategoryByID(id uint) (*models.Category, error) {
	var cat models.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}
