package models

import (
	"github.com/jinzhu/gorm"
)

type VehicleCategory struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	IsActive    bool
}
