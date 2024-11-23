package main

import (
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/database"
)

func main() {
	db := database.Connect()
	database.Migrate(db)
}
