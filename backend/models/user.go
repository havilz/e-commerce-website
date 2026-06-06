package models

import "time"

type User struct {
	ID uint `gorm:"primarykey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role string `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}