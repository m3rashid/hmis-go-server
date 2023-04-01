package utils

import (
	"gorm.io/gorm"
)

type Resource struct {
	Name string
	AvailablePermissions int
}

type Permission struct {
	ResourceId int
	Level int
}

type Role struct {
	gorm.Model
	Permissions []Permission `gorm:"type:integer[]"`
}

/**
 * Idea is taken from how linux filesystem handles user permissions 
	* 1 for read
	* 2 for write
	* 4 for edit
	* 8 for delete

	* The final permission is the sum of all its permissions
	* If the user has read and edit permissions, level is 1+4=5
	* 
	* ========== Algorithm ========== 
	*	subtract the nearest powers of 2 from the number
	* and keep those powers of two in an array
	* the array now defines the user permissions in decreasing order of authority
 */
