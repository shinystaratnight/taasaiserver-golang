package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Currency string
	IsActive bool
}
