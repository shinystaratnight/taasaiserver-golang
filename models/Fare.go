package models

import (
	"github.com/jinzhu/gorm"
)

type Fare struct {
	gorm.Model
	VehicleTypeID    uint    `gorm:"not null;"`
	OperatorID       uint    `gorm:"not null;"`
	BaseFare         float64 `gorm:"not null"`
	MinimumFare float64 `gorm:"not null"`
	WaitingTimeLimit float64 `gorm:"not null"`
	WaitingFee float64 `gorm:"not null"`
	CancellationTimeLimit float64 `gorm:"not null"`
	CancellationFee float64 `gorm:"not null"`
	DurationFare     float64 `gorm:"not null"`
	DistanceFare     float64 `gorm:"not null"`
	Tax              float64 `gorm:"not null"`
	TrafficFactor              float64 `gorm:"not null"`
	IsActive         bool
}
