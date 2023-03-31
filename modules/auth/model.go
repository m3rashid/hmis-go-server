package auth

import (
	"errors"
	"fmt"

	"github.com/m3rashid-org/hmis-go-server/utils/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"index:idx_email,unique"`
	Password    string
	Permissions []string
}

type Address struct {
	Pincode             int
	State               string
	Country             string `gorm:"default:india"`
	Street              string
	BuildingHouseNumber string
	RoomNumber          string
}

func FindUserByEmail(email string) (*User, db.SearchResult) {
	var user User
	result := db.Result(db.DB.Where("email = ?", email).First(&user).Error)
	return &user, result

}

func FindUserByID(id string) (*User, db.SearchResult) {
	var user User
	result := db.Result(db.DB.Where("id = ?", id).First(&user).Error)
	return &user, result
}

func FindUserByColum(colum string, value interface{}) (*User, bool) {
	var user User
	qs := fmt.Sprintf("%s = ?", colum)
	err := db.DB.Where(qs, value).First(&user).Error
	return &user, errors.Is(err, gorm.ErrRecordNotFound)
}
