package database

import (
	"fmt"
	"log"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/models"
	"github.com/Devanand-Sharma/medication-dispenser-backend/pkg/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to the database and return instance
func Connect() *gorm.DB {
	// Load Config
	cfg := database.LoadConfig()

	// Print Config
	// fmt.Printf("DBUser: %v, DBPassword: %v, DBName: %v, DBHost: %v, DBPort: %v\n", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	// Create DSN String
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Toronto", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v", err)
	}

	return db
}

// Migrate schemas to the database
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Medication{},
		&models.Prescription{},
		&models.RefillDate{},
		&models.Dosage{},
		&models.AdministeredTime{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
