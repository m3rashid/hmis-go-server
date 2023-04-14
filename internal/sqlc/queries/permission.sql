-- name: CreateNewPermission :one
INSERT INTO "Permission" ("name", "description", "resourceId")
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetAllPermissions :many
SELECT *
FROM "Permission"
WHERE "isDeleted" = FALSE;


-- name: GetPermissionByName :one
SELECT *
FROM "Permission"
WHERE "name" = $1
	AND "isDeleted" = FALSE;


-- name: GetPermissionById :one
SELECT *
FROM "Permission"
WHERE "id" = $1
	AND "isDeleted" = FALSE;


-- name: GetPermissionByResource :many
SELECT *
FROM "Permission"
WHERE "resourceId" = $1
	AND "isDeleted" = FALSE;


-- name: GetPermissionDetailsWithRoles :one
SELECT *
FROM "Permission"
	INNER JOIN "Role" ON "Permission"."id" = "Role"."permissionId"
WHERE "Permission"."id" = $1;


-- name: UpdatePermissionAllFields :one
UPDATE "Permission"
SET "name" = $1,
	"description" = $2,
	"resourceId" = $3
WHERE "id" = $4
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdatePermissionName :one
UPDATE "Permission"
SET name = $1
WHERE "id" = $2
	AND DELETED = FALSE
RETURNING *;


-- name: UpdatePermissionDescription :one
UPDATE "Permission"
SET "description" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdatePermissionResourceId :one
UPDATE "Permission"
SET "resourceId" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: GetAllDeletedPermissions :many
SELECT *
FROM "Permission"
WHERE "isDeleted" = FALSE;


-- name: UnDeletePermission :one
UPDATE "Permission"
SET "isDeleted" = FALSE
WHERE "id" = $1
RETURNING *;