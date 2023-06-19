// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user_profile.sql

package models

import (
	"context"
	"time"
)

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO "Profile" (
		"bio",
		"roomNumber",
		"sex",
		"age",
		"contactNumber",
		"maritalStatus",
		"bloodGroup",
		"origin",
		"lastVisit",
		"designation",
		"department",
		"userId"
	)
VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12
	)
RETURNING id, bio, "roomNumber", sex, age, "contactNumber", "maritalStatus", "bloodGroup", origin, "lastVisit", designation, department, "userId", "createdAt", "updatedAt", "isDeleted"
`

type CreateUserProfileParams struct {
	Bio           string    `json:"bio"`
	RoomNumber    int32     `json:"roomNumber"`
	Sex           int32     `json:"sex"`
	Age           int32     `json:"age"`
	ContactNumber string    `json:"contactNumber"`
	MaritalStatus int32     `json:"maritalStatus"`
	BloodGroup    int32     `json:"bloodGroup"`
	Origin        string    `json:"origin"`
	LastVisit     time.Time `json:"lastVisit"`
	Designation   string    `json:"designation"`
	Department    string    `json:"department"`
	UserId        int32     `json:"userId"`
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, createUserProfile,
		arg.Bio,
		arg.RoomNumber,
		arg.Sex,
		arg.Age,
		arg.ContactNumber,
		arg.MaritalStatus,
		arg.BloodGroup,
		arg.Origin,
		arg.LastVisit,
		arg.Designation,
		arg.Department,
		arg.UserId,
	)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Bio,
		&i.RoomNumber,
		&i.Sex,
		&i.Age,
		&i.ContactNumber,
		&i.MaritalStatus,
		&i.BloodGroup,
		&i.Origin,
		&i.LastVisit,
		&i.Designation,
		&i.Department,
		&i.UserId,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return i, err
}