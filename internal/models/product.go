package models

import "gorm.io/gorm"

// Product struct represents the items available for flash sale
type Product struct {
	gorm.Model
	Name       string  `gorm:"not null" json:"name"`
	Image      string  `json:"image"`
	Price      float64 `gorm:"not null" json:"price"`
	Stock      int     `gorm:"not null;default:0" json:"stock"`       // Current available stock
	TotalStock int     `gorm:"not null;default:0" json:"total_stock"` // Initial stock for tracking progress
}
