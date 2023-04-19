package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	City           string `json:"city" gorm:"type:varchar"`
	State          string `json:"state" gorm:"type:varchar"`
	Country        string `json:"country" gorm:"type:varchar"`
	PinCode        string `json:"pinCode" gorm:"type:varchar"`
	RoomNumber     string `json:"roomNumber" gorm:"type:varchar"`
	BuildingNumber string `json:"buildingNumber" gorm:"type:varchar"`
	UserID         uint   `json:"userId" gorm:"type:smallint;unique;not null;index"`
	User           User   `json:"user,omitempty" gorm:"foreignKey:UserID; references:ID"`
}
