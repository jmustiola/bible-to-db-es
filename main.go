package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/hiahir357/bible-to-db/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConnection struct {
	DB *database.Queries
}

func main() {
	// loading env variables
	godotenv.Load(".env")

	// getting DB_URL
	var dbUrlStr string = os.Getenv("DB_URL")
	if dbUrlStr == "" {
		log.Fatal("DB_URL env variable not found")
	}

	// database connection
	conn, err := sql.Open("postgres", dbUrlStr)
	if err != nil {
		log.Fatal("Cannot connect to databse", err)
	}
	queries := database.New(conn)
	dbConnection := DBConnection{
		DB: queries,
	}

	wg := &sync.WaitGroup{}

	version := dbConnection.createVersion("Reina-Valera 1960", "RV1960")
	for i, filename := range JSON_FILENAMES {
		wg.Add(1)
		book, err := jsonToBook(filename)
		if err != nil {
			log.Println("Error parsing the json file to book", filename)
		}
		go create(version.ID, book, int32(i), &dbConnection, wg)
	}
	wg.Wait()

}

func create(versionId uuid.UUID, book Book, bookOrder int32, dbConnection *DBConnection, wg *sync.WaitGroup) {
	defer wg.Done()

	b, err := dbConnection.createBook(versionId, book, bookOrder)
	if err != nil {
		log.Fatal("Error while creating book:", err)
	}
	for _, chapter := range book.Chapters {
		c, err := dbConnection.createChapter(b.ID, chapter)
		if err != nil {
			log.Fatal("Error while creating chapter:", err)
		}
		for _, verse := range chapter.Verses {
			_, err := dbConnection.createVerse(c.ID, verse)
			if err != nil {
				log.Fatal("Error while creating verse:", err)
			}
		}
	}
	log.Printf("%v - Book %s created successfully", bookOrder, book.Name)
}
