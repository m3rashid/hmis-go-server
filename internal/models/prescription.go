package models

type PrescribedMedicine struct {
	Dosage         uint          `bson:"dosage" json:"dosage"`
	PrescriptionID uint          `bson:"prescriptionId" json:"prescriptionId"`
	Prescription   Prescription  `bson:"prescription" json:"prescription,omitempty"`
	MedicineIDs    []uint        `bson:"medicineIds" json:"medicineIds"`
	Medicines      []Consumables `bson:"consumables" json:"consumables,omitempty"`
}

type Prescription struct {
	Remarks       string      `bson:"remarks" json:"remarks"`
	AppointmentID uint        `bson:"appointmentId" json:"appointmentId"`
	Appointment   Appointment `bson:"appointment" json:"appointment,omitempty"`
}
