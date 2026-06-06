package seed

import (
	"log"

	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
)

// SeedProducts seeds starting products if not exists
func SeedProducts(db *gorm.DB) {
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		log.Println("[seed] products already exist, skipping...")
		return
	}

	products := []models.Product{
		{
			Name:        "Wireless Headphones",
			Description: "Premium noise-cancelling wireless headphones with 30h battery life",
			Price:       999000,
			Stock:       50,
			ImageURL:    "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400",
			Category:    "Electronics",
		},
		{
			Name:        "Mechanical Keyboard",
			Description: "TKL mechanical keyboard with Cherry MX switches and RGB backlight",
			Price:       750000,
			Stock:       30,
			ImageURL:    "https://images.unsplash.com/photo-1587829741301-dc798b83add3?w=400",
			Category:    "Electronics",
		},
		{
			Name:        "Running Shoes",
			Description: "Lightweight running shoes with cushioned sole",
			Price:       650000,
			Stock:       100,
			ImageURL:    "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=400",
			Category:    "Fashion",
		},
		{
			Name:        "Stainless Water Bottle",
			Description: "500ml double-wall insulated stainless steel water bottle",
			Price:       120000,
			Stock:       200,
			ImageURL:    "https://images.unsplash.com/photo-1602143407151-7111542de6e8?w=400",
			Category:    "Lifestyle",
		},
		{
			Name:        "Backpack 30L",
			Description: "Waterproof laptop backpack 30L with USB charging port",
			Price:       450000,
			Stock:       75,
			ImageURL:    "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=400",
			Category:    "Fashion",
		},
	}

	for i := range products {
		db.Create(&products[i])
	}

	log.Printf("[seed] %d sample products created", len(products))
}
