package database

import (
	"fmt"
	"github.com/aliemreipek/go-flash-sale/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB holds the database connection instance
var DB *gorm.DB

// ConnectDB establishes connection to the PostgreSQL database
func ConnectDB() {
	// 1. Build Data Source Name (DSN) from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// 2. Open connection using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Show SQL queries in console
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected to Database successfully! ✅")

	log.Println("Running Migrations...")
	// This creates "products" and "orders" tables in PostgreSQL automatically
	err = db.AutoMigrate(&models.Product{}, &models.Order{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err)
	}
	log.Println("Migrations Completed! 🚀")

	// Seed data if empty
	SeedProducts(db)

	// 3. Assign connection to global variable
	DB = db
}
