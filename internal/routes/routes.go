package routes

import (
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Configure Routes
func ConfigureRouter() *gin.Engine {
	// Gin Router
	router := gin.Default()

	// Future Middleware

	// Group Routes Under /api
	api := router.Group("/api")
	{
		// Medication Routes
		api.GET("/medications", handlers.FetchMedications)
		api.POST("/medications", handlers.CreateMedication)
		api.PATCH("/medications/:id", handlers.UpdateMedication)
		api.DELETE("/medications/:id", handlers.RemoveMedication)
	}

	return router
}
