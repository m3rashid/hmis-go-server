package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name          string       `json:"name" gorm:"type:varchar;unique;not null;index"`
	Description   string       `json:"description" gorm:"type:varchar;default:-"`
	TotalLevelSum uint         `json:"totalLevelSum" gorm:"type:smallint;default:0"`
	Permissions   []Permission `json:"permissions" gorm:"many2many:permission_resource"`
}

type Permission struct {
	gorm.Model
	RoleId     uint     `json:"roleId" gorm:"type:smallint;unique"` // relation to role
	Role       Role     `json:"role,omitempty" gorm:"foreignKey:RoleId; references:ID"`
	ResourceID uint     `json:"resourceId" gorm:"type:smallint;not null;index"` // relation to resource
	Resource   Resource `json:"resource,omitempty" gorm:"foreignKey:ResourceID; references:ID"`
	Users      []User   `json:"users,omitempty" gorm:"many2many:user_permission"`
}

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"type:varchar;unique;not null;index"`
	Description string       `json:"description" gorm:"type:varchar;default:-"`
	User        []User       `json:"user,omitempty" gorm:"many2many:user_role"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:permission_role"`
}
