package models

import (
	"github.com/jinzhu/gorm"
)

type Driver struct {
	gorm.Model
	Name                        string `gorm:"not null"`
	CompanyLocationAssignmentID uint   `gorm:"not null"`
	DialCode                    int64  `gorm:"not null"`
	MobileNumber                string `gorm:"not null"`
	LicenseNumber               string `gorm:"not null;unique"`
	AuthToken                   string
	Image                       string
	FcmID                       string
	IsActive                    bool
}
