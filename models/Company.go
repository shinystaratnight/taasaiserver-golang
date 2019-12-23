package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name       string  `gorm:"not null;unique"`
	Email      string  `gorm:"not null;unique"`
	Password   string  `gorm:"not null"`
	Commission float64 `gorm:"not null"`
	AuthToken  string
	IsActive   bool
}
