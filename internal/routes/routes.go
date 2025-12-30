package routes

import (
	"github.com/aliemreipek/go-flash-sale/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Grouping routes under /api

	// Product Routes
	api.Get("/products", controllers.GetProducts)

	// Order Routes
	api.Post("/orders", controllers.CreateOrder)

}
