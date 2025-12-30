package main

import (
	"log"

	"github.com/aliemreipek/go-flash-sale/internal/database"
	"github.com/aliemreipek/go-flash-sale/internal/routes"
	"github.com/aliemreipek/go-flash-sale/internal/services"
	"github.com/aliemreipek/go-flash-sale/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// 1. Load config
	configs.LoadEnv()

	// 2. Connect to Infrastructure
	database.ConnectDB()
	services.ConnectRabbitMQ()

	// 3. Start the background consumer worker
	// We pass the database connection so the worker can write to DB
	services.RabbitMQ.StartConsumer(database.DB)

	// 4. Initialize Fiber app
	app := fiber.New()

	// 5. Enable CORS (Cross-Origin Resource Sharing)
	// This is required for the frontend to communicate with the backend
	app.Use(cors.New())

	// 6. Register Routes
	routes.SetupRoutes(app)

	// 7. Start Server
	log.Fatal(app.Listen(":" + configs.GetPort()))
}
