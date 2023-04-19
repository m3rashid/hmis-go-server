package models

import "time"

type Leave struct {
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `bson:"endDate" json:"endDate"`
	Reason    string    `bson:"reason" json:"reason"`
	ProfileID uint      `bson:"profileId" json:"profileId"`
	User      Profile   `bson:"profile" json:"profile,omitempty"`
}
