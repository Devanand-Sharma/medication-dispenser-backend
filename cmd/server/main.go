package main

import (
	"fmt"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/routes"
)

func main() {
	router := routes.ConfigureRouter()

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
