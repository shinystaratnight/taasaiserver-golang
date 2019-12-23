package models

import (
	"github.com/jinzhu/gorm"
)

type SentRideRequest struct {
	gorm.Model
	Ride     Ride   `gorm:"foreignkey:RideID"`
	Driver   Driver `gorm:"foreignkey:DriverID"`
	DriverID uint   `gorm:"not null"`
	RideID   uint   `gorm:"not null"`
	IsActive bool
}
