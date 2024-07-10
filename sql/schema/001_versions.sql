-- +goose Up
CREATE TABLE versions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL,
    abbr VARCHAR(10)
);


-- +goose Down
DROP TABLE versions;