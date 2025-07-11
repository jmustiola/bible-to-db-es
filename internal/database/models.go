// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	BookOrder   int32
	NumChapters int32
	NumVerses   int32
	VersionID   uuid.UUID
}

type Chapter struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	NumChapter int32
	NumVerses  int32
	BookID     uuid.UUID
}

type Verse struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	NumVerse  int32
	Text      string
	ChapterID uuid.UUID
}

type Version struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Abbr      sql.NullString
}
