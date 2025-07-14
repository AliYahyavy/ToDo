package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func ErrHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RowCount() string {
	file, err := os.Open("data.csv")
	ErrHandler(err)
	defer file.Close()

	reader := csv.NewReader(file)
	var counter int = 1

	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		ErrHandler(err)
		counter++
	}
	return strconv.Itoa(counter)
}

func OpenFile(path string) (*os.File, *csv.Reader) {
	file, openFileErr := os.OpenFile(path, os.O_RDONLY, 0444)
	ErrHandler(openFileErr)

	return file, csv.NewReader(file)
}
