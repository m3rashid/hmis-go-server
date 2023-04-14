-- name: CreateNewRole :one
INSERT INTO "Role" ("name", "description") VALUES ($1, $2) RETURNING "id";

-- name: CreateNewPermission :one
INSERT INTO "Permission" ("name", "description", "resourceId") VALUES ($1, $2, $3) RETURNING "id";
