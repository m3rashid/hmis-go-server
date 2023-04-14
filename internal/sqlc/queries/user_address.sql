-- name: CreateNewAddress :one
INSERT INTO "Address" (
		"city",
		"state",
		"country",
		"pinCode",
		"roomNumber",
		"buildingNumber",
		"userId"
	)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;