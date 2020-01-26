package models

import (
	"github.com/jinzhu/gorm"
)

type Zone struct {
	gorm.Model
	Name       string `gorm:"not null;unique"`
	OperatorID uint   `gorm:"not null"`
	IsActive   bool
}
