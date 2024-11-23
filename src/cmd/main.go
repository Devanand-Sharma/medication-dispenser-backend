package main

import (
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/database"
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/routes"
)

func main() {
	db := database.Connect()
	router := routes.NewRouter(db)
	engine := router.ConfigureRouter()

	engine.Run(":8080")
}
