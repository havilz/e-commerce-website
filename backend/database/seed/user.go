package seed

import (
	"log"

	"github.com/havilz/ecommerce-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUser seeds the initial regular user if not exists
func SeedUser(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Where("role = ?", "user").Count(&count)
	if count > 0 {
		log.Println("[seed] regular user already exists, skipping...")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte("secret123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("[seed] failed to hash user password: %v", err)
	}

	user := models.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: string(hashed),
		Role:     "user",
	}

	if err := db.Create(&user).Error; err != nil {
		log.Printf("[seed] user seed failed: %v", err)
		return
	}

	log.Println("[seed] user created → email: john@example.com | password: secret123")
}
