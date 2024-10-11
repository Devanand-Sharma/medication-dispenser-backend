package main

import (
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/database"
)

func main() {
	// Connect to the database
	db := database.Connect()

	// Migrate the schema
	database.Migrate(db)
}
