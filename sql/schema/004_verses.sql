-- +goose Up
CREATE TABLE verses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    num_verse INT NOT NULL,
    text TEXT NOT NULL,
    chapter_id UUID NOT NULL REFERENCES chapters(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE verses;