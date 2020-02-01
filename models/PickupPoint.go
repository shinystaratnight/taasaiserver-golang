package models

import (
	"github.com/jinzhu/gorm"
)

type PickupPoint struct {
	gorm.Model
	Name       string `gorm:"not null;unique"`
	ZoneID uint   `gorm:"not null"`
	IsActive   bool
}