package models

import "time"

type CartItem struct{
	ID uint `gorm:"primarykey;autoIncrement" json:"id"`
	UserID uint `gorm:"not null;index;uniqueIndex:idx_user_product" json:"user_id"`
	ProductID uint `gorm:"not null;uniqueIndex:idx_user_product" json:"product_id"`
	Quantity int `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignkey:UserID"`
	Product Product `gorm:"foreignkey:ProductID"`
}