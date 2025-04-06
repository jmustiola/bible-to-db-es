package main

import (
	"log"
	"os"
)

func getProgramEnvs() (dbURLString, dataSourcePath string) {
	dbURLString = os.Getenv("DATABASE_URL")
	if dbURLString == "" {
		log.Fatal("DB_URL env variable not found")
	}

	dataSourcePath = os.Getenv("DATA_SOURCE_PATH")
	if dataSourcePath == "" {
		log.Fatal("DATA_SOURCE_PATH env variable not found")
	}
	return
}
