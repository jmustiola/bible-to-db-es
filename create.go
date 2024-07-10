package main

import (
	"context"
	"database/sql"
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
		log.Fatal("Error creating version in DB", err)
	}
	log.Println("Version created succesfully")
	return v
}

func (dbConnection *DBConnection) createBook(version database.Version, book Book, book_order int32, wg *sync.WaitGroup) {
	defer wg.Done()
	b, err := dbConnection.DB.CreateBook(context.Background(), database.CreateBookParams{
		ID:          uuid.New(),
		VersionID:   version.ID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        book.Name,
		NumChapters: book.TotalChapters,
		NumVerses:   book.TotalVerses,
		BookOrder:   book_order,
	})
	if err != nil {
		log.Fatal("Error creating book in DB", err)
	}

	for _, chapter := range book.Chapters {
		c, err := dbConnection.DB.CreateChapter(context.Background(), database.CreateChapterParams{
			ID:         uuid.New(),
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
			NumChapter: chapter.NumChapter,
			NumVerses:  chapter.TotalVerses,
			BookID:     b.ID,
		})
		if err != nil {
			log.Fatal("Error creating book in DB", err)
		}

		for _, verse := range chapter.Verses {
			_, err := dbConnection.DB.CreateVerse(context.Background(), database.CreateVerseParams{
				ID:        uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				NumVerse:  verse.VersNum,
				Text:      verse.Text,
				ChapterID: c.ID,
			})
			if err != nil {
				log.Fatal("Error creating book in DB", err)
			}
		}
	}
	log.Printf("%v - Book %s created successfully", book_order, book.Name)

}
