-- name: CreateChapter :one
INSERT INTO 
    chapters (id, created_at, updated_at, num_chapter, num_verses, book_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;