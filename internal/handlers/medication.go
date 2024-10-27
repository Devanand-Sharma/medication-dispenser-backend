package handlers

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/database"
	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func FetchMedications(c *gin.Context) {
	// Connect to db
	db := database.Connect()

	// Define a variable to hold the medication data
	var medications []models.Medication
	// Use Preload to fetch related ScheduledTimes
	db.Preload("ScheduledTimes").Find(&medications)

	// Return the medications as JSON
	c.JSON(http.StatusOK, medications)
}

func CreateMedication(c *gin.Context) {
	// Connect to db
	db := database.Connect()

	// Define a variable to hold the medication data
	var medication models.Medication

	// Bind the JSON request to the medication model
	if err := c.ShouldBind(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the medication entry in the database
	if err := db.Create(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Print the created medication in console
	fmt.Printf("Created Medication: %v\n", medication)

	// Return the created medication as JSON
	c.JSON(http.StatusCreated, medication)
}

func UpdateMedication(c *gin.Context) {
	// Connect to db
	db := database.Connect()

	// Get the ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Fetch the existing medication
	var medication models.Medication
	if err := db.First(&medication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	// Bind the JSON request to the medication model
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the medication entry in the database
	if err := db.Save(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the updated medication as JSON
	c.JSON(http.StatusOK, medication)
}

func RemoveMedication(c *gin.Context) {
	// Connect to db
	db := database.Connect()

	// Get the ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Fetch the existing medication
	var medication models.Medication
	if err := db.First(&medication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	// Delete the medication entry from the database
	if err := db.Delete(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the deleted medication as JSON
	c.JSON(http.StatusOK, medication.ID)
}
