package order

import (
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order *models.Order) error
	FindByUserID(userID uint) ([]models.Order, error)
	FindByID(orderID, userID uint) (*models.Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *repository) FindByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.
		Preload("Items.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}

func (r *repository) FindByID(orderID, userID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.
		Preload("Items.Product").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
