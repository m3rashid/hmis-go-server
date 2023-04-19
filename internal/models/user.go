package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID           uint         `json:"userId" gorm:"type:smallint;unique"`
	Email            string       `json:"email" gorm:"type:varchar;unique"`
	Name             string       `json:"name" gorm:"type:varchar"`
	Password         string       `json:"password" gorm:"type:varchar"`
	Addresses        []Address    `json:"addresses" gorm:"many2many:user_address"`
	Roles            []Role       `json:"roles,omitempty" gorm:"many2many:user_role"`
	OtherPermissions []Permission `json:"otherPermissions,omitempty" gorm:"many2many:user_permission"`
	ProfileID        uint         `json:"profileId" gorm:"type:smallint;unique;not null;index"`
	Profile          Profile      `json:"profile,omitempty" gorm:"foreignKey:ProfileID; references:ID"`
}

type Profile struct {
	gorm.Model
	Bio            string         `json:"bio" gorm:"type:varchar"`
	RoomNumber     uint           `json:"roomNumber" gorm:"type:smallint;default:0"`
	Sex            int            `json:"sex" gorm:"type:smallint;default:0"`
	Age            uint           `json:"age" gorm:"type:smallint;default:0"`
	ContactNumber  string         `json:"contactNumber" gorm:"type:varchar"`
	MaritalStatus  uint           `json:"maritalStatus" gorm:"type:smallint;default:0"`
	BloodGroup     uint           `json:"bloodGroup" gorm:"type:smallint;default:0"`
	Origin         string         `json:"origin" gorm:"type:varchar"`
	LastVisit      time.Time      `json:"lastVisit" gorm:"type:timestamp"`
	Designation    string         `json:"designation" gorm:"type:varchar"`
	Department     string         `json:"department" gorm:"type:varchar"`
	UserId         uint           `json:"userId" gorm:"type:smallint;unique;not null;index"`
	Availabilities []Availability `json:"availabilities,omitempty" gorm:"many2many:profile_availability"`
	Appointments   []Appointment  `json:"appointments,omitempty" gorm:"many2many:profile_appointment"`
	Leaves         []Leave        `json:"leaves,omitempty" gorm:"many2many:profile_leave"`
}
