package handlers

import (
	"fmt"
	"net/http"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/database"
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func ListMedications(c *gin.Context) {
	db := database.Connect()

	var medications []models.Medication
	db.Find(&medications)

	// Print the medications in console
	for _, medication := range medications {
		fmt.Println(medication)
	}

	// Return the medications as JSON
	c.JSON(http.StatusOK, medications)
}
