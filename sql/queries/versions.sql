-- name: CreateVersion :one
INSERT INTO 
    versions (id, created_at, updated_at, name, abbr)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;