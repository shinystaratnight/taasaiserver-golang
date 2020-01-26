package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Ride struct {
	gorm.Model

	PassengerID   uint `gorm:"not null;index:idx_ride"`
	OperatorID    uint `gorm:"not null;index:idx_ride"`
	ZoneID        uint `sql: "default: null"`
	VehicleTypeID uint `gorm:"not null;index:idx_ride"`

	DriverID uint `gorm:"index:idx_ride",sql: "default: null"`

	PickupLocation  string  `sql: "default: null"`
	PickupLatitude  float64 `gorm:"not null"`
	PickupLongitude float64 `gorm:"not null"`

	DropLocation  string  `sql: "default: null"`
	DropLatitude  float64 `gorm:"not null"`
	DropLongitude float64 `gorm:"not null"`

	RideDateTime  time.Time `gorm:"not null"`
	RideStartTime time.Time `gorm:"not null"`
	RideEndTime   time.Time `gorm:"not null"`

	RideType    int64 `gorm:"not null"`
	IsRideLater bool  `gorm:"not null"`

	Distance         float64 `sql: "default: null"`
	Duration         float64 `sql: "default: null"`
	DurationReadable string  `sql: "default: null"`

	FareID        uint    `gorm:"index:idx_ride",sql: "default: null"`
	ZoneFareID    uint    `gorm:"index:idx_ride",sql: "default: null"`
	DistanceFare  float64 `sql: "default: null"`
	DurationFare  float64 `sql: "default: null"`
	Tax           float64 `sql: "default: null"`
	IsPaid        bool    `sql: "default: null"`
	TransactionID string  `sql: "default: null"`
	TotalFare     float64 `sql: "default: null"`

	PassengerRating float64 `sql: "default: null"`
	DriverRating    float64 `sql: "default: null"`
	PassengerReview string  `sql: "default: null"`
	DriverReview    string  `sql: "default: null"`

	RideStatus int64 `gorm:"not null"`
	IsActive   bool
}
