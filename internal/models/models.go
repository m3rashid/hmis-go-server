package models

// TODO: Model Not Complete

import (
	"time"

	_ "gorm.io/driver/postgres"
)

type Address struct {
	BaseModel      `gorm:"embedded"`
	City           string // `gorm:""`
	State          string // `gorm:""`
	Country        string // `gorm:""`
	Pincode        string // `gorm:""`
	RoomNumber     string // `gorm:""`
	BuildingNumber string // `gorm:""`
	UserId         uint   `gorm:"unique"`
	User           User   `gorm:"foreignKey:UserId; references:Id"`
}

type Resource struct {
	BaseModel     `gorm:"embedded"`
	Name          string       // `gorm:""`
	Description   string       `gorm:"default:-"`
	TotalLevelSum uint         `gorm:"default:0"`
	Permissions   []Permission `gorm:"many2many:permission_resource"`
}

type Permission struct {
	BaseModel  `gorm:"embedded"`
	RoleId     uint     `gorm:"unique"` // relation to role
	Role       Role     `gorm:"foreignKey:RoleId; references:Id"`
	ResourceId uint     // `gorm:""`       // relation to resource
	Resource   Resource `gorm:"foreignKey:ResourceId; references:Id"`
}

type Role struct {
	BaseModel   `gorm:"embedded"`
	Name        string       // `gorm:""`
	Description string       `gorm:"default:-"`
	IsDeleted   bool         `gorm:"default:false"`
	User        []User       `gorm:"many2many:user_role"`
	Permissions []Permission `gorm:"many2many:permission_role"`
}

type Availability struct {
	BaseModel `gorm:"embedded"`
	Day       uint      // `gorm:""`
	StartTime time.Time // `gorm:""`
	EndTime   time.Time // `gorm:""`
	ProfileId uint      // `gorm:""`
	Profile   Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type User struct {
	BaseModel `gorm:"embedded"`
	UserId    uint      `gorm:"unique"`
	Email     string    `gorm:"unique"`
	Name      string    // `gorm:""`
	Password  string    // `gorm:""`
	IsDeleted bool      `gorm:"default:false"`
	Addresses []Address // `gorm:""`
	RoleId    uint      // `gorm:""`
	Role      Role      `gorm:"foreignKey:RoleId; references:Id"`
	ProfileId uint      // `gorm:""`
	Profile   Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type Profile struct {
	BaseModel     `gorm:"embedded"`
	Bio           string        // `gorm:""`
	RoomNumber    uint          // `gorm:""`
	Sex           int           `gorm:"default:0"`
	Age           uint          // `gorm:""`
	ContactNumber string        // `gorm:""`
	MaritalStatus uint          `gorm:"default:0"`
	BloodGroup    uint          `gorm:"default:0"`
	Origin        string        // `gorm:""`
	LastVisit     time.Time     // `gorm:""`
	Designation   string        // `gorm:""`
	Department    string        // `gorm:""`
	UserId        uint          `gorm:"unique"`
	Availability  []uint        // `gorm:""`
	Appointments  []Appointment // TODO
	Leaves        []Leave       // TODO
}

type Leave struct {
	BaseModel `gorm:"embedded"`
	StartDate time.Time // `gorm:""`
	EndDate   time.Time // `gorm:""`
	Reason    string    // `gorm:""`
	ProfileId uint      // `gorm:""`
	User      Profile   `gorm:"foreignKey:ProfileId; references:Id"`
}

type Appointment struct {
	BaseModel      `gorm:"embedded"`
	UserOneId      uint      // `gorm:""` // doctor
	UserOne        Profile   `gorm:"foreignKey:UserOneId; references:Id"`
	UserTwoId      uint      // `gorm:""` // patient
	UserTwo        Profile   `gorm:"foreignKey:UserTwoId; references:Id"`
	Status         uint      `gorm:"default:0"`
	DateTime       time.Time // `gorm:""`
	Remarks        string    // `gorm:""`
	ReferredBy     uint      // `gorm:""`
	ReferredByUser Profile   `gorm:"foreignKey:ReferredBy; references:Id"`
	Pending        bool      `gorm:"default:true"`
}

type Consumables struct {
	BaseModel       `gorm:"embedded"`
	Name            string    // `gorm:""`
	Type            uint      // `gorm:""`
	QuantityLeft    uint      // `gorm:""`
	QuantityPerUnit uint      `gorm:"default:0"`
	BatchNumber     string    // `gorm:""`
	Manufacturer    string    // `gorm:""`
	ExpiryDate      time.Time // `gorm:""`
}

type NonConsumables struct {
	BaseModel         `gorm:"embedded"`
	Name              string        // `gorm:""`
	AvailableQuantity uint          // `gorm:""`
	Type              uint          // `gorm:""`
	ServicingTime     time.Duration // `gorm:""`
	ServicingTimeUnit uint          // `gorm:""`
}

type PrescribedMedicine struct {
	BaseModel      `gorm:"embedded"`
	Dosage         uint         // `gorm:""`
	PrescriptionId uint         // `gorm:""`
	Prescription   Prescription `gorm:"foreignKey:PrescriptionId; references:Id"`
	ConsumableId   uint         // `gorm:""`
	Consumables    Consumables  `gorm:"foreignKey:ConsumableId; references:Id"`
}

type Prescription struct {
	BaseModel     `gorm:"embedded"`
	Remarks       string      `gorm:"default:-"`
	AppointmentId uint        // `gorm:""`
	Appointment   Appointment `gorm:"foreignKey:AppointmentId; references:Id"`
}
