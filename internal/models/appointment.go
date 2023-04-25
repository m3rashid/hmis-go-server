package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorID     uint      `json:"doctorId" gorm:"not null;index"`
	PatientID    uint      `json:"patientId" gorm:"not null;index"`
	Status       uint      `json:"status" gorm:"default:0"`
	DateTime     time.Time `json:"dateTime" gorm:"type:timestamp"`
	Remarks      string    `json:"remarks" gorm:"type:varchar"`
	ReferredByID uint      `json:"referredById" gorm:"index"`
}
