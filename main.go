package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

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
		go dbConnection.createBook(version, book, int32(i+1), wg)
	}
	wg.Wait()

}
