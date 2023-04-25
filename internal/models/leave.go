package models

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	gorm.Model
	StartDate time.Time `json:"startDate" gorm:"type:timestamp"`
	EndDate   time.Time `json:"endDate" gorm:"type:timestamp"`
	Reason    string    `json:"reason" gorm:"type:varchar"`
	ProfileID uint      `json:"profileId" gorm:"unique;not null;index"`
}
