package main

import (
	"github.com/google/uuid"
	"github.com/hiahir357/bible-to-db/internal/database"
)

type DBConnection struct {
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
	BookOrder int32
	VersionId uuid.UUID
	Book      Book
}
