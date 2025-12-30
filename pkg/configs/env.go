package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		// It's okay if .env file is missing in production (env vars might be set by OS)
		log.Println("Note: .env file not found, using system environment variables")
	}
}

// GetPort returns the port number from env or default to 3000
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "3000"
	}
	return port
}
