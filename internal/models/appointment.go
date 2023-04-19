package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UserOneID      uint      `json:"userOneId" gorm:"type:smallint;not null;index"`
	UserOne        Profile   `json:"userOne,omitempty" gorm:"foreignKey:UserOneID; references:ID"`
	UserTwoID      uint      `json:"userTwoId" gorm:"type:smallint;not null;index"`
	UserTwo        Profile   `json:"userTwo,omitempty" gorm:"foreignKey:UserTwoID; references:ID"`
	Status         uint      `json:"status" gorm:"type:smallint;default:0"`
	DateTime       time.Time `json:"dateTime" gorm:"type:timestamp"`
	Remarks        string    `json:"remarks" gorm:"type:varchar"`
	ReferredByID   uint      `json:"referredById" gorm:"type:smallint;index"`
	ReferredByUser Profile   `json:"referredByUser,omitempty" gorm:"foreignKey:ReferredByID; references:ID"`
	Pending        bool      `json:"pending" gorm:"type:boolean;default:true"`
}
