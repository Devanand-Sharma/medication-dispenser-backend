package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// Load config from environment variables
func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not found in root directory")
	}

	config := &Config{
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_NAME"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
	}

	if config.DBHost == "" || config.DBUser == "" || config.DBName == "" {
		log.Fatal("Required database environment variables are not set")
	}

	return config
}
