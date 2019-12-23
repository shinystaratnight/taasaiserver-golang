package models

import (
	"github.com/jinzhu/gorm"
)

type CompanyLocationAssignment struct {
	gorm.Model
	Company    Company  `gorm:"foreignkey:CompanyID"`
	CompanyID  uint     `gorm:"not null"`
	Location   Location `gorm:"foreignkey:LocationID"`
	LocationID uint     `gorm:"not null;index:idx_driver"`
	IsActive   bool
}
