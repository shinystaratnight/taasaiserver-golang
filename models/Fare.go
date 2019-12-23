package models

import (
	"github.com/jinzhu/gorm"
)

type Fare struct {
	gorm.Model
	VehicleType      VehicleType `gorm:"foreignkey:VehicleTypeID"`
	Location         Location    `gorm:"foreignkey:LocationID"`
	VehicleTypeID    uint        `gorm:"not null;index:idx_fare"`
	LocationID       uint        `gorm:"not null;index:idx_fare"`
	BaseFare         float64     `gorm:"not null"`
	BaseFareDistance float64     `gorm:"not null"`
	BaseFareDuration float64     `gorm:"not null"`
	DurationFare     float64     `gorm:"not null"`
	DistanceFare     float64     `gorm:"not null"`
	Tax              float64     `gorm:"not null"`
	IsActive         bool
}
