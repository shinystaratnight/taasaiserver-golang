package models

import "github.com/jinzhu/gorm"

type DriverDocument struct {
	gorm.Model
	OperatorID     uint  `gorm:"not null"`
	Name  string `gorm:"not null"`
	IsActive       bool
}
