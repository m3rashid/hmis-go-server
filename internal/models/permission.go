package models

type Resource struct {
	Name          string       `bson:"name" json:"name"`
	Description   string       `bson:"description" json:"description"`
	TotalLevelSum uint         `bson:"totalLevelSum" json:"totalLevelSum"`
	Permissions   []Permission `bson:"permissions" json:"permissions"`
}

type Permission struct {
	RoleId     uint     `bson:"roleId" json:"roleId"`
	Role       Role     `bson:"role" json:"role,omitempty"`
	ResourceID uint     `bson:"resourceId" json:"resourceId"`
	Resource   Resource `bson:"resource" json:"resource,omitempty"`
	Users      []User   `bson:"users" json:"users,omitempty"`
}

type Role struct {
	Name        string       `bson:"name" json:"name"`
	Description string       `bson:"description" json:"description"`
	User        []User       `bson:"user" json:"user,omitempty"`
	Permissions []Permission `bson:"permissions" json:"permissions,omitempty"`
}
