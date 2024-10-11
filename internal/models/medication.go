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
	Time     time.Time `gorm:"not null"`
	DosageID uint
}

type AdministeredTime struct {
	gorm.Model
	Time       time.Time          `gorm:"not null"`
	Status     AdministeredStatus `gorm:"not null"`
	IsReminder bool               `gorm:"not null"`
	DosageID   uint
}

type Dosage struct {
	gorm.Model
	Frequency         MedicationFrequency `gorm:"not null"`
	FrequencyCount    int
	ScheduledTimes    []ScheduledTime
	AdministeredTimes []AdministeredTime
	StartDate         time.Time `gorm:"not null"`
	EndDate           time.Time
	IsReminder        bool
	MedicationID      uint
}

type RefillDate struct {
	gorm.Model
	Date           time.Time `gorm:"not null"`
	PrescriptionID uint
}

type Prescription struct {
	gorm.Model
	TotalQuantity     int `gorm:"not null"`
	RemainingQuantity int `gorm:"not null"`
	ThresholdQuantity int `gorm:"not null"`
	RefillDates       []RefillDate
	IsRefillReminder  bool `gorm:"not null"`
	MedicationID      uint
}

type Medication struct {
	gorm.Model
	Name         string          `gorm:"not null"`
	Condition    string          `gorm:"not null"`
	Route        MedicationRoute `gorm:"not null"`
	Dose         int             `gorm:"not null"`
	Prescription Prescription
	Dosage       Dosage
	Instructions string
	Image        string
	Color        string
}
