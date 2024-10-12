package main

import (
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin Router
	r := gin.Default()

	r.GET("/medications", handlers.ListMedications)

	r.Run(":8080")
}
