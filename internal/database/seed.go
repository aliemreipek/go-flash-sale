package database

import (
	"log"

	"github.com/aliemreipek/go-flash-sale/internal/models"
	"gorm.io/gorm"
)

// SeedProducts inserts dummy data if the database is empty
func SeedProducts(db *gorm.DB) {
	var count int64
	db.Model(&models.Product{}).Count(&count)

	if count > 0 {
		return // Data already exists, do nothing
	}

	products := []models.Product{
		{
			Name:       "iPhone 15 Pro Max",
			Image:      "iphone15.jpg",
			Price:      1200.00,
			Stock:      100, // Only 100 items available!
			TotalStock: 100,
		},
		{
			Name:       "PlayStation 5",
			Image:      "ps5.jpg",
			Price:      500.00,
			Stock:      50,
			TotalStock: 50,
		},
		{
			Name:       "AirPods Pro 2",
			Image:      "airpods.jpg",
			Price:      250.00,
			Stock:      200,
			TotalStock: 200,
		},
	}

	if err := db.Create(&products).Error; err != nil {
		log.Printf("Failed to seed products: %v\n", err)
	} else {
		log.Println("Database seeded successfully! 🌱")
	}
}
