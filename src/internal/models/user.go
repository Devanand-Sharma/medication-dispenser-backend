package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(100);primary_key"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
