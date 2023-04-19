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
	ProfileID uint      `json:"profileId" gorm:"type:smallint;unique;not null;index"`
	Profile   Profile   `json:"profile,omitempty" gorm:"foreignKey:ProfileID; references:ID"`
}
