package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        uint   `json:"userId" gorm:"unique"`
	Name          string `json:"name" gorm:"type:varchar"`
	Email         string `json:"email" gorm:"type:varchar;unique"`
	EmailVerified bool   `json:"emailVerified" gorm:"type:boolean;default:false"`
	Password      string `json:"password" gorm:"type:varchar"`
	Roles         []uint `json:"roles,omitempty" gorm:"type:int[];many2many:user_role"`
	ProfileID     uint   `json:"profileId" gorm:"unique;not null;index"`
}

type Profile struct {
	gorm.Model
	Bio                      string    `json:"bio" gorm:"type:varchar"`
	RoomNumber               uint      `json:"roomNumber" gorm:"default:0"`
	Age                      uint      `json:"age" gorm:"default:0"`
	Sex                      int       `json:"sex" gorm:"default:0"`
	Phone                    string    `json:"phone" gorm:"type:varchar"`
	PhoneVerified            bool      `json:"phoneVerified" gorm:"type:boolean;default:false"`
	MaritalStatus            uint      `json:"maritalStatus" gorm:"default:0"`
	ProfilePicture           string    `json:"profilePicture" gorm:"type:varchar"`
	Addresses                []uint    `json:"addresses,omitempty" gorm:"type:int[];many2many:profile_address"`
	BloodGroup               uint      `json:"bloodGroup" gorm:"default:0"`
	Origin                   string    `json:"origin" gorm:"type:varchar"`
	LastVisit                time.Time `json:"lastVisit" gorm:"type:timestamp"`
	Designation              string    `json:"designation" gorm:"type:varchar"`
	Department               string    `json:"department" gorm:"type:varchar"`
	UserHealthId             uint      `json:"userHealthId" gorm:"unique;not null;index"`
	Leaves                   []uint    `json:"leaves,omitempty" gorm:"type:int[];many2many:profile_leave"`
	Availabilities           []uint    `json:"availabilities,omitempty" gorm:"type:int[];many2many:profile_availability"`
	AppointmentsAsDoctor     []uint    `json:"appointmentsAsDoctor,omitempty" gorm:"type:int[];many2many:profile_appointment_doctor"`
	AppointmentsAsPatient    []uint    `json:"appointmentsAsPatient,omitempty" gorm:"type:int[];many2many:profile_appointment_patient"`
	AppointmentsAsReferredBy []uint    `json:"appointmentsAsReferredBy,omitempty" gorm:"type:int[];many2many:profile_appointment_patient"`
}
