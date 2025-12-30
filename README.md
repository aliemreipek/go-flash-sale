# Go Flash Sale System ⚡

![Build Status](https://github.com/aliemreipek/go-flash-sale/actions/workflows/ci.yml/badge.svg)

A high-performance, asynchronous e-commerce backend designed to handle **High Concurrency** scenarios (e.g., Flash Sales, Black Friday). This project solves the **Race Condition** problem by utilizing **RabbitMQ** for message queuing and **PostgreSQL** transactions for data integrity.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2.52-000000?style=flat&logo=gofiber)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.x-FF6600?style=flat&logo=rabbitmq)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI%2FCD-2088FF?style=flat&logo=github-actions)

## 🏗️ Architecture & Workflow

Traditional synchronous APIs fail under heavy load because the database becomes a bottleneck. This system uses an **Event-Driven Architecture**:

1.  **Ingestion (API):** The Go Fiber server receives the order request (`POST /api/orders`).
2.  **Validation:** Incoming requests are validated (Unit Tested) to ensure data integrity before processing.
3.  **Queuing (Producer):** Instead of locking the database, the request is serialized and pushed to a **RabbitMQ** queue (`orders_queue`). The user receives an immediate acknowledgement.
4.  **Processing (Worker):** A background **Goroutine (Consumer)** picks up messages from the queue sequentially.
5.  **Persistence (Atomic Transaction):**
    * The worker checks the real-time stock in **PostgreSQL**.
    * If stock > 0, it creates the order and decrements the stock within a single **Database Transaction**.
    * If the operation fails, changes are rolled back to ensure **Data Integrity**.

## 🛠️ Tech Stack

* **Language:** Golang (Go 1.21+)
* **Web Framework:** Fiber v2.52 (Fastest HTTP engine for Go)
* **Message Broker:** RabbitMQ (AMQP Protocol)
* **Database:** PostgreSQL 16
* **ORM:** GORM
* **Testing:** Go Standard Library (`testing` package)
* **CI/CD:** GitHub Actions (Automated Build & Test)
* **Containerization:** Docker & Docker Compose

## 🚀 Getting Started

### Prerequisites
* Docker & Docker Compose
* Go 1.21+ (Optional, if running locally without Docker)
* Make (Optional, for easy commands)

### 1. Clone the Repository
```bash
git clone https://github.com/aliemreipek/go-flash-sale.git
cd go-flash-sale
```

### 2. Start Infrastructure
Start PostgreSQL and RabbitMQ containers using Docker Compose:
```bash
make up
# Or: docker-compose up -d
```

### 3. Run the Application
Start the API server:
```bash
make run
# Or: go run cmd/api/main.go
```
*The server will start at `http://localhost:3000`*

---

## 🧪 Testing the System

### A. Run Unit Tests (New)
To verify the logic and validation rules, run the unit tests:
```bash
go test -v ./...
```

### B. Functional Tests (Manual)

**1. List Products**
Check the initial stock status:
```bash
curl http://localhost:3000/api/products
```

**2. Place an Order (Async)**
Simulate a user buying a product (User ID: 1, Product ID: 1):
```bash
curl -X POST http://localhost:3000/api/orders \
     -H "Content-Type: application/json" \
     -d '{"user_id": 1, "product_id": 1}'
```
*Response:* `{"status":"success", "message":"Order received! Processing in background. 🐇"}`

**3. Monitor the Queue**
You can monitor the message flow via the RabbitMQ Management Dashboard:
* **URL:** `http://localhost:15672`
* **Username:** `guest`
* **Password:** `guest`

## 📂 Project Structure

```
.
├── .github
│   └── workflows     # CI/CD Pipeline (GitHub Actions)
├── cmd
│   └── api
│       └── main.go       # Application entry point & Server config
├── internal
│   ├── controllers       # HTTP Handlers (Order & Product logic)
│   ├── database          # Database connection & Auto-migration
│   ├── models            # GORM Structs & Validators
│   ├── routes            # API Endpoint definitions
│   └── services          # RabbitMQ Producer & Consumer logic
├── pkg
│   └── configs           # Environment variables (.env) loading
├── docker-compose.yml    # Infrastructure (DB & MQ) setup
├── Makefile              # Helper commands (run, up, down, clean)
└── README.md             # Project documentation
```

## 📝 License
This project is licensed under the MIT License.