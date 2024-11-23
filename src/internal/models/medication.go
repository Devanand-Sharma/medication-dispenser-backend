package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicationRoute int

const (
	Tablet MedicationRoute = iota
	Capsule
	SolutionTsp
	SolutionTbsp
	Drops
	Inhaler
	Injection
	Powder
	Other
)

type MedicationFrequency int

const (
	OnceADay MedicationFrequency = iota
	XTimesADay
	EveryXHours
	EveryXDays
	EveryXWeeks
	EveryXMonths
	OnlyAsNeeded
)

type AdministeredStatus int

const (
	Taken AdministeredStatus = iota
	Skipped
)

type ScheduledTime struct {
	gorm.Model
	Time time.Time `json:"time" gorm:"not null"`

	// Many-to-one relationship with Medication
	MedicationID uint `json:"medication_id" gorm:"not null"`
}

type AdministeredTime struct {
	gorm.Model
	Time   time.Time          `gorm:"not null"`
	Status AdministeredStatus `gorm:"not null"`

	// Many-to-one relationship with Medication
	MedicationID uint `gorm:"not null;unique"`
}

type RefillDate struct {
	gorm.Model
	Date time.Time `gorm:"not null"`

	// Many-to-one relationship with Medication
	MedicationID uint `gorm:"not null;unique"`
}

type Medication struct {
	gorm.Model
	Name              string              `json:"name" binding:"required" gorm:"not null"`
	Condition         string              `json:"condition" binding:"required" gorm:"not null"`
	Route             MedicationRoute     `json:"route" binding:"required" gorm:"not null"`
	Dose              int                 `json:"dose" binding:"required,min=1" gorm:"not null"`
	TotalQuantity     int                 `json:"total_quantity" binding:"required,min=1" gorm:"not null"`
	RemainingQuantity int                 `json:"remaining_quantity" binding:"required" gorm:"not null"`
	ThresholdQuantity int                 `json:"threshold_quantity" binding:"required" gorm:"not null"`
	IsRefillReminder  bool                `json:"is_refill_reminder" binding:"required" gorm:"not null"`
	Frequency         MedicationFrequency `json:"frequency" binding:"gte=0,lte=6" gorm:"not null"`
	FrequencyCount    *int                `json:"frequency_count" binding:"omitempty,min=1"`
	StartDate         time.Time           `json:"start_date" binding:"required" gorm:"not null"`
	EndDate           *time.Time          `json:"end_date"`
	IsReminder        bool                `json:"is_reminder" binding:"required" gorm:"not null"`
	Instructions      string              `json:"instructions"`

	// One-to-many relationships
	ScheduledTimes    []ScheduledTime    `json:"scheduled_times" gorm:"constraint:OnDelete:CASCADE;"`
	AdministeredTimes []AdministeredTime `json:"administered_times" gorm:"constraint:OnDelete:CASCADE;"`
	RefillDates       []RefillDate       `json:"refill_dates" gorm:"constraint:OnDelete:CASCADE;"`
}
