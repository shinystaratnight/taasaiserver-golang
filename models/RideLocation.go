package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type RideLocation struct {
	gorm.Model
	RideID   uint      `gorm:"not null"`
	Time     time.Time `gorm:"not null"`
	IsActive bool
}
