package models

import "gorm.io/gorm"

type PrescribedMedicine struct {
	gorm.Model
	Dosage         uint         `json:"dosage" gorm:"type:varchar;default:0"`
	PrescriptionID uint         `json:"prescriptionId" gorm:"not null;index"`
	Prescription   Prescription `json:"prescription,omitempty" gorm:"foreignKey:PrescriptionID;references:ID"`
	MedicineIDs    []uint       `json:"medicineIds" gorm:"type:int[];not null;index"`
	Medicines      []uint       `json:"consumables,omitempty" gorm:"type:int[];many2many:prescribed_medicines"`
}

type Prescription struct {
	gorm.Model
	Remarks       string `json:"remarks" gorm:"type:varchar;default:-"`
	AppointmentID uint   `json:"appointmentId" gorm:"not null;index"`
}
