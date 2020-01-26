package models

import (
	"github.com/jinzhu/gorm"
)

type VehicleType struct {
	gorm.Model
	Name              string          `gorm:"not null;unique"`
	Image             string          `gorm:"not null"`
	VehicleCategoryID uint            `gorm:"not null"`
	Description       string          `gorm:"not null"`
	ImageActive       string          `gorm:"not null"`
	SeatCapacity      int64           `gorm:"not null"`
	IsActive          bool
}
