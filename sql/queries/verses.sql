-- name: CreateVerse :one
INSERT INTO 
    verses (id, created_at, updated_at, num_verse, text, chapter_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFilteredVersesByWord :many
SELECT b.name as Libro,
       c.num_chapter as Capitulo,
       v.num_verse as Versiculo,
       v.text as Texto
FROM public.verses as v
JOIN public.chapters as c ON c.id = v.chapter_id
JOIN public.books as b ON b.id = c.book_id
WHERE UNACCENT(v.text) ~* UNACCENT('\m' || $1 || '\M');