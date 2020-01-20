package models

import (
	"github.com/jinzhu/gorm"
)

type RideEventLog struct {
	gorm.Model

	RideID     uint   `gorm:"not null;index:idx_ride"`
	RideStatus int64  `gorm:"not null"`
	Message    string `gorm:"not null"`
	IsActive   bool
}
