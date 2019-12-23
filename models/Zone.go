package models

import (
	"github.com/jinzhu/gorm"
)

type Zone struct {
	gorm.Model
	Name       string   `gorm:"not null;unique"`
	Location   Location `gorm:"foreignkey:LocationID"`
	LocationID uint     `gorm:"not null"`
	IsActive   bool
}
