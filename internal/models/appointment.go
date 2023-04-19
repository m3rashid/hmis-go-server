package models

import (
	"time"
)

type Appointment struct {
	UserOneID      uint      `bson:"userOneId" json:"userOneId"`
	UserOne        Profile   `bson:"userOne,omitempty" json:"userOne,omitempty"`
	UserTwoID      uint      `bson:"userTwoId" json:"userTwoId"`
	UserTwo        Profile   `bson:"userTwo,omitempty" json:"userTwo,omitempty"`
	Status         uint      `bson:"status" json:"status"`
	DateTime       time.Time `bson:"dateTime" json:"dateTime"`
	Remarks        string    `bson:"remarks" json:"remarks"`
	ReferredByID   uint      `bson:"referredById" json:"referredById"`
	ReferredByUser Profile   `bson:"referredByUser" json:"referredByUser,omitempty"`
	Pending        bool      `bson:"pending" json:"pending"`
}
