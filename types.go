package main

import (
	"github.com/google/uuid"
	"github.com/hiahir357/bible-to-db/internal/database"
)

type Repository struct {
	DB *database.Queries
}

type ProgramEnvs struct {
	DATABASE_URL     string
	DATA_SOURCE_PATH string
}

type Result struct {
	Message string
	Error   error
}

type BookCreationParams struct {
	VersionId uuid.UUID
	Book      Book
}
