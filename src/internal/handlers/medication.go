package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Devanand-Sharma/medication-dispenser-backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MedicationHandler struct {
	db *gorm.DB
}

func NewMedicationHandler(db *gorm.DB) *MedicationHandler {
	return &MedicationHandler{db: db}
}

func (h *MedicationHandler) FetchMedications(c *gin.Context) {
	var medications []models.Medication
	h.db.Preload("ScheduledTimes").
		Preload("AdministeredTimes").
		Preload("RefillDates").
		Find(&medications)
	c.JSON(http.StatusOK, medications)
}

func (h *MedicationHandler) CreateMedication(c *gin.Context) {
	var medication models.Medication

	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Validation failed: %v", err.Error()),
		})
		return
	}

	if err := h.db.Create(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Database error: %v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, medication)
}

func (h *MedicationHandler) UpdateMedication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var medication models.Medication
	if err := h.db.First(&medication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)
}

func (h *MedicationHandler) RemoveMedication(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var medication models.Medication
	if err := h.db.First(&medication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	if err := h.db.Delete(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication.ID)
}
