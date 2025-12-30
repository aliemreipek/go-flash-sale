# Project Variables
APP_NAME=flashsale

# 1. Run the Go Application
run:
	@echo "Starting Go Application..."
	go run cmd/api/main.go

# 2. Start Infrastructure (Docker)
up:
	@echo "Starting Docker Containers..."
	docker-compose up -d

# 3. Stop Infrastructure (Docker)
down:
	@echo "Stopping Docker Containers..."
	docker-compose down

# 4. Connect to Database (Windows/GitBash Compatible)
db:
	@echo "Connecting to Database..."
	winpty docker exec -it flashsale_db psql -U postgres -d flashsale

# 5. Tail RabbitMQ Logs
logs-mq:
	docker logs -f flashsale_rabbitmq

# 6. Clean Everything (Including Data Volumes) - WARNING
clean:
	docker-compose down -v
	@echo "All data cleared! System is clean."