package models

import (
	"github.com/jinzhu/gorm"
)

type RideMessage struct {
	gorm.Model

	RideID   uint `gorm:"not null;index:idx_ride"`
	Message  string
	From int //driver - 0 , passenger - 1
	IsActive   bool
}

