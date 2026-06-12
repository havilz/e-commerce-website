package seed

import (
	"log"

	"github.com/havilz/ecommerce-backend/config"
	"github.com/havilz/ecommerce-backend/pkg/database"
)

// Seed runs all the seeders for the application
func Seed() {
	// Load config
	cfg := config.Load()

	// Initialize database
	db := database.InitDB(cfg.DBPath)

	log.Println("[seed] starting database seeding...")
	SeedAdmin(db)
	SeedUser(db)
	cats := SeedCategories(db)
	SeedProducts(db, cats)
	log.Println("[seed] database seeding completed successfully!")
}
