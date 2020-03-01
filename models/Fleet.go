package models

import (
	"github.com/jinzhu/gorm"
)

type Fleet struct {
	gorm.Model
	Name       string `gorm:"not null"`
	OperatorID string `gorm:"not null"`
	Email      string `gorm:"not null;unique"`
	Password   string `gorm:"not null"`
	IsActive   bool
}
