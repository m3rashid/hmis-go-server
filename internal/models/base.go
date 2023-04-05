package models

import (
	_ "gorm.io/gorm"
)

type BaseModel struct {
	Id        uint  `gorm:"primaryKey; autoIncrement; not null"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
