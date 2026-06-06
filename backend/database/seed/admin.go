package seed

import (
	"log"

	"github.com/havilz/ecommerce-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedAdmin seeds the initial admin user if not exists
func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&count)
	if count > 0 {
		log.Println("[seed] admin already exists, skipping...")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("[seed] failed to hash admin password: %v", err)
	}

	admin := models.User{
		Name:     "Administrator",
		Email:    "admin@ecommerce.com",
		Password: string(hashed),
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("[seed] admin seed failed: %v", err)
		return
	}

	log.Println("[seed] admin created → email: admin@ecommerce.com | password: admin123")
}
