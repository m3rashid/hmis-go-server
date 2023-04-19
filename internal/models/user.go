package models

import "time"

type User struct {
	UserID           uint         `bson:"userId" json:"userId"`
	Email            string       `bson:"email" json:"email"`
	Name             string       `bson:"name" json:"name"`
	Password         string       `bson:"password" json:"password"`
	Addresses        []Address    `bson:"addresses" json:"addresses"`
	Roles            []Role       `bson:"roles" json:"roles,omitempty"`
	OtherPermissions []Permission `bson:"otherPermissions" json:"otherPermissions,omitempty"`
	ProfileID        uint         `bson:"profileId" json:"profileId"`
	Profile          Profile      `bson:"profile" json:"profile,omitempty"`
}

type Profile struct {
	Bio            string         `bson:"bio" json:"bio"`
	RoomNumber     uint           `bson:"roomNumber" json:"roomNumber"`
	Sex            int            `bson:"sex" json:"sex"`
	Age            uint           `bson:"age" json:"age"`
	ContactNumber  string         `bson:"contactNumber" json:"contactNumber"`
	MaritalStatus  uint           `bson:"maritalStatus" json:"maritalStatus"`
	BloodGroup     uint           `bson:"bloodGroup" json:"bloodGroup"`
	Origin         string         `bson:"origin" json:"origin"`
	LastVisit      time.Time      `bson:"lastVisit" json:"lastVisit"`
	Designation    string         `bson:"designation" json:"designation"`
	Department     string         `bson:"department" json:"department"`
	UserId         uint           `bson:"userId" json:"userId"`
	Availabilities []Availability `bson:"availabilities" json:"availabilities,omitempty"`
	Appointments   []Appointment  `bson:"appointments" json:"appointments,omitempty"`
	Leaves         []Leave        `bson:"leaves" json:"leaves,omitempty"`
}
