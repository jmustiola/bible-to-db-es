package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func jsonToBook(filename string) (Book, error) {
	bookFile, err := os.OpenFile("./data/json/"+filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return Book{}, err
	}

	defer bookFile.Close()

	var book Book
	info, err := bookFile.Stat()
	if err != nil {
		return Book{}, err
	}

	if info.Size() != 0 {
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

	return Book{}, fmt.Errorf("error: %v", err)
}
