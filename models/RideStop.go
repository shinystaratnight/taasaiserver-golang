package models

import (
	"github.com/jinzhu/gorm"
)

type RideStop struct {
	gorm.Model

	RideID   uint `gorm:"not null;index:idx_ride"`

	Location  string
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`

	IsReached bool

	IsActive   bool
}

