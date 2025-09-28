-- +goose Up
CREATE TABLE chapters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    num_chapter INT NOT NULL,
    num_verses INT NOT NULL,
    book_id UUID NOT NULL REFERENCES books(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE chapters;