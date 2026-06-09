package seed

import (
	"log"

	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) []models.Category {
	var count int64
	db.Model(&models.Category{}).Count(&count)

	names := []string{
		"Electronics",
		"Fashion",
		"Lifestyle",
		"Home",
		"Sports",
		"Books",
		"Beauty",
		"Gaming",
	}

	if count > 0 {
		log.Println("[seed] categories already exist, returning loaded ones...")
		var cats []models.Category
		db.Order("name ASC").Find(&cats)
		return cats
	}

	var cats []models.Category
	for _, name := range names {
		cats = append(cats, models.Category{Name: name})
	}

	if err := db.Create(&cats).Error; err != nil {
		log.Println("[seed] failed to seed categories:", err)
		return nil
	}

	log.Printf("[seed] %d sample categories created", len(cats))
	return cats
}
