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
		Preload("Items.Product").
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
