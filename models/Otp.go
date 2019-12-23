package models

import "github.com/jinzhu/gorm"

type Otp struct {
	gorm.Model
	DialCode     int64  `gorm:"not null"`
	CountryCode  string `gorm:"not null"`
	MobileNumber string `gorm:"not null"`
	Otp          string `gorm:"not null"`
	IsUsed       bool
}
