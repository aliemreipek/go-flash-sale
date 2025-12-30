package services

import (
	"context"
	"encoding/json"
	"github.com/aliemreipek/go-flash-sale/internal/models"
	"gorm.io/gorm"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQService handles connection and publishing
type RabbitMQService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// RabbitMQ is the global variable to keep the service instance
var RabbitMQ *RabbitMQService

// ConnectRabbitMQ establishes the connection to RabbitMQ server
func ConnectRabbitMQ() {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	// 1. Open Connection
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	// 2. Open Channel (Virtual connection inside the TCP connection)
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	// 3. Declare a Queue
	// We make sure the queue exists before trying to use it
	_, err = ch.QueueDeclare(
		"orders_queue", // name
		true,           // durable (survives broker restart)
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	log.Println("Connected to RabbitMQ successfully! 🐇")

	RabbitMQ = &RabbitMQService{
		Connection: conn,
		Channel:    ch,
	}
}

// PublishOrder sends an order request to the queue
func (r *RabbitMQService) PublishOrder(orderData interface{}) error {
	// Serialize struct to JSON
	body, err := json.Marshal(orderData)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Push message to the queue
	err = r.Channel.PublishWithContext(ctx,
		"",             // exchange (default)
		"orders_queue", // routing key (queue name)
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

// StartConsumer listens to the queue and processes orders
func (r *RabbitMQService) StartConsumer(db *gorm.DB) {
	// 1. Consume messages from queue
	msgs, err := r.Channel.Consume(
		"orders_queue", // Queue name
		"",             // Consumer name (empty = auto generated)
		true,           // Auto-Ack (True = Message is deleted from queue as soon as read)
		false,          // Exclusive
		false,          // No-local
		false,          // No-wait
		nil,            // Args
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	// 2. Run in a background goroutine (Thread)
	// We use a channel to keep the main function alive if needed, but here we just loop
	go func() {
		log.Println("Consumer started. Waiting for messages...")

		for d := range msgs {
			// A. Parse the message body
			var orderRequest struct {
				UserID    int  `json:"user_id"`
				ProductID uint `json:"product_id"`
			}

			if err := json.Unmarshal(d.Body, &orderRequest); err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			log.Printf("Received Order Request for Product ID: %d", orderRequest.ProductID)

			// B. Process the Order (Database Transaction)
			// We need a transaction to ensure atomicity (Create Order + Decrease Stock)
			tx := db.Begin()

			// Check Stock
			var product models.Product
			if err := tx.First(&product, orderRequest.ProductID).Error; err != nil {
				tx.Rollback()
				log.Println("Product not found!")
				continue
			}

			if product.Stock <= 0 {
				tx.Rollback()
				log.Println("Out of stock! Order rejected.")
				// In a real app, you might notify the user here
				continue
			}

			// Create Order
			newOrder := models.Order{
				UserID:    orderRequest.UserID,
				ProductID: orderRequest.ProductID,
				Status:    "Success",
			}

			if err := tx.Create(&newOrder).Error; err != nil {
				tx.Rollback()
				log.Println("Failed to create order")
				continue
			}

			// Decrease Stock
			product.Stock = product.Stock - 1
			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				log.Println("Failed to update stock")
				continue
			}

			// Commit Transaction
			tx.Commit()
			log.Println("Order processed successfully! Stock updated. ✅")
		}
	}()
}
