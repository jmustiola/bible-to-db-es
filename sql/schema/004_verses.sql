-- +goose Up
CREATE TABLE verses (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    num_verse INT NOT NULL,
    text TEXT NOT NULL,
    chapter_id UUID NOT NULL REFERENCES chapters(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE verses;