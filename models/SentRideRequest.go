package models

import (
	"github.com/jinzhu/gorm"
)

type SentRideRequest struct {
	gorm.Model
	DriverID uint `gorm:"not null"`
	RideID   uint `gorm:"not null"`
	IsActive bool
}
