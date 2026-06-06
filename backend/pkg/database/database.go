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

	log.Println("[database] SQLite connected and migrated successfully")
	return db
}