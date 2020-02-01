package models

import (
	"github.com/jinzhu/gorm"
)

type Driver struct {
	gorm.Model

	Name         string `gorm:"not null"`
	DialCode     int64  `gorm:"not null"`
	MobileNumber string `gorm:"not null"`

	OperatorID int `gorm:"not null"`

	VehicleName   string `gorm:"not null"`
	VehicleTypeID uint   `gorm:"not null;index:idx_vehicle"`
	VehicleBrand  string `gorm:"not null"`
	VehicleModel  string
	VehicleColor  string `gorm:"not null"`
	VehicleNumber string `gorm:"not null;unique;unique_index:idx_vehicle"`
	VehicleImage  string

	AuthToken   string
	DriverImage string
	FcmID       string

	IsProfileCompleted bool
	IsOnline bool
	IsRide bool

	IsActive bool
}
