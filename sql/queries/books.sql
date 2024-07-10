-- name: CreateBook :one
INSERT INTO 
    books (id, created_at, updated_at, name, book_order, num_chapters, num_verses, version_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;