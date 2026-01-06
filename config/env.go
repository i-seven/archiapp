package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ImageDir string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	ImageDir = getEnv("IMAGE_DIR", "./storage/images")

	// Create directory if it doesn't exist
	if err := os.MkdirAll(ImageDir, 0755); err != nil {
		log.Printf("Warning: Could not create image directory: %v", err)
	}

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
