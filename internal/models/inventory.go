package models

import "time"

type Consumables struct {
	Name            string    `bson:"name" json:"name"`
	Type            uint      `bson:"type" json:"type"`
	QuantityLeft    uint      `bson:"quantityLeft" json:"quantityLeft"`
	QuantityPerUnit uint      `bson:"quantityPerUnit" json:"quantityPerUnit"`
	BatchNumber     string    `bson:"batchNumber" json:"batchNumber"`
	Manufacturer    string    `bson:"manufacturer" json:"manufacturer"`
	ExpiryDate      time.Time `bson:"expiryDate" json:"expiryDate"`
}

type NonConsumables struct {
	Name              string `bson:"name" json:"name"`
	AvailableQuantity uint   `bson:"availableQuantity" json:"availableQuantity"`
	Type              uint   `bson:"type" json:"type"`
	ServicingTime     uint   `bson:"servicingTime" json:"servicingTime"`
	ServicingTimeUnit uint   `bson:"servicingTimeUnit" json:"servicingTimeUnit"`
}
