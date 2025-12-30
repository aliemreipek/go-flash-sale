package models

import "gorm.io/gorm"

// Order struct records the successful purchases
type Order struct {
	gorm.Model
	UserID    int     `gorm:"not null" json:"user_id"` // Simplified for now (no Auth system yet)
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"` // Relationship
	Status    string  `gorm:"default:'pending'" json:"status"`     // pending, success, failed
}
