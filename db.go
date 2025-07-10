package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
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
	// log.Println("Version created succesfully")
	return v
}

func (dbConnection *DBConnection) createBook(versionId uuid.UUID, book Book) (*database.Book, error) {
	b, err := dbConnection.DB.CreateBook(context.Background(), database.CreateBookParams{
		ID:          uuid.New(),
		VersionID:   versionId,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        book.Name,
		NumChapters: book.TotalChapters,
		NumVerses:   book.TotalVerses,
		BookOrder:   int32(book.BookOrder),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating book in DB: %w", err)
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
		return nil, fmt.Errorf("error creating chapter in DB: %w", err)
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
		return nil, fmt.Errorf("error creating verse in DB: %w", err)
	}
	return &v, nil
}

func (dbConnection *DBConnection) processBookCreation(params BookCreationParams, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	b, err := dbConnection.createBook(params.VersionId, params.Book)
	if err != nil {
		result <- Result{Message: "Fatal error", Error: err}
		return
	}
	for _, chapter := range params.Book.Chapters {
		c, err := dbConnection.createChapter(b.ID, chapter)
		if err != nil {
			result <- Result{Message: "Fatal error", Error: err}
			return
		}
		for _, verse := range chapter.Verses {
			_, err := dbConnection.createVerse(c.ID, verse)
			if err != nil {
				result <- Result{Message: "Fatal error", Error: err}
				return
			}
		}
	}
	result <- Result{Message: fmt.Sprintf("%v - Book %s created successfully", params.Book.BookOrder, params.Book.Name), Error: nil}
}

func (dbConnection *DBConnection) filterByWord(word string) []database.GetFilteredVersesByWordRow {
	sqlString := sql.NullString{}
	if word != "" {
		sqlString.String = word
		sqlString.Valid = true
	}
	result, err := dbConnection.DB.GetFilteredVersesByWord(context.Background(), sqlString)
	if err != nil {
		log.Fatal("Error getting filtered verses by word", err)
	}
	return result
}
