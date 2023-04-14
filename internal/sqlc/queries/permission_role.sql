-- name: CreateNewRole :one
INSERT INTO "Role" ("name", "description")
VALUES ($1, $2)
RETURNING *;


-- name: GetAllRoles :many
SELECT *
FROM "Role"
WHERE "isDeleted" = FALSE;


-- name: GetRoleByName :one
SELECT *
FROM "Role"
WHERE "name" = $1
	AND "isDeleted" = FALSE;


-- name: GetRoleById :one
SELECT *
FROM "Role"
WHERE "id" = $1
	AND "isDeleted" = FALSE;


-- name: GetRoleDetailsWithPermissions :one
SELECT *
FROM "Role"
	INNER JOIN "Permission" ON "Role"."id" = "Permission"."roleId"
WHERE "Role"."id" = $1;


-- name: UpdateRoleAllFields :one
UPDATE "Role"
SET "name" = $1,
	"description" = $2
WHERE "id" = $3
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdateRoleName :one
UPDATE "Role"
SET "name" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdateRoleDescription :one
UPDATE "Role"
SET "description" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: GetAllDeletedRoles :many
SELECT *
FROM "Role"
WHERE "isDeleted" = FALSE;


-- name: UnDeleteRole :one
UPDATE "Role"
SET "isDeleted" = FALSE
WHERE "id" = $1
RETURNING *;