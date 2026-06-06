package models

import "time"

type Product struct {
	ID uint `gorm:"primarykey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(200);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock int `gorm:"not null;default:0" json:"stock"`
	ImageURL string `gorm:"type:varchar(500)" json:"image_url"`
	Category string `gorm:"type:varchar(100)" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}