package cart

import (
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindByUserID(userID uint) ([]models.CartItem, error)
	FindItemByID(itemID, userID uint) (*models.CartItem, error)
	FindByUserAndProduct(userID, productID uint) (*models.CartItem, error)
	Create(item *models.CartItem) error
	Update(item *models.CartItem) error
	Delete(itemID, userID uint) error
	ClearByUserID(userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByUserID(userID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	if err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *repository) FindItemByID(itemID, userID uint) (*models.CartItem, error) {
	var item models.CartItem
	if err := r.db.Preload("Product").Where("id = ? AND user_id = ?", itemID, userID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *repository) FindByUserAndProduct(userID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *repository) Create(item *models.CartItem) error {
	return r.db.Create(item).Error
}

func (r *repository) Update(item *models.CartItem) error {
	return r.db.Save(item).Error
}

func (r *repository) Delete(itemID, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", itemID, userID).Delete(&models.CartItem{}).Error
}

func (r *repository) ClearByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}
