package product

import (
	"github.com/havilz/ecommerce-backend/models"
	"github.com/havilz/ecommerce-backend/pkg/security"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(search, category string, page, limit int) ([]models.Product, int64, error)
	FindByID(id uint) (*models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(search, category string, page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := r.db.Model(&models.Product{})

	if search != "" {
		escapedSearch := security.SanitizeSQLWildcards(search)
		// SQLite standard escape syntax is ESCAPE '\'
		query = query.Where("name LIKE ? ESCAPE '\\' OR description LIKE ? ESCAPE '\\'", "%"+escapedSearch+"%", "%"+escapedSearch+"%")
	}
	if category != "" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").Where("categories.name = ?", category)
	}

	query.Count(&total)

	offset := (page - 1) * limit
	if err := query.Preload("Category").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *repository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
