package routes

import (
	"time"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	db     *gorm.DB
	engine *gin.Engine
}

func NewRouter(db *gorm.DB) *Router {
	engine := gin.Default()

	// Add CORS middleware with configuration for Flutter
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	return &Router{
		db:     db,
		engine: engine,
	}
}

func (r *Router) ConfigureRouter() *gin.Engine {
	medicationHandler := handlers.NewMedicationHandler(r.db)
	userHandler := handlers.NewUserHandler(r.db)

	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.POST("/sync", userHandler.SyncUser)
		}

		// Medication routes
		medications := v1.Group("/medications")
		{
			medications.GET("", medicationHandler.FetchMedications)
			medications.POST("", medicationHandler.CreateMedication)
			medications.PATCH("/:id", medicationHandler.UpdateMedication)
			medications.DELETE("/:id", medicationHandler.RemoveMedication)
		}
	}

	return r.engine
}
