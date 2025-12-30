package controllers

import (
	"github.com/aliemreipek/go-flash-sale/internal/services"
	"github.com/gofiber/fiber/v2"
)

// OrderRequest DTO (Data Transfer Object) for incoming JSON
type OrderRequest struct {
	UserID    int  `json:"user_id"`
	ProductID uint `json:"product_id"`
}

// CreateOrder handles the incoming order request and pushes it to the queue
func CreateOrder(c *fiber.Ctx) error {
	var request OrderRequest

	// 1. Parse JSON body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// 2. Publish order to RabbitMQ Queue
	// We are NOT writing to the database here directly!
	err := services.RabbitMQ.PublishOrder(request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to queue order",
			"error":   err.Error(),
		})
	}

	// 3. Return immediate success response
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Order received! Processing in background. 🐇",
	})
}
