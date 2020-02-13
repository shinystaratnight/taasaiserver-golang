package models

import (
	"github.com/jinzhu/gorm"
)

type Operator struct {
	gorm.Model
	Name       string  `gorm:"not null;unique"`
	LocationName       string  `gorm:"not null;unique"`
	Email      string  `gorm:"not null;unique"`
	Password   string  `gorm:"not null"`
	PlatformCommission float64 `gorm:"not null"`
	OperatorCommission float64 `gorm:"not null"`
	DriverWorkTime int
	DriverRestTime int
	ReferAmount float64
	ReferType int //0-flat , 1-percentage
	Currency   string
	AuthToken  string `json:"-"`
	IsActive   bool
}
