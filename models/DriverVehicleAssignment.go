package models

import (
	"github.com/jinzhu/gorm"
)

type DriverVehicleAssignment struct {
	gorm.Model
	Driver    Driver  `gorm:"foreignkey:DriverID"`
	Vehicle   Vehicle `gorm:"foreignkey:VehicleID"`
	DriverID  uint    `gorm:"not null"`
	VehicleID uint    `gorm:"not null"`
	IsOnline  bool
	IsRide    bool
	IsActive  bool
}
