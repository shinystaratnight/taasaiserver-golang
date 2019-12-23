package models

import (
	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	Name            string  `gorm:"not null"`
	DriverCountFrom int64   `gorm:"not null"`
	DriverCountTo   int64   `gorm:"not null"`
	PricePerDriver  float64 `gorm:"not null"`
	IsFree          bool    `gorm:"not null"`
	IsActive        bool    `gorm:"not null"`
}
