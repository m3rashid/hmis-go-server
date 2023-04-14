-- name: CreateUserProfile :one
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
RETURNING *;