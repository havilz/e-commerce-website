package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint        `gorm:"not null;index" json:"user_id"`
	TotalPrice float64     `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string      `gorm:"type:varchar(50);default:'pending';index" json:"status"`
	Address    string      `gorm:"type:text;not null" json:"address"`
	CreatedAt  time.Time   `gorm:"index" json:"created_at"`

	User  User        `gorm:"foreignKey:UserID" json:"-"`
	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
}

type OrderItem struct {
	ID              uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID         uint    `gorm:"not null;index" json:"order_id"`
	ProductID       uint    `gorm:"not null;index" json:"product_id"`
	Quantity        int     `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64 `gorm:"type:decimal(10,2);not null" json:"price_at_purchase"`

	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
