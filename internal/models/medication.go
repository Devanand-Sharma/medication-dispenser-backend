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
	Name              string              `json:"name" gorm:"not null"`
	Condition         string              `json:"condition" gorm:"not null"`
	Route             MedicationRoute     `json:"route" gorm:"not null"`
	Dose              int                 `json:"dose" gorm:"not null"`
	TotalQuantity     int                 `json:"total_quantity" gorm:"not null"`
	RemainingQuantity int                 `json:"remaining_quantity" gorm:"not null"`
	ThresholdQuantity int                 `json:"threshold_quantity" gorm:"not null"`
	IsRefillReminder  bool                `json:"is_refill_reminder" gorm:"not null"`
	Frequency         MedicationFrequency `json:"frequency" gorm:"not null"`
	FrequencyCount    *int                `json:"frequency_count"`
	StartDate         time.Time           `json:"start_date" gorm:"not null"`
	EndDate           *time.Time          `json:"end_date"`
	IsReminder        bool                `json:"is_reminder" gorm:"not null"`
	Instructions      string              `json:"instructions"`

	// One-to-many relationship with RefillDate
	ScheduledTimes    []ScheduledTime    `json:"scheduled_times" gorm:"constraint:OnDelete:CASCADE;"`
	AdministeredTimes []AdministeredTime `json:"administered_times" gorm:"constraint:OnDelete:CASCADE;"`
	RefillDates       []RefillDate       `json:"refill_dates" gorm:"constraint:OnDelete:CASCADE;"`
}
