package models

import "time"

type Consumables struct {
	Name            string    `json:"name" gorm:"unique;not null;index"`
	Type            uint      `json:"type" gorm:"default:0"`
	QuantityLeft    uint      `json:"quantityLeft" gorm:"default:0"`
	QuantityPerUnit uint      `json:"quantityPerUnit" gorm:"default:0"`
	BatchNumber     string    `json:"batchNumber" gorm:"type:varchar"`
	Manufacturer    string    `json:"manufacturer" gorm:"type:varchar"`
	ExpiryDate      time.Time `json:"expiryDate" gorm:"type:timestamp"`
}

type NonConsumables struct {
	Name              string `json:"name" gorm:"unique;not null;index"`
	AvailableQuantity uint   `json:"availableQuantity" gorm:"default:0"`
	Type              uint   `json:"type" gorm:"default:0"`
	ServicingTime     uint   `json:"servicingTime" gorm:"default:0"`
	ServicingTimeUnit uint   `json:"servicingTimeUnit" gorm:"default:0"`
}
