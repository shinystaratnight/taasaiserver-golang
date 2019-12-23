package models

import (
	"github.com/jinzhu/gorm"
)

type Vehicle struct {
	gorm.Model
	Name                        string                    `gorm:"not null"`
	CompanyLocationAssignment   CompanyLocationAssignment `gorm:"foreignkey:CompanyLocationAssignmentID"`
	CompanyLocationAssignmentID uint                      `gorm:"not null"`
	VehicleType                 VehicleType               `gorm:"foreignkey:VehicleTypeID"`
	VehicleTypeID               uint                      `gorm:"not null;index:idx_vehicle"`
	Brand                       string                    `gorm:"not null"`
	VehicleModel                string
	Color                       string `gorm:"not null"`
	VehicleNumber               string `gorm:"not null;unique;unique_index:idx_vehicle"`
	Image                       string
	IsActive                    bool
}
