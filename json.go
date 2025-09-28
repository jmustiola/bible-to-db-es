package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func getChapterJsonContent(book Book) []byte {
	jsonBytes, err := json.Marshal(book.Chapters)
	if err != nil {
		log.Fatal(err)
	}
	return jsonBytes
}

func jsonToBook(filepath string) (Book, error) {
	bookFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return Book{}, err
	}

	defer bookFile.Close()

	var book Book
	info, err := bookFile.Stat()
	if err != nil {
		return Book{}, err
	}

	if info.Size() == 0 {
		return Book{}, fmt.Errorf("file is empty")
	}

	bytes, err := io.ReadAll(bookFile)
	if err != nil {
		return Book{}, err
	}

	err = json.Unmarshal(bytes, &book)
	if err != nil {
		return Book{}, err
	}

	return book, nil

}
