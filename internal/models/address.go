package models

type Address struct {
	City           string `bson:"city" json:"city"`
	State          string `bson:"state" json:"state"`
	Country        string `bson:"country" json:"country"`
	PinCode        string `bson:"pinCode" json:"pinCode"`
	RoomNumber     string `bson:"roomNumber" json:"roomNumber"`
	BuildingNumber string `bson:"buildingNumber" json:"buildingNumber"`
	UserID         uint   `bson:"userId" json:"userId"`
	User           User   `bson:"user" json:"user,omitempty"`
}
