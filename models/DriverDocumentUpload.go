package models

import "github.com/jinzhu/gorm"

type DriverDocumentUpload struct {
	gorm.Model
	DocID     uint  `gorm:"not null"`
	DriverID  uint `gorm:"not null"`
	Image  string `gorm:"not null"`
	IsActive       bool
}
