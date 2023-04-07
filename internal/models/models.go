package models

// TODO: Model Not Complete

import (
	"time"

	_ "gorm.io/driver/postgres"
)

type Address struct {
	BaseModel      `gorm:"embedded"`
	City           string `gorm:"not null"`
	State          string `gorm:"not null"`
	Country        string `gorm:"not null"`
	Pincode        string `gorm:"not null"`
	RoomNumber     string // `gorm:""`
	BuildingNumber string // `gorm:""`
	UserId         uint   `gorm:"unique"`
	User           User   `gorm:"foreignKey:UserId; references:Id"`
}

type Resource struct {
	BaseModel     `gorm:"embedded"`
	Name          string       `gorm:"not null"`
	Description   string       `gorm:"default:-"`
	TotalLevelSum uint         `gorm:"default:0"`
	Permissions   []Permission `gorm:"many2many:permission_resource; joinForeignKey:resource_id; joinReferences:permission_id"`
}

type Permission struct {
	BaseModel  `gorm:"embedded"`
	RoleId     uint     `gorm:"unique"` // relation to role
	Role       Role     `gorm:"foreignKey:RoleId; references:Id"`
	ResourceId uint     `gorm:"not null"` // relation to resource
	Resource   Resource `gorm:"foreignKey:ResourceId; references:Id"`
}

type Role struct {
	BaseModel   `gorm:"embedded"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"default:-"`
	IsDeleted   bool         `gorm:"default:false"`
	User        []User       `gorm:"many2many:user_role; joinForeignKey:user_role_id; joinReferences:user_id"`
	Permissions []Permission `gorm:"many2many:permission_role; joinForeignKey:permission_role_id; joinReferences:permission_id"`
}

type Availability struct {
	BaseModel `gorm:"embedded"`
	Day       uint      `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	ProfileId uint      `gorm:"not null"`
	Profile   Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type User struct {
	BaseModel `gorm:"embedded"`
	UserId    uint      `gorm:"unique"`
	Email     string    `gorm:"unique"`
	Name      string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	IsDeleted bool      `gorm:"default:false"`
	Addresses []Address `gorm:"; joinForeignKey:address_user_id; joinReferences:user_id"`
	RoleId    uint      `gorm:"not null"`
	Role      Role      `gorm:"foreignKey:RoleId; references:Id"`
	ProfileId uint      `gorm:"not null"`
	Profile   Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type Profile struct {
	BaseModel     `gorm:"embedded"`
	Bio           string        // `gorm:""`
	RoomNumber    uint          // `gorm:""`
	Sex           int           `gorm:"default:0"`
	Age           uint          `gorm:"not null"`
	ContactNumber string        `gorm:"not null"`
	MaritalStatus uint          `gorm:"default:0"`
	BloodGroup    uint          `gorm:"default:0"`
	Origin        string        // `gorm:""`
	LastVisit     time.Time     `gorm:"default:now()"`
	Designation   string        // `gorm:""`
	Department    string        // `gorm:""`
	UserId        uint          `gorm:"unique"`
	Availability  []uint        `gorm:"; joinForeignKey:availability_profile_id; joinReferences:availability_id"`
	Appointments  []Appointment `gorm:"; joinForeignKey:appointment_profile_id; joinReferences:appointment_id"`
	Leaves        []Leave       `gorm:"; joinForeignKey:leave_profile_id; joinReferences:leave_id"`
}

type Leave struct {
	BaseModel `gorm:"embedded"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Reason    string    `gorm:"not null"`
	ProfileId uint      `gorm:"not null"`
	User      Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type Appointment struct {
	BaseModel      `gorm:"embedded"`
	UserOneId      uint      `gorm:"not null"` // doctor
	UserOne        Profile   `gorm:"foreignKey:UserOneId; references:Id"`
	UserTwoId      uint      `gorm:"not null"` // patient
	UserTwo        Profile   `gorm:"foreignKey:UserTwoId; references:Id"`
	Status         uint      `gorm:"default:0"`
	DateTime       time.Time `gorm:"not null; default:now()"`
	Remarks        string    // `gorm:""`
	ReferredBy     uint      // `gorm:""`
	ReferredByUser Profile   `gorm:"foreignKey:ReferredBy; references:Id"`
	Pending        bool      `gorm:"default:true"`
}

type Consumables struct {
	BaseModel       `gorm:"embedded"`
	Name            string    `gorm:"not null"`
	Type            uint      `gorm:""`
	QuantityLeft    uint      `gorm:""`
	QuantityPerUnit uint      `gorm:"default:0"`
	BatchNumber     string    // `gorm:""`
	Manufacturer    string    // `gorm:""`
	ExpiryDate      time.Time // `gorm:""`
}

type NonConsumables struct {
	BaseModel         `gorm:"embedded"`
	Name              string        `gorm:"not null"`
	AvailableQuantity uint          `gorm:"not null"`
	Type              uint          // `gorm:""`
	ServicingTime     time.Duration // `gorm:""`
	ServicingTimeUnit uint          `gorm:"not null"`
}

type PrescribedMedicine struct {
	BaseModel      `gorm:"embedded"`
	Dosage         uint         `gorm:"not null"`
	PrescriptionId uint         `gorm:"not null"`
	Prescription   Prescription `gorm:"foreignKey:PrescriptionId; references:Id"`
	ConsumableId   uint         `gorm:"not null"`
	Consumables    Consumables  `gorm:"foreignKey:ConsumableId; references:Id"`
}

type Prescription struct {
	BaseModel     `gorm:"embedded"`
	Remarks       string      `gorm:"default:-"`
	AppointmentId uint        `gorm:"not null"`
	Appointment   Appointment `gorm:"foreignKey:AppointmentId; references:Id"`
}
