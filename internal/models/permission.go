package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	DisplayName string `json:"displayName" gorm:"type:varchar;not null"`
	ActualName  string `json:"actualName" gorm:"type:varchar;unique;not null;index"`
	Description string `json:"description" gorm:"type:varchar;default:-"`
	Type        string `json:"type" gorm:"type:varchar;not null;index"`
}

type Permission struct {
	gorm.Model
	DisplayName  string `json:"displayName" gorm:"type:varchar;not null"`
	ActualName   string `json:"actualName" gorm:"type:varchar;unique;not null;index"`
	Description  string `json:"description" gorm:"type:varchar;default:-"`
	ResourceType string `json:"resourceType" gorm:"type:varchar;not null"`
	Scope        string `json:"scope" gorm:"type:varchar;not null"`
	Permission   string `json:"permission" gorm:"type:varchar;not null"`
}

type Role struct {
	gorm.Model
	DisplayName string `json:"displayName" gorm:"type:varchar;not null"`
	ActualName  string `json:"actualName" gorm:"type:varchar;unique;not null;index"`
	Description string `json:"description" gorm:"type:varchar;default:-"`
	Permissions []uint `json:"permissions,omitempty" gorm:"type:int[];many2many:permission_role"`
}
