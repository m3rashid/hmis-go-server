package profile

import (
	"github.com/m3rashid-org/hmis-go-server/modules/auth"
	"gorm.io/gorm"
)

type sexType string

const (
	MALE   sexType = "MALE"
	FEMALE sexType = "FEMALE"
	OTHERS sexType = "OTHERS"
)

type maritalStatus string

const (
	SINGLE   maritalStatus = "SINGLE"
	MARRIED  maritalStatus = "MARRIED"
	DIVORCED maritalStatus = "DIVORCED"
	WIDOWED  maritalStatus = "WIDOWED"
)

type role string

const (
	DOCTOR        role = "DOCTOR"
	ADMIN         role = "ADMIN"
	RECEPTIONIST  role = "RECEPTIONIST"
	PHARMACIST    role = "PHARMACIST"
	INVENTORY_MGR role = "INVENTORY_MGR"
	CO_ADMIN      role = "CO_ADMIN"
	OTHER         role = "OTHER"
	PATIENT       role = "PATIENT"
)

type ProfileCommons struct {
	gorm.Model
	auth.User
	auth.Address
	Contact       string
	Sex           sexType       `gorm:"type:sex"`
	Role          role          `gorm:"type:role"`
	MaritalStatus maritalStatus `gorm:"type:maritalStatus"`
}

type UserProfile struct {
	ProfileCommons
	Designation string
}

type PatientProfile struct {
	ProfileCommons
	// visits
}
