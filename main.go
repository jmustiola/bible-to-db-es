package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/hiahir357/bible-to-db/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// loading env variables
	godotenv.Load(".env")

	// getting ENVs
	dbUrlStr, dataSourcePath := getProgramEnvs()

	// database connection
	conn, err := sql.Open("postgres", dbUrlStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer conn.Close()

	queries := database.New(conn)
	dbConnection := DBConnection{
		DB: queries,
	}

	switch os.Args[1] {
	case "gen":
		wg := &sync.WaitGroup{}
		result := make(chan Result)

		version := dbConnection.createVersion("Reina-Valera 1960", "RV1960")
		for _, filename := range JSON_FILENAMES {
			wg.Add(1)
			book, err := jsonToBook(filepath.Join(dataSourcePath, filename))
			if err != nil {
				log.Println("Error parsing the json file to model", filename, ":", err)
			}
			params := BookCreationParams{
				VersionId: version.ID,
				Book:      book,
			}
			go dbConnection.processBookCreation(params, result, wg)
		}

		go func() {
			wg.Wait()
			close(result)
		}()

		for res := range result {
			if res.Error != nil {
				log.Fatal("Error creating book in DB", res.Error)
			}
			log.Println(res.Message)
		}
	case "--search":
		if os.Args[2] != "" {
			result := dbConnection.filterByWord(os.Args[2])
			for _, row := range result {
				log.Println(row)
			}
		}

	}

}
