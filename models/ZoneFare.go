package models

import (
	"github.com/jinzhu/gorm"
)

type ZoneFare struct {
	gorm.Model
	VehicleType      VehicleType `gorm:"foreignkey:VehicleTypeID"`
	Zone             Zone        `gorm:"foreignkey:ZoneID"`
	VehicleTypeID    uint        `gorm:"not null;index:idx_fare"`
	ZoneID           uint        `gorm:"not null;index:idx_fare"`
	BaseFare         float64     `gorm:"not null"`
	BaseFareDistance float64     `gorm:"not null"`
	BaseFareDuration float64     `gorm:"not null"`
	DurationFare     float64     `gorm:"not null"`
	DistanceFare     float64     `gorm:"not null"`
	Tax              float64     `gorm:"not null"`
	IsActive         bool
}
