package models

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique;unique_index:idx_admin_email"`
	Password  string `gorm:"not null"`
	AuthToken string
	Image     string
	IsActive  bool `gorm:"not null"`
}
