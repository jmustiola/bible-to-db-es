-- +goose Up
CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    book_order INT NOT NULL,
    num_chapters INT NOT NULL,
    num_verses INT NOT NULL,
    version_id UUID NOT NULL REFERENCES versions(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE books;