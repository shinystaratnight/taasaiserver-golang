package models

import (
	"github.com/jinzhu/gorm"
)

type Passenger struct {
	gorm.Model
	Name          string `gorm:"not null"`
	DialCode      int64  `gorm:"not null"`
	CountryCode   string `gorm:"not null"`
	MobileNumber  string `gorm:"not null"`
	ReferralCode  string
	ReferredBy    string
	AuthToken     string
	Image         string
	FcmID         string
	WalletBalance float64
	IsActive      bool
}
