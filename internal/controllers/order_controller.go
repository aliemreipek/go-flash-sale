package controllers

import (
	"github.com/aliemreipek/go-flash-sale/internal/models"
	"github.com/aliemreipek/go-flash-sale/internal/services"
	"github.com/gofiber/fiber/v2"
)

// CreateOrder handles the incoming order request and pushes it to the queue
func CreateOrder(c *fiber.Ctx) error {
	var req models.OrderRequest

	// 1. Parse JSON body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// 2. Check if the request data is valid
	if err := req.Validate(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// 3. Publish order to RabbitMQ Queue
	// We are NOT writing to the database here directly!
	err := services.RabbitMQ.PublishOrder(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to queue order",
			"error":   err.Error(),
		})
	}

	// 4. Return immediate success response
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Order received! Processing in background. 🐇",
	})
}
