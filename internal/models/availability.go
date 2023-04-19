package models

import "time"

type Availability struct {
	Day       uint      `bson:"day" json:"day"`
	StartTime time.Time `bson:"startTime" json:"startTime"`
	EndTime   time.Time `bson:"endTime" json:"endTime"`
	ProfileID uint      `bson:"profileId" json:"profileId"`
	Profile   Profile   `bson:"profile" json:"profile,omitempty"`
}
