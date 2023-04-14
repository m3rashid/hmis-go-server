-- name: CreateAdmin :one
INSERT INTO "User" ("userId", "name", "email", "password") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateDeveloper :one
INSERT INTO "User" ("userId", "name", "email", "password") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateUser :one
INSERT INTO "User" ("userId", "name", "email", "password") VALUES ($1, $2, $3, $4) RETURNING *;
