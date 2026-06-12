package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/havilz/ecommerce-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dbpath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("[database] failed to connect to sqlite : %v", err)
	}

	if err := db.AutoMigrate(models.AllModels()...); err != nil {
		log.Fatalf("[database] failed to auto-migrate: %v", err)
	}

	// Data migration for category text column to category_id
	if db.Migrator().HasColumn(&models.Product{}, "category") {
		log.Println("[database] legacy category column exists. Checking for data migration...")
		var products []struct {
			ID       uint
			Category string
		}
		if err := db.Table("products").
			Select("id, category").
			Where("(category_id = 0 OR category_id IS NULL) AND category IS NOT NULL AND category != ''").
			Find(&products).Error; err == nil && len(products) > 0 {
			
			log.Printf("[database] found %d products to migrate...", len(products))
			for _, p := range products {
				var cat models.Category
				if err := db.Where("name = ?", p.Category).First(&cat).Error; err != nil {
					cat = models.Category{Name: p.Category}
					if err := db.Create(&cat).Error; err != nil {
						log.Printf("[database] failed to auto-create category %s: %v", p.Category, err)
						continue
					}
				}
				if err := db.Table("products").Where("id = ?", p.ID).Update("category_id", cat.ID).Error; err != nil {
					log.Printf("[database] failed to update product ID %d: %v", p.ID, err)
				}
			}
			log.Println("[database] legacy category migration completed successfully")
		}
	}

	log.Println("[database] SQLite connected and migrated successfully")
	return db
}