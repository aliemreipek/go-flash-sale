package controllers

import (
	"github.com/aliemreipek/go-flash-sale/internal/database"
	"github.com/aliemreipek/go-flash-sale/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts retrieves all available products
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	// 1. Fetch all products from database using Global DB instance
	result := database.DB.Find(&products)

	// 2. Check for errors
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not fetch products",
			"error":   result.Error.Error(),
		})
	}

	// 3. Return products as JSON
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   products,
	})
}
