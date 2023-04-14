-- name: CreateResource :one
INSERT INTO "Resource" ("name", "description", "totalLevelsSum")
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetAllResources :many
SELECT *
FROM "Resource"
WHERE "isDeleted" = FALSE;


-- name: GetResourceByName :one
SELECT *
FROM "Resource"
WHERE "name" = $1
	AND "isDeleted" = FALSE;


-- name: GetResourceById :one
SELECT *
FROM "Resource"
WHERE "id" = $1
	AND "isDeleted" = FALSE;


-- name: GetResourceByLevel :many
SELECT *
FROM "Resource"
WHERE "totalLevelsSum" = $1
	AND "isDeleted" = FALSE;


-- name: GetResourceByLevelRange :many
SELECT *
FROM "Resource"
WHERE "totalLevelsSum" >= $1
	AND "totalLevelsSum" <= $2
	AND "isDeleted" = FALSE;


-- name: GetResourceDetailsWithPermissions :one
SELECT *
FROM "Resource"
	INNER JOIN "Permission" ON "Resource"."id" = "Permission"."resourceId"
WHERE "Resource"."id" = $1;


-- name: UpdateResourceAllFields :one
UPDATE "Resource"
SET "name" = $1,
	"description" = $2,
	"totalLevelsSum" = $3
WHERE "id" = $4
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdateResourceName :one
UPDATE "Resource"
SET "name" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdateResourceDescription :one
UPDATE "Resource"
SET "description" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: UpdateResourceTotalLevelSum :one
UPDATE "Resource"
SET "totalLevelsSum" = $1
WHERE "id" = $2
	AND "isDeleted" = FALSE
RETURNING *;


-- name: GetAllDeletedResources :many
SELECT *
FROM "Resource"
WHERE "isDeleted" = FALSE;


-- name: UnDeleteResource :one
UPDATE "Resource"
SET "isDeleted" = FALSE
WHERE "id" = $1
RETURNING *;