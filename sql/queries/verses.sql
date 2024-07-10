-- name: CreateVerse :one
INSERT INTO 
    verses (id, created_at, updated_at, num_verse, text, chapter_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;