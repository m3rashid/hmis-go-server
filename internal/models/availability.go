package models

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	Day       uint      `json:"day" gorm:"type:smallint"`
	StartTime time.Time `json:"startTime" gorm:"type:timestamp"`
	EndTime   time.Time `json:"endTime" gorm:"type:timestamp"`
	ProfileID uint      `json:"profileId" gorm:"not null;index"`
}
