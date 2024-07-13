package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hiahir357/bible-to-db/internal/database"
)

func (dbConnection *DBConnection) createVersion(versionName string, versionAbbr string) database.Version {
	abbr := sql.NullString{}
	if versionAbbr != "" {
		abbr.String = versionAbbr
		abbr.Valid = true
	}

	v, err := dbConnection.DB.CreateVersion(context.Background(), database.CreateVersionParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      versionName,
		Abbr:      abbr,
	})
	if err != nil {
		log.Fatal("Error creating version in DB ", err)
	}
	log.Println("Version created succesfully")
	return v
}

func (dbConnection *DBConnection) createBook(versionId uuid.UUID, book Book, book_order int32) (*database.Book, error) {
	b, err := dbConnection.DB.CreateBook(context.Background(), database.CreateBookParams{
		ID:          uuid.New(),
		VersionID:   versionId,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        book.Name,
		NumChapters: book.TotalChapters,
		NumVerses:   book.TotalVerses,
		BookOrder:   book_order,
	})
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (dbConnection *DBConnection) createChapter(bookId uuid.UUID, chapter Chapter) (*database.Chapter, error) {
	c, err := dbConnection.DB.CreateChapter(context.Background(), database.CreateChapterParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		NumChapter: chapter.NumChapter,
		NumVerses:  chapter.TotalVerses,
		BookID:     bookId,
	})
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (dbConnection *DBConnection) createVerse(chapterId uuid.UUID, verse Verse) (*database.Verse, error) {
	v, err := dbConnection.DB.CreateVerse(context.Background(), database.CreateVerseParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		NumVerse:  verse.VersNum,
		Text:      verse.Text,
		ChapterID: chapterId,
	})
	if err != nil {
		return nil, err
	}
	return &v, nil
}
