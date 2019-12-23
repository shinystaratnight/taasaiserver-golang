package models
import "github.com/jinzhu/gorm"

type EmergencyContact struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	MobileNumber string `gorm:"not null"`
	Name string `gorm:"not null"`
	IsPassenger bool `gorm:"not null"`
	IsActive bool `gorm:"not null"`
}