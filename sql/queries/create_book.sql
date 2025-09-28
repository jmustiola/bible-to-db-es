-- name: CreateBookProcedure :exec
CALL insert_book(sqlc.arg(book_id)::UUID, sqlc.arg(book_data)::JSONB);